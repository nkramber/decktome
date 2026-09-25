# Deploy after a merge, and roll back

`docs/setup-gcp.md` builds the deployment one time. This page runs after every merge to `main`. It also tells you how to return to the version before a bad deploy.

The project is `decktome-prod`, and the region is `us-central1`. The deploy of 2026-09-07 set both. Write in the commands: `TAG` is the short commit of the merge.

## 1. A merge to main deploys itself

Cloud Build builds and releases every merge to `main` (D-584). No step of section 4 or section 5 runs by hand. Two triggers read the diff of the merge, so a merge pays for what it changed.

| Trigger | It fires on | It does |
|---|---|---|
| `deploy-api`, `cloudbuild/api.yaml` | `go/**`, `docker/**` | Builds both images, deploys the API, points both jobs at the new image, and waits for `/readyz` |
| `deploy-web`, `cloudbuild/web.yaml` | `web/**`, `firebase.json` | Builds the web app, releases it, and reads the site |

A merge of documents alone starts no build at all.

Cloud Build has no queue, and the two triggers run apart. So each build reads the live commit before its deploy (REV-072, D-943). `/readyz` names the commit of the API in `version`, and `/version.json` names the commit of the web. `docs/tools/deploy_order.py` holds the rules:

- The API build skips its deploy, its jobs, and its check when the live API runs a later commit. The step `guard` writes `/workspace/.deploy-skip`, and each later step reads it.
- The web build of a merge that changed `go/` or `docker/` waits for the API of that commit, or a later one. The wait stops at 25 minutes, and the web build then fails. So a failed API build also stops the web of its merge.
- The web build skips its release when the live web runs a later commit.
- The API check waits until `/readyz` reads `ok` with the commit of the build. The web check reads the commit of `/version.json`.

A live commit that no read gives, or that git can not place, never stops a deploy. The log then says so. Two builds can still overlap during the one minute of one deploy.

Cloud Build runs the deploy, and not GitHub Actions, to keep the free Actions minutes for the checks (D-584). Cloud Build gives 2,500 build-minutes a month, and a minute costs $0.006 after that. GitHub gives 2,000 minutes a month for a private repository, and a minute costs $0.008. The two pools are apart, so a deploy never takes a minute the checks need.

The workflow `deploy` of GitHub Actions stays as the way back. Nothing starts it on its own, and a dispatch runs the same steps.

CAUTION: the triggers read `main` and no other branch (D-579). A pull request starts no build.

### 1.1 The one-time setup

The setup of `decktome-prod` ran on 2026-09-07, and both triggers are live. These steps stand for a new project, and for a repair of this one.

1. Connect the repository, in the region `us-central1`. Open Cloud Build, then Repositories, then the 2nd gen tab. Select Create host connection, choose GitHub, and authorize it. Then select Link repository and choose `nkramber/decktome`. This step needs a browser.

Note: the host connection asks for a Cloud KMS key, and it is optional. Leave it empty. Google then encrypts the stored credential with a key it manages, as it does for Firestore, the bucket, and Secret Manager. A key of your own adds a keyring and an IAM binding for the Cloud Build service agent. It also adds a charge for each key version, and a rotation to keep.

2. Read the names the connection got, and build the repository path from them:

```
gcloud builds connections list --region=us-central1
gcloud builds repositories list --connection=<the connection> --region=us-central1
```

3. Create the two triggers. Every part is regional, so the region of the trigger must equal the region of the connection:

```
REPO=projects/decktome-prod/locations/us-central1/connections/decktome-repository/repositories/nkramber-decktome
SA=projects/-/serviceAccounts/gh-deployer@decktome-prod.iam.gserviceaccount.com
gcloud builds triggers create github --name=deploy-api --region=us-central1 \
  --repository=$REPO --branch-pattern='^main$' --service-account=$SA \
  --build-config=cloudbuild/api.yaml --included-files='go/**,docker/**'
gcloud builds triggers create github --name=deploy-web --region=us-central1 \
  --repository=$REPO --branch-pattern='^main$' --service-account=$SA \
  --build-config=cloudbuild/web.yaml --included-files='web/**,firebase.json,firestore.indexes.json'
```

Note: the web build deploys `hosting,firestore:indexes`, so `deploy-web` also watches `firestore.indexes.json` (D-603). Without that path, a change to the index file deploys nothing, and F-73 comes back.

CAUTION: a trigger of a 2nd-gen repository needs `--service-account`. Without it the API answers `INVALID_ARGUMENT` and names no field. `--repo-owner` and `--repo-name` name a 1st-gen repository, and they never reach a 2nd-gen connection. A trigger with no `--region` lands in `global`, and it finds no connection of `us-central1`. Read `gcloud builds triggers list --region=us-central1` after each one.

