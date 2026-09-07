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
  --build-config=cloudbuild/web.yaml --included-files='web/**,firebase.json'
```

CAUTION: a trigger of a 2nd-gen repository needs `--service-account`. Without it the API answers `INVALID_ARGUMENT` and names no field. `--repo-owner` and `--repo-name` name a 1st-gen repository, and they never reach a 2nd-gen connection. A trigger with no `--region` lands in `global`, and it finds no connection of `us-central1`. Read `gcloud builds triggers list --region=us-central1` after each one.

Note: `us-central1` matches Artifact Registry, Cloud Run, Firestore, and the bucket, so a build pushes an image inside one region.

4. Grant the deployer what a build needs. `gh-deployer` is the account the GitHub workflow signs in as (D-582), so one identity deploys on both paths:

```
GH=gh-deployer@decktome-prod.iam.gserviceaccount.com
for role in roles/run.admin roles/artifactregistry.writer \
            roles/firebasehosting.admin roles/logging.logWriter; do
  gcloud projects add-iam-policy-binding decktome-prod \
    --member=serviceAccount:$GH --role=$role --condition=None
done
```

CAUTION: a build takes a service account you made, and never the one Google manages. A trigger that names `PROJECT_NUMBER@cloudbuild.gserviceaccount.com` fails before its first step. The message reads "provide a user-managed service account or leave unset". `roles/logging.logWriter` belongs on the list, because the build writes its log with `CLOUD_LOGGING_ONLY`.

Note: the project policy of `decktome-prod` holds a condition, so a binding without `--condition=None` fails in a script. The message names the flag.

5. Let the deployer run the API and the jobs as their own accounts. This binding names those two accounts, and never every account of the project:

```
for sa in mtg-api mtg-worker; do
  gcloud iam service-accounts add-iam-policy-binding \
    $sa@decktome-prod.iam.gserviceaccount.com \
    --member=serviceAccount:$GH --role=roles/iam.serviceAccountUser
done
```

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

## 4. Deploy the API and the jobs

Build the two images from the repo root. Cloud Run runs `linux/amd64` images only.

```
docker build --platform linux/amd64 -f docker/api.Dockerfile \
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

## 5. Deploy the web app

The four `VITE_FIREBASE_` values never change. `docs/setup-gcp.md` section 13 holds them.

1. Run `nvm use 22`. The Vite build needs Node 22.
2. Run the build with the variables of section 13.
3. Run `firebase deploy --only hosting` from the repo root.

The site is live on `https://decktome.com` and on `https://decktome-prod.web.app`. Firebase Hosting keeps every earlier version, and section 8 returns to one.

## 6. Deploy the Firestore rules and indexes

1. Run `firebase deploy --only firestore` from the repo root.

Note: this command adds a new index. It removes no index that the repo dropped. Remove an unused index by hand in the Firebase console.

## 7. Check the deployment

1. Run `curl -s "$(gcloud run services describe mtg-api --region us-central1 --format='value(status.url)')/readyz"`. It must read `"status":"ok"`.
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

CAUTION: a rollback of the code does not undo a change of the data. A new version can write a document that an old version cannot read. Read the store code of the merge before a rollback, and check `go/internal/store` for a schema change.

## 9. What a rollback does not repair

- A Firestore document that the new version wrote stays as it is.
- An index that a deploy added stays until you remove it by hand.
- A secret version stays active until you disable it.
- A card snapshot in the bucket stays. The worker keeps three versions (`cards.KeepVersions`).

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
| The Vite build needs Node 22 | `docs/setup.md` and the toolchain of the repo |
