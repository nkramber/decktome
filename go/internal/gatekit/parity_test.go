package gatekit

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
	"testing"
)

// TestEveryToolBuildsWithTheSignalsOfTheApp is F-129 (D-706, D-708). The
// app hands the shortlist the top-list rate of the quality model, and the
// commander pool its commander signal. A tool that builds a deck without
// them measures a shortlist the app never builds, and no bar says so.
func TestEveryToolBuildsWithTheSignalsOfTheApp(t *testing.T) {
	files, err := filepath.Glob(filepath.Join("..", "..", "cmd", "*", "*.go"))
	if err != nil {
		t.Fatal(err)
	}
	files = append(files, filepath.Join("..", "agentsvc", "build.go"))
	checked := 0
	for _, path := range files {
		if strings.HasSuffix(path, "_test.go") || parityExempt[filepath.Base(filepath.Dir(path))] {
			continue
		}
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			t.Fatalf("parse %s: %v", path, err)
		}
		n, gaps := parityGaps(fset, f)
		checked += n
		for _, gap := range gaps {
			t.Error(gap)
		}
	}
	// Four tools and the app make seven requests, so a scan that reads
	// fewer has lost a file.
	if checked < 7 {
		t.Fatalf("the scan read %d requests, want 7 or more", checked)
	}
}

// TestParityScanFollowsARequestVariable proves the scan reads a request that
// a variable carries to the call, as well as a literal in the call.
func TestParityScanFollowsARequestVariable(t *testing.T) {
	const src = `package tool

func build() {
	req := candidates.Request{Theme: "lifegain"}
	list, err := cb.Build(idx, req)
	pool, err := cb.CommanderPool(idx, candidates.Request{Theme: "lifegain", CommanderSignal: signal})
	_, _, _ = list, pool, err
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "tool.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	checked, gaps := parityGaps(fset, f)
	if checked != 2 || len(gaps) != 1 || !strings.Contains(gaps[0], "Build takes a candidates.Request that sets no MetaBoost") {
		t.Errorf("checked %d requests, gaps %v; want 2 requests and the Build gap", checked, gaps)
	}
}

// parityExempt names the tools that build no deck. candidates-review lists
// shortlists for a person to grade the theme ranking (PR-6), and D-708
// covers the tools that build a deck.
var parityExempt = map[string]bool{"candidates-review": true}

// parityGaps reads every candidates.Request that a Build or a CommanderPool
// call takes in one file: a literal in the call, or a variable that the
// same function binds to a literal. It returns how many requests it read,
// and one line for each request that misses a signal of the app.
func parityGaps(fset *token.FileSet, f *ast.File) (int, []string) {
	want := map[string]string{"Build": "MetaBoost", "CommanderPool": "CommanderSignal"}
	checked := 0
	var gaps []string
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil {
			continue
		}
		bound := boundRequests(fn.Body)
		ast.Inspect(fn.Body, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			field, ok := want[sel.Sel.Name]
			if !ok {
				return true
			}
			for _, arg := range call.Args {
				lit, ok := arg.(*ast.CompositeLit)
				if id, isIdent := arg.(*ast.Ident); isIdent {
					lit, ok = bound[id.Name]
				}
				if !ok || !isCandidatesRequest(lit.Type) {
					continue
				}
				checked++
				if !setsField(lit, field) {
					gaps = append(gaps, fmt.Sprintf("%s: %s takes a candidates.Request that sets no %s",
						fset.Position(call.Pos()), sel.Sel.Name, field))
				}
			}
			return true
		})
	}
	return checked, gaps
}

// boundRequests maps each name that one function body assigns a
// candidates.Request literal to.
func boundRequests(body *ast.BlockStmt) map[string]*ast.CompositeLit {
	out := map[string]*ast.CompositeLit{}
	ast.Inspect(body, func(n ast.Node) bool {
		as, ok := n.(*ast.AssignStmt)
		if !ok {
			return true
		}
		for i, rhs := range as.Rhs {
			lit, ok := rhs.(*ast.CompositeLit)
			if !ok || !isCandidatesRequest(lit.Type) || i >= len(as.Lhs) {
				continue
			}
			if id, ok := as.Lhs[i].(*ast.Ident); ok {
				out[id.Name] = lit
			}
		}
		return true
	})
	return out
}

func isCandidatesRequest(expr ast.Expr) bool {
	sel, ok := expr.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "Request" {
		return false
	}
	pkg, ok := sel.X.(*ast.Ident)
	return ok && pkg.Name == "candidates"
}

func setsField(lit *ast.CompositeLit, field string) bool {
	for _, e := range lit.Elts {
		kv, ok := e.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		if key, ok := kv.Key.(*ast.Ident); ok && key.Name == field {
			return true
		}
	}
	return false
}