Note: `us-central1` matches Artifact Registry, Cloud Run, Firestore, and the bucket, so a build pushes an image inside one region.

4. Grant the deployer what a build needs. `gh-deployer` is the account the GitHub workflow signs in as (D-582), so one identity deploys on both paths:

```
GH=gh-deployer@decktome-prod.iam.gserviceaccount.com
for role in roles/run.admin roles/artifactregistry.writer \
            roles/firebasehosting.admin roles/logging.logWriter \
            roles/firebaserules.admin roles/datastore.indexAdmin; do
  gcloud projects add-iam-policy-binding decktome-prod \
    --member=serviceAccount:$GH --role=$role --condition=None
done
```

CAUTION: a build takes a service account you made, and never the one Google manages. A trigger that names `PROJECT_NUMBER@cloudbuild.gserviceaccount.com` fails before its first step. The message reads "provide a user-managed service account or leave unset". `roles/logging.logWriter` belongs on the list, because the build writes its log with `CLOUD_LOGGING_ONLY`.

CAUTION: the web build fails with a 403 without the last two roles (D-605). The index deploy reads `firestore.rules` through the Firebase Rules API, and then it writes the indexes. `roles/firebaserules.admin` can also write the rules of the project, so it widens the account.

Note: the project policy of `decktome-prod` holds a condition, so a binding without `--condition=None` fails in a script. The message names the flag.

5. Let the deployer run the API and the jobs as their own accounts. This binding names those two accounts, and never every account of the project:

```
for sa in mtg-api mtg-worker; do
  gcloud iam service-accounts add-iam-policy-binding \
    $sa@decktome-prod.iam.gserviceaccount.com \
    --member=serviceAccount:$GH --role=roles/iam.serviceAccountUser
done
```

6. Give the web build its own account (D-931). The web build runs `pnpm install`, and an install script runs as the build account. `gh-deployer` holds `roles/run.admin` and `roles/firebaserules.admin`, so a bad package can deploy an API revision or open the rules. `web-deployer` releases Hosting and the indexes, and it can not write a rule.

State on 2026-09-25: `web-deployer` holds `roles/firebasehosting.admin`, `roles/datastore.indexAdmin`, `roles/logging.logWriter`, and the custom role `webDeployRulesTest`. The trigger `deploy-web` runs as `web-deployer` (PR-81, D-944). The steps below stand for a new project, and for a repair of this one.

```
WEB=web-deployer@decktome-prod.iam.gserviceaccount.com
gcloud iam roles create webDeployRulesTest --project=decktome-prod \
  --title="Web deploy: test and read Firebase rules" --stage=GA \
  --permissions=firebaserules.rulesets.test,firebaserules.rulesets.get,firebaserules.rulesets.list,firebaserules.releases.get,firebaserules.releases.list
gcloud projects add-iam-policy-binding decktome-prod --member=serviceAccount:$WEB \
  --role=projects/decktome-prod/roles/webDeployRulesTest --condition=None
```

CAUTION: do not switch the trigger before the custom role holds. The index deploy calls the `:test` method of the Rules API (D-605). `roles/firebaserules.viewer` holds no `firebaserules.rulesets.test`, so the web build then fails with a 403, and Hosting does not deploy.

Then switch the trigger. `gcloud builds triggers update github` refuses a trigger of a 2nd-gen repository, and a PATCH of the whole trigger body works (D-603). Version 533.0.0 of gcloud has no `builds triggers export` outside the beta group. The PATCH below ran on 2026-09-25, and it changed `serviceAccount` alone.

```
gcloud builds triggers describe deploy-web --region=us-central1 --format=json > /tmp/deploy-web.json
sed -i '' 's#serviceAccounts/gh-deployer@#serviceAccounts/web-deployer@#' /tmp/deploy-web.json
curl -s -X PATCH -H "Authorization: Bearer $(gcloud auth print-access-token)" \
  -H "Content-Type: application/json" --data @/tmp/deploy-web.json \
  https://cloudbuild.googleapis.com/v1/projects/decktome-prod/locations/us-central1/triggers/deploy-web
gcloud builds triggers describe deploy-web --region=us-central1 --format='value(serviceAccount)'
```

Read the next `deploy-web` build with `gcloud builds list --region=us-central1 --limit=2`. When it fails with a 403, put `gh-deployer` back with the same commands and the opposite `sed`. Then read the log for the permission it named.

