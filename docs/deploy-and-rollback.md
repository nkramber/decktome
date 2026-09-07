# Deploy after a merge, and roll back

`docs/setup-gcp.md` builds the deployment one time. This page runs after every merge to `main`. It also tells you how to return to the version before a bad deploy.

The project is `decktome-prod`, and the region is `us-central1`. The deploy of 2026-09-07 set both. Write in the commands: `TAG` is the short commit of the merge.

## 1. A merge to main deploys itself

The workflow `deploy` builds and releases every merge to `main` (D-582). No step of section 4 or section 5 runs by hand any more. The workflow reads the changed paths first, so a merge of documents alone deploys nothing.

| The merge touches | The workflow does |
|---|---|
| Any file under `go/` or `docker/` | Builds both images, deploys the API, points both jobs at the new image, and waits for `/readyz` |
| Any file under `web/`, or `firebase.json` | Builds the web app and releases it, then reads the site |
| Only `docs/` | Nothing |

Sections 2 to 6 stay for two jobs: the first deploy of a new project, and a step the workflow cannot do. Section 8 is the rollback, and it stays a command you run.

CAUTION: the workflow deploys `main` and no other branch (D-579). It starts on a push to `main` and on a dispatch, and never on a pull request (D-578).

### 1.1 The one-time setup

The workflow signs in with Workload Identity Federation, so no key of a service account lives in the repository. Run these commands one time, as the owner of the project.

1. Create the deployer:

```
gcloud iam service-accounts create gh-deployer --display-name="github deployer"
```

2. Create the pool and the provider. The condition names the repository, so no other repository can sign in as the deployer:

```
gcloud iam workload-identity-pools create github --location=global \
  --display-name="GitHub Actions"
gcloud iam workload-identity-pools providers create-oidc github \
  --location=global --workload-identity-pool=github \
  --issuer-uri="https://token.actions.githubusercontent.com" \
  --attribute-mapping="google.subject=assertion.sub,attribute.repository=assertion.repository" \
  --attribute-condition="assertion.repository=='nkramber/decktome'"
```

3. Let the repository act as the deployer:

```
gcloud iam service-accounts add-iam-policy-binding \
  gh-deployer@decktome-prod.iam.gserviceaccount.com \
  --role=roles/iam.workloadIdentityUser \
  --member="principalSet://iam.googleapis.com/projects/492774632746/locations/global/workloadIdentityPools/github/attribute.repository/nkramber/decktome"
```

4. Grant the deployer what a deploy needs, and nothing else:

```
for role in roles/run.admin roles/artifactregistry.writer roles/firebasehosting.admin; do
  gcloud projects add-iam-policy-binding decktome-prod \
    --member=serviceAccount:gh-deployer@decktome-prod.iam.gserviceaccount.com --role=$role
done
```

5. Let the deployer run the API and the jobs as their own accounts. This binding names those two accounts, and never every account of the project:

```
for sa in mtg-api mtg-worker; do
  gcloud iam service-accounts add-iam-policy-binding \
    $sa@decktome-prod.iam.gserviceaccount.com \
    --member=serviceAccount:gh-deployer@decktome-prod.iam.gserviceaccount.com \
    --role=roles/iam.serviceAccountUser
done
```

6. Write the four Firebase values as repository variables. They reach the browser in the bundle, so they are variables and never secrets:

```
gh variable set VITE_FIREBASE_PROJECT_ID --body "decktome-prod"
gh variable set VITE_FIREBASE_AUTH_DOMAIN --body "decktome-prod.firebaseapp.com"
gh variable set VITE_FIREBASE_API_KEY --body "<the apiKey of firebase apps:sdkconfig>"
gh variable set VITE_FIREBASE_APP_ID --body "<the appId of firebase apps:sdkconfig>"
```

Note: one deploy costs about 5 to 8 minutes of Actions time for the API and about 3 to 4 for the web app. A merge that touches documents alone costs about one minute.

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
