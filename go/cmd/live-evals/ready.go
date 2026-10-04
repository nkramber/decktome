package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// notReady is the answer of `ready` for a pull request that is not ready
// for the owner. main exits 3 on it, so the loop script tells a wait
// from a fault.
type notReady []string

func (n notReady) Error() string { return "not ready: " + strings.Join(n, "; ") }

// prView is the part of `gh pr view --json` that the check reads.
type prView struct {
	State             string `json:"state"`
	IsDraft           bool   `json:"isDraft"`
	Mergeable         string `json:"mergeable"`
	URL               string `json:"url"`
	StatusCheckRollup []struct {
		Typename   string `json:"__typename"`
		Name       string `json:"name"`
		Context    string `json:"context"`
		Status     string `json:"status"`
		Conclusion string `json:"conclusion"`
		State      string `json:"state"`
		StartedAt  string `json:"startedAt"`
	} `json:"statusCheckRollup"`
}

// readiness answers each reason that a pull request is not ready for
// the owner, and none for a ready one (D-1137). It reads the state on
// GitHub and never the word of the eval session.
func readiness(v prView, unresolved int) []string {
	var why []string
	if v.State != "OPEN" {
		why = append(why, "state "+v.State)
	}
	if v.IsDraft {
		why = append(why, "draft")
	}
	if v.Mergeable != "MERGEABLE" {
		why = append(why, "mergeable "+v.Mergeable)
	}
	// A check that ran more than once counts by its newest run, the run
	// the ruleset reads. RFC 3339 times in UTC sort as text.
	newest := map[string]string{}
	for _, c := range v.StatusCheckRollup {
		if name := c.Name + c.Context; c.StartedAt > newest[name] {
			newest[name] = c.StartedAt
		}
	}
	gate := false
	for _, c := range v.StatusCheckRollup {
		name := c.Name
		if name == "" {
			name = c.Context
		}
		if c.StartedAt < newest[c.Name+c.Context] {
			continue
		}
		result := c.Conclusion
		if c.Typename == "StatusContext" {
			result = c.State
		} else if c.Status != "COMPLETED" {
			why = append(why, name+" "+strings.ToLower(c.Status))
			continue
		}
		switch result {
		case "SUCCESS", "NEUTRAL", "SKIPPED":
		default:
			why = append(why, name+" "+strings.ToLower(result))
			continue
		}
		if name == "review-gate" && result == "SUCCESS" {
			gate = true
		}
	}
	if !gate {
		why = append(why, "no passed review-gate check")
	}
	if unresolved > 0 {
		why = append(why, strconv.Itoa(unresolved)+" open review threads")
	}
	return why
}

const threadsQuery = `query($owner:String!,$name:String!,$number:Int!){repository(owner:$owner,name:$name){pullRequest(number:$number){reviewThreads(first:100){nodes{isResolved}}}}}`

func ready(ctx context.Context, gh runner, pr int, out io.Writer) error {
	raw, err := gh(ctx, "gh", "pr", "view", strconv.Itoa(pr), "--json", "state,isDraft,mergeable,url,statusCheckRollup")
	if err != nil {
		return fmt.Errorf("gh pr view: %w", err)
	}
	var v prView
	if err := json.Unmarshal(raw, &v); err != nil {
		return fmt.Errorf("gh pr view: %w", err)
	}
	repo, err := gh(ctx, "gh", "repo", "view", "--json", "nameWithOwner", "--jq", ".nameWithOwner")
	if err != nil {
		return fmt.Errorf("gh repo view: %w", err)
	}
	owner, name, ok := strings.Cut(strings.TrimSpace(string(repo)), "/")
	if !ok {
		return fmt.Errorf("gh repo view: no owner in %q", repo)
	}
	raw, err = gh(ctx, "gh", "api", "graphql", "-f", "query="+threadsQuery, "-f", "owner="+owner, "-f", "name="+name,
		"-F", "number="+strconv.Itoa(pr))
	if err != nil {
		return fmt.Errorf("gh api graphql: %w", err)
	}
	var threads struct {
		Data struct {
			Repository struct {
				PullRequest struct {
					ReviewThreads struct {
						Nodes []struct {
							IsResolved bool `json:"isResolved"`
						} `json:"nodes"`
					} `json:"reviewThreads"`
				} `json:"pullRequest"`
			} `json:"repository"`
		} `json:"data"`
	}
	if err := json.Unmarshal(raw, &threads); err != nil {
		return fmt.Errorf("gh api graphql: %w", err)
	}
	unresolved := 0
	for _, n := range threads.Data.Repository.PullRequest.ReviewThreads.Nodes {
		if !n.IsResolved {
			unresolved++
		}
	}
	if why := readiness(v, unresolved); len(why) > 0 {
		return notReady(why)
	}
	_, err = fmt.Fprintf(out, "ready %s\n", v.URL)
	return err
}