Note: the Workload Identity Federation setup of D-582 stays. The GitHub workflow reads it when somebody runs the deploy by hand.

The Cloud Build files hold the four Firebase values as substitutions, so a build reads them from the repository. The GitHub workflow reads them from the repository variables of `nkramber/decktome` instead, and `gh variable list` prints what they hold.

CAUTION: never run a command that holds a placeholder in angle brackets. The command writes the placeholder itself as the value. On 2026-09-07 a run of such a command wrote `<the apiKey of firebase apps:sdkconfig>` over the real key. Every sign-in then read `auth/invalid-api-key`, because the app carries that value to the browser (F-65). Read the value first, then write it, then read it back.

## 2. What each merge changes

Read the merged diff first. This table names the part to deploy.

| The merge touches | Deploy this | Section |
|---|---|---|
| Any file under `go/` | The API and the two jobs | 4 |
| Any file under `web/` | The web app | 5 |
| `firestore.rules` or `firestore.indexes.json` | The rules and the indexes | 6 |
| The `hosting` block of `firebase.json` | The web app | 5 |
| Only `docs/` | Nothing | - |

The API binary and the worker binary both import `go/internal/quality`. Therefore every change under `go/` needs two new images.

## 3. Before a deploy you run by hand

1. Run `git checkout main`.
2. Run `git pull`.
3. Run `git status`. The tree must be clean.
4. Run `git rev-parse --abbrev-ref HEAD`. It must print `main`.
5. Run `gcloud config configurations activate decktome`.
6. Run `gcloud config get-value project`. It must print `decktome-prod`.
7. Run `TAG=$(git rev-parse --short HEAD)`. Every command below reads this tag.

CAUTION: deploy `main` alone (D-579). No other branch goes to production, for any reason. A clean tree is not enough, because a branch commit is clean too. Read step 4 before each build.

## 3. Check the code

Run the free checks before a deploy. Each one costs nothing.

1. Run `make verify`. It runs every check of the `verify` workflow on this machine, for nothing (D-578).

`make verify` covers the proto, go, web, shell, eval, and docker lanes. The emulator lane needs `make store-check` against a local Firestore emulator, and `govulncheck` runs on the weekly schedule.

A failed check stops the deploy. Repair the code on a branch, and merge the repair first.

CAUTION: the ruleset of `main` requires the `review-gate` check, each pull request job of `verify`, and each resolved thread, with no bypass (D-815, D-828). A repair of code merges only after a Codex record approves it and each check is green. For an urgent fault, roll back first with section 8, which needs no merge.

## 4. Deploy the API and the jobs

Build the two images from the repo root. Cloud Run runs `linux/amd64` images only.

```
docker build --platform linux/amd64 -f docker/api.Dockerfile \
  --build-arg VERSION=$(git rev-parse HEAD) \
  -t us-central1-docker.pkg.dev/decktome-prod/mtg/api:$TAG .
docker build --platform linux/amd64 -f docker/worker.Dockerfile \
  -t us-central1-docker.pkg.dev/decktome-prod/mtg/worker:$TAG .
docker push us-central1-docker.pkg.dev/decktome-prod/mtg/api:$TAG
docker push us-central1-docker.pkg.dev/decktome-prod/mtg/worker:$TAG
```

Deploy the API. The command makes a new revision and moves all traffic to it.

```
gcloud run deploy mtg-api \
  --image us-central1-docker.pkg.dev/decktome-prod/mtg/api:$TAG \
  --region us-central1
```

Note: this command keeps every environment variable, secret, and flag of the revision before it. A redeploy needs `--allow-unauthenticated` no more, because that flag writes an IAM policy one time.

Update the two jobs. A job has no traffic, so the next execution reads the new image.

```
gcloud run jobs update mtg-snapshot --region us-central1 \
  --image us-central1-docker.pkg.dev/decktome-prod/mtg/worker:$TAG
gcloud run jobs update mtg-meta --region us-central1 \
  --image us-central1-docker.pkg.dev/decktome-prod/mtg/worker:$TAG
```

To change an environment variable or a secret, add the flag to the same command. Section 11 of `docs/setup-gcp.md` names every flag of the first deploy.

The Pushover secrets of PR-75 need one step outside a deploy (D-893). Section 8 of `docs/setup-gcp.md` creates them. Then mount them on the service one time:

```
gcloud run services update mtg-api --region us-central1 \
  --update-secrets=PUSHOVER_APP_TOKEN=pushover-app-token:1,PUSHOVER_USER_KEY=pushover-user-key:1
```

The session of 2026-09-24 ran this step on the image of `930d1a6`, and revision `mtg-api-00078-hv2` holds the two secrets. Each later deploy keeps them. The API log reads `feedback notices on` at start when both are set. To stop the notices, run the same command with `--remove-secrets=PUSHOVER_APP_TOKEN,PUSHOVER_USER_KEY`.

## 5. Deploy the web app

The four `VITE_FIREBASE_` values never change. `docs/setup-gcp.md` section 13 holds them.

1. Run `nvm use 22`. The Vite build needs Node 22.
2. Run the build with the variables of section 13.
3. Write the commit to the release: `printf '{"commit":"%s"}\n' "$(git rev-parse HEAD)" > web/apps/web/dist/version.json`.
4. Run `firebase deploy --only hosting` from the repo root.

The site is live on `https://decktome.com` and on `https://decktome-prod.web.app`. Firebase Hosting keeps every earlier version, and section 8 returns to one.

Note: an installed app takes a new release on its next load, and the page reloads when the new service worker activates (D-692). Before #158, the page kept the old shell for that load (F-122). The release of #196 proved that rule on 2026-09-20, and `make self-reload-check` proves it again for nothing. `docs/reference/self-reload-2026-09-20.md` holds the method and the three limits of the check. The Hosting rewrite answers a missing file with `index.html`. So a check for a removed file reads the `index.html` of the release and the precache list of `sw.js`.

## 6. Deploy the Firestore rules and indexes

1. Run `firebase deploy --only firestore` from the repo root.

Note: this command adds a new index. It removes no index that the repo dropped. Remove an unused index by hand in the Firebase console.

## 7. Check the deployment

1. Run `curl -s "$(gcloud run services describe mtg-api --region us-central1 --format='value(status.url)')/readyz"`. It must read `"status":"ok"`, and `version` must name the commit of the merge.
2. Open `https://decktome.com` and sign in.
3. Build one small deck. The build proves the model keys and the card snapshot.
4. Read the logs for an error:

```
gcloud logging read 'resource.type="cloud_run_revision"
  AND resource.labels.service_name="mtg-api"' --limit 30 --freshness=30m \
  --format='value(timestamp,severity,textPayload,jsonPayload.message)'
```

Note: `/readyz` reads `starting` for about 90 seconds after a cold start. Do not read that as a failure. `/healthz` never answers on a `run.app` URL, because the Google frontend takes that path.

CAUTION: `gcloud run services logs read` crashed with a `TypeError` on gcloud 533.0.0 against this service. The command above reads the same logs.

## 8. Roll back

### 8.1 The API

Cloud Run keeps every revision. A rollback moves the traffic, and it needs no build.

1. Run `gcloud run revisions list --service mtg-api --region us-central1`.
2. Read the name of the last good revision.
3. Run `gcloud run services update-traffic mtg-api --region us-central1 --to-revisions=REVISION=100`.

The change takes seconds. Run `--to-latest` to return the traffic to the newest revision.

CAUTION: the pin holds until `--to-latest`. A deploy during the pin makes a new revision that takes no traffic. The deploy then fails at its check step, because `docs/tools/traffic_check.py` reads the traffic of the newest revision (REV-013). Run `--to-latest` after the fix merges, then run the deploy again.

CAUTION: during the pin, `/readyz` names the commit of the old revision. So the web build of a merge that changed `go/` or `docker/` waits for its API, and it fails after 25 minutes (D-943). Run the web build again after `--to-latest`.

A revision older than `mtg-api-00077-vwp` holds no Pushover secret, so a rollback to it sends no notice of a verdict. The store still keeps each verdict.

### 8.2 The jobs

A job keeps no revision history for a rollback command. Point the job at the earlier image tag.

```
gcloud run jobs update mtg-snapshot --region us-central1 \
  --image us-central1-docker.pkg.dev/decktome-prod/mtg/worker:OLDTAG
```

Read the tags of the earlier images with `gcloud artifacts docker tags list us-central1-docker.pkg.dev/decktome-prod/mtg/worker`.

### 8.3 The web app

firebase-tools 14.14.0 has no `hosting:rollback` command. Two paths return the site to an earlier version.

- The console: open the Firebase console, then Hosting, then the release history. Select the three dots of a good version, then Rollback. This path is the fast one.
- The repo: check out the earlier commit, build the web app again, and deploy it. This path also returns the source to a known state.

### 8.4 The Firestore rules

1. Run `git checkout GOODCOMMIT -- firestore.rules`.
2. Run `firebase deploy --only firestore:rules`.
3. Run `git checkout main -- firestore.rules` to restore the tree.

CAUTION: a rollback of the code does not undo a change of the data. A new version can write a document that an old version cannot read. Read the store code of the merge before a rollback. Check the package under `go/internal` that writes the document for a schema change.

Note: a stored deck or chat session is gzip JSON of a proto message. The decoder of `go/internal/gzstore` drops each field that the proto of the old version does not know, so the old version reads the document (REV-014). A write of the old version then loses those fields. A new proto field alone needs no other step.

## 9. What a rollback does not repair

- A Firestore document that the new version wrote stays as it is.
- An index that a deploy added stays until you remove it by hand.
- A secret version stays active until you disable it.
- A Pushover notice that the API sent stays on the device and with Pushover.
- A card snapshot in the bucket stays. The worker keeps three versions (`cards.KeepVersions`).
- Section 11 restores the Firestore data from a daily backup.

## 11. Restore the Firestore data

The owner turned on a daily backup of the `(default)` database on 2026-09-23 (REV-037, D-936). A read of 2026-09-25 gave these facts:

- The schedule runs each day, and each backup stays 10 days (`864000s`).
- Two backups read READY: 2026-09-24 at 07:07 UTC, and 2026-09-25 at 07:19 UTC.
- Point-in-time recovery reads disabled, and delete protection reads disabled.

A restore writes a new database. It never writes over `(default)`. Do these steps:

1. Run `gcloud firestore backups list --project=decktome-prod --format="value(name,snapshotTime,state)"`.
2. Choose the newest READY backup before the damage.
3. Set `BACKUP` to its full name, and `DEST` to a new id, for example `restore-20260925`.
4. Run `gcloud firestore databases restore --project=decktome-prod --source-backup="$BACKUP" --destination-database="$DEST"`.
5. Read the damaged documents in `$DEST`, and compare each one with `(default)`.
6. Copy each document back into `(default)` with a script that the owner reads first.
7. Run `gcloud firestore databases delete --project=decktome-prod --database="$DEST"` after the repair.

CAUTION: a copy back writes production user data. Copy only the documents that the damage changed. A later write of a user sits in `(default)` alone, and a whole copy removes it.

## 12. Refresh the image digests

Each step of `cloudbuild/api.yaml` and `cloudbuild/web.yaml` names its image by digest (D-931). So a new release of an image changes no build, and it brings no fix either. `docs/tools/test_deploy_workflow.py` refuses a step with no digest.

1. Run `gcloud container images describe gcr.io/cloud-builders/docker:latest --format='value(image_summary.digest)'`.
2. Run the same command for `gcr.io/google.com/cloudsdktool/cloud-sdk:latest`.
3. Run `docker buildx imagetools inspect node:22.23.2`, and read its `Digest` line.
4. Replace each old digest in the two files, and open a pull request.

The digests of 2026-09-25 are the first ones. A merge that touches `go/` then builds with the new digests of `cloudbuild/api.yaml`.

## 10. Sources and dates

| Fact | Source |
|---|---|
| The Cloud Run rollback command and the `--to-revisions` form | `gcloud run services update-traffic --help`, gcloud 533.0.0, read 2026-09-07 |
| The `--image` flag of a job update | `gcloud run jobs update --help`, read 2026-09-07 |
| firebase-tools has no `hosting:rollback` command | `firebase --help`, firebase-tools 14.14.0, read 2026-09-07 |
| `hosting:clone` reads a channel, not a version | `firebase hosting:clone --help`, read 2026-09-07 |
| Both binaries import `go/internal/quality` | `go list -deps ./cmd/api` and `./cmd/worker`, read 2026-09-07 |
| The project, the region, and every name | `docs/setup-gcp.md` and the deploy of 2026-09-07 |
| `/readyz` is the readiness path, and the frontend takes `/healthz` | The deploy of 2026-09-07, five paths compared |
| `--update-secrets` keeps the other secrets, and a deploy keeps the secrets of the revision before it | The update of `mtg-api` on 2026-09-24, revisions `mtg-api-00077-vwp` and `mtg-api-00078-hv2` |
| The Vite build needs Node 22 | `docs/setup.md` and the toolchain of the repo |
| The restore writes a new database, with `--source-backup` and `--destination-database` | `gcloud firestore databases restore --help`, gcloud 533.0.0, read 2026-09-25 |
| The backup schedule, the two backups, and the disabled recovery and protection | `gcloud firestore backups schedules list`, `gcloud firestore backups list`, and `gcloud firestore databases describe`, read 2026-09-25 |
| `roles/firebaserules.viewer` holds no `firebaserules.rulesets.test` | `gcloud iam roles describe`, read 2026-09-25 |
