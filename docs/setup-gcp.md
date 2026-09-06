# Set up the app on Google Cloud

This page tells you how to put the app on the internet for invited users. It starts at the domain and ends at a deployed web app, an API, and two scheduled jobs. It also estimates the monthly cost for five active users who build three decks a week each.

The page reads the repo as it stands on 2026-09-05. PR-22 is the deploy slice of the roadmap, and it waits for its turn. Section 1 names the code that PR-22 must add before the last steps work. Every price on this page is a list price read on 2026-09-05, and section 17 names the source and the date of each one.

CAUTION: the official pricing pages of Cloud Run, Firestore, and Cloud Storage render in a browser only. Their numbers come from the Google Cloud free-tier document and from two dated third-party reads. Confirm each number in the Google Cloud pricing calculator before an invoice matters.

Write in the commands: `PROJECT_ID` is your project id, `REGION` is `us-central1`, and `DOMAIN` is your domain, for example `decks.example.com`. Section 3 explains the region.

## 1. What the deployment holds

The roadmap (PR-22, D-310, D-314) fixes the shape. One table names each part.

| Part | Google Cloud product | What it does |
|---|---|---|
| The API | Cloud Run service, `docker/api.Dockerfile` | Serves the Connect RPCs, verifies the Firebase ID token, reads the allowlist, calls the model providers |
| The snapshot job | Cloud Run job, `docker/worker.Dockerfile`, `-once` | Refreshes the Scryfall card snapshot in the bucket every 15 minutes (D-61) |
| The meta job | Cloud Run job, `docker/worker.Dockerfile`, `-meta` | Reads the deck list sources and fits the quality model daily at 06:00 UTC (D-492) |
| The database | Firestore, Native mode, Standard edition | Sessions, decks, collections, usage, and the allowlist document `config/allowlist` (D-420) |
| The bucket | Cloud Storage, `PROJECT_ID-cards` | The card snapshots, three versions kept (`cards.KeepVersions`), and the meta store |
| The secrets | Secret Manager | `OPENAI_API_KEY`, `ANTHROPIC_API_KEY`, and `TOPDECK_API_KEY` (D-492) |
| Sign-in | Firebase Authentication, email and password | Open sign-up, and the API refuses every email that is not on the allowlist (D-314) |
| The web app | Firebase Hosting on your domain | The Vite build of `web/apps/web`, with a free managed certificate |
| The images | Artifact Registry | The two container images |
| The schedules | Cloud Scheduler, two jobs | Runs the two Cloud Run jobs |
| The alarm | Cloud Billing budget | Sends an email at a spend threshold |

The repo holds the two Dockerfiles and the Firestore rules and indexes. The Go code reads `PROJECT_ID`, `CARDS_BUCKET`, `ALLOWED_ORIGINS`, and the three keys from the environment. `gcpenv.OnCloudRun` reads `K_SERVICE`, which Cloud Run sets, and the API then refuses every request with no token (`cmd/api/main.go`).

Four things are not in the repo yet, and PR-22 adds them:

- A `hosting` block in `firebase.json` for the static app (D-544). The file holds the Firestore and emulator blocks alone.
- The real Firebase web configuration. `web/apps/web/src/lib/firebase.ts` initializes the app with `apiKey: "demo-key"` and `projectId: "mtg-local"`. A production build needs the values of section 5 through environment variables.
- The allowlist interceptor and `make allow EMAIL=...` (D-314, D-420).
- The per-user spend cap of $5 a month (D-421).

## 2. Before you start

You need three accounts and their credentials:

- A Google account for the Google Cloud console and the Firebase console.
- A payment method. A Cloud Billing account needs one, also during the free trial.
- The provider keys of `.env`: `OPENAI_API_KEY`, `ANTHROPIC_API_KEY`, and `TOPDECK_API_KEY`.

You need these tools on the Mac. `docs/setup.md` installs the first four.

1. Go 1.27.0, Node 22.23.2, pnpm 9.2.0, and firebase-tools 14.14.0.
2. The gcloud CLI. Run `brew install --cask gcloud-cli`, then `gcloud init`.
3. Docker Desktop, for the image builds of section 10. Cloud Build is the alternative there.

Choose three names before the first command:

- `PROJECT_ID`: lower case, 6 to 30 characters, unique across Google Cloud. It can not change later.
- `REGION`: this page uses `us-central1`. It is a Tier 1 price region, and it is one of the three regions of the Cloud Storage free tier. Every Cloud Run product is available there.
- `DOMAIN`: the name the users type.

## 3. Buy the domain

Google Domains closed to new registrations on 2023-09-07, when Squarespace bought it. Three ways remain, and each one works with the steps of this page.

| Registrar | Price of a `.com` | Notes |
|---|---|---|
| Cloud Domains, inside Google Cloud | The console and the CLI show the yearly price before you confirm | Squarespace is the registrar of record, and Google bills your Cloud Billing account. Renewal is automatic. You must choose a DNS host at registration: a Cloud DNS zone or custom name servers. |
| Squarespace Domains | $12 the first year, $20 a year after that (third-party read, 2026-09-05) | WHOIS privacy included. DNS included. |
| Cloudflare Registrar | About $10 to $12 a year, at the registry cost (third-party read, 2026-09-05) | The lowest renewal. DNS included. Set the Firebase records to DNS only, not proxied. |

Choose Cloud Domains when you want one invoice. Choose Cloudflare when you want the lowest renewal. Do not choose Google Domains DNS: Cloud Domains registrations can not use it since 2023-10-19.

### 3.1 Cloud Domains

1. Create the project and the billing account first (section 4). Cloud Domains needs both.
2. Run `gcloud services enable domains.googleapis.com`.
3. Run `gcloud domains registrations search-domains NAME`. It lists the free names near `NAME`.
4. Run `gcloud domains registrations get-register-parameters DOMAIN`. It prints the availability, the yearly price, and the privacy modes.
5. If you want Cloud DNS, run `gcloud dns managed-zones create mtg-zone --description="mtg deck builder" --dns-name=DOMAIN.`. Note the trailing dot.
6. Run `gcloud domains registrations register DOMAIN`. It asks for the DNS host, the contact data, and the privacy mode.
7. Confirm the yearly price when the command shows it.

CAUTION: Cloud DNS bills per zone and per query. This page did not verify those two prices on 2026-09-05. The registrar's own DNS costs nothing, and this page uses it.

### 3.2 Squarespace or Cloudflare

1. Buy the domain on the registrar's site.
2. Keep the DNS at the registrar. Section 14 adds the Firebase records there.

## 4. Create the project and the billing account

1. Open https://console.cloud.google.com and sign in.
2. Create the billing account. The Free Trial gives "$300 Welcome credit to spend over 90 days", and "you must provide a credit card or other payment method". Google does not bill usage during the trial. At the end of the 90 days the trial account closes, and its projects stop. The same happens when the credit is gone. Upgrade to a paid account before you invite a user.
3. Create the project: `gcloud projects create PROJECT_ID --name="mtg deck builder"`.
4. Link the billing account: `gcloud billing accounts list`, then `gcloud billing projects link PROJECT_ID --billing-account=ACCOUNT_ID`.
5. Run `gcloud config set project PROJECT_ID`.
6. Enable the APIs in one command:

```
gcloud services enable run.googleapis.com artifactregistry.googleapis.com \
  firestore.googleapis.com secretmanager.googleapis.com cloudscheduler.googleapis.com \
  cloudbuild.googleapis.com iam.googleapis.com firebase.googleapis.com \
  identitytoolkit.googleapis.com storage.googleapis.com
```

7. Create the budget alert. Open Billing, then Budgets & alerts, then Create budget. Set the scope to the project. Set the amount to $30 a month. Add thresholds at 50, 90, and 100 percent, on actual spend, and one at 100 percent on forecasted spend. Send the emails to the billing administrators.

CAUTION: a budget sends emails and stops nothing. Google writes: "Setting an alerts-only budget doesn't automatically cap Google Cloud or Google Maps Platform usage or spending." The per-user cap of D-421 is the control that stops spend, and it lives in the API.

## 5. Add Firebase to the project

1. Run `firebase login`.
2. Run `firebase projects:addfirebase PROJECT_ID`. It adds the Firebase resources to the project you made.
3. In the repo root, run `firebase use --add`, choose `PROJECT_ID`, and name the alias `prod`.
4. Run `firebase apps:create WEB "mtg deck builder"`. Note the app id it prints.
5. Run `firebase apps:sdkconfig WEB APP_ID`. It prints the web configuration: `apiKey`, `authDomain`, `projectId`, and `appId`. PR-22 reads these through `VITE_` variables at build time.
6. Open https://console.firebase.google.com, choose the project, then Authentication, then Sign-in method. Enable Email/Password.
7. In Authentication, open Settings, then Authorized domains, and add `DOMAIN`.

The project moved from the Spark plan to the Blaze plan when you linked the billing account. The Hosting rewrites to Cloud Run need Blaze. Authentication stays free to 50,000 monthly active users on both plans.

## 6. Create the Firestore database

1. Run `gcloud firestore databases create --location=REGION --type=firestore-native --edition=standard`. The database id stays `(default)`.
2. Run `firebase deploy --only firestore`. It writes the deny-all rules and the indexes of the repo.

The rules deny every client read and write. The Go services use the Admin SDK, which bypasses the rules, so the API is the only door. Firestore allows exactly one free database per project, and the first one you create is it.

## 7. Create the bucket

1. Run `gcloud storage buckets create gs://PROJECT_ID-cards --location=REGION --uniform-bucket-level-access`.

The name matters. `gcpenv` reads `CARDS_BUCKET`, and its default is `PROJECT_ID-cards`. The bucket holds the snapshots under `scryfall/` and the meta store under `meta/`. The worker keeps three complete snapshot versions, and the version of 2026-09-04 holds 109 MB on the owner's disk. The meta store holds 263 MB. No object is public.

## 8. Store the secrets

Run these six commands. Each key comes from `.env`.

```
gcloud secrets create openai-api-key --replication-policy=automatic
printf '%s' "$OPENAI_API_KEY" | gcloud secrets versions add openai-api-key --data-file=-
gcloud secrets create anthropic-api-key --replication-policy=automatic
printf '%s' "$ANTHROPIC_API_KEY" | gcloud secrets versions add anthropic-api-key --data-file=-
gcloud secrets create topdeck-api-key --replication-policy=automatic
printf '%s' "$TOPDECK_API_KEY" | gcloud secrets versions add topdeck-api-key --data-file=-
```

Three secrets with one version each stay inside the free tier of six active versions. Each version is `1`, and section 11 pins that number. Google recommends a pinned version over `latest` for a secret in an environment variable, because Cloud Run reads it at instance start.

## 9. Create the service accounts

Three service accounts keep the permissions apart.

```
gcloud iam service-accounts create mtg-api --display-name="mtg api"
gcloud iam service-accounts create mtg-worker --display-name="mtg worker"
gcloud iam service-accounts create mtg-scheduler --display-name="mtg scheduler"
```

Grant the roles. Replace `SA_API`, `SA_WORKER`, and `SA_SCHEDULER` with the three emails, of the form `NAME@PROJECT_ID.iam.gserviceaccount.com`.

```
gcloud projects add-iam-policy-binding PROJECT_ID --member=serviceAccount:SA_API --role=roles/datastore.user
gcloud projects add-iam-policy-binding PROJECT_ID --member=serviceAccount:SA_WORKER --role=roles/datastore.user
gcloud storage buckets add-iam-policy-binding gs://PROJECT_ID-cards --member=serviceAccount:SA_API --role=roles/storage.objectViewer
gcloud storage buckets add-iam-policy-binding gs://PROJECT_ID-cards --member=serviceAccount:SA_WORKER --role=roles/storage.objectAdmin
gcloud secrets add-iam-policy-binding openai-api-key --member=serviceAccount:SA_API --role=roles/secretmanager.secretAccessor
gcloud secrets add-iam-policy-binding anthropic-api-key --member=serviceAccount:SA_API --role=roles/secretmanager.secretAccessor
gcloud secrets add-iam-policy-binding topdeck-api-key --member=serviceAccount:SA_WORKER --role=roles/secretmanager.secretAccessor
```

The API verifies Firebase ID tokens with Google's public keys, so it needs no Firebase role. The scheduler account gets the Cloud Run Invoker role on each job in section 12.

## 10. Build and push the images

1. Run `gcloud artifacts repositories create mtg --repository-format=docker --location=REGION`.
2. Run `gcloud auth configure-docker REGION-docker.pkg.dev`.
3. Build both images from the repo root. The Dockerfiles copy the `go/` directory, so the root is the context. Cloud Run runs `linux/amd64` images only, so an Apple Silicon Mac must name the platform.

```
TAG=$(git rev-parse --short HEAD)
docker build --platform linux/amd64 -f docker/api.Dockerfile -t REGION-docker.pkg.dev/PROJECT_ID/mtg/api:$TAG .
docker build --platform linux/amd64 -f docker/worker.Dockerfile -t REGION-docker.pkg.dev/PROJECT_ID/mtg/worker:$TAG .
docker push REGION-docker.pkg.dev/PROJECT_ID/mtg/api:$TAG
docker push REGION-docker.pkg.dev/PROJECT_ID/mtg/worker:$TAG
```

Cloud Build is the alternative when the Mac has no Docker. It needs a `cloudbuild.yaml` that names the two Dockerfiles, which the repo does not hold yet. Each billing account gets 2,500 free build-minutes a month on the default pool, and an `e2-standard-2` minute costs $0.006 in `us-central1` after that, since 2025-11-01.

## 11. Deploy the API

Run one command. Replace `DOMAIN` in `ALLOWED_ORIGINS` with the origin the browser uses, with the scheme and no path.

```
gcloud run deploy mtg-api \
  --image REGION-docker.pkg.dev/PROJECT_ID/mtg/api:$TAG \
  --region REGION --platform managed \
  --service-account SA_API \
  --allow-unauthenticated \
  --set-env-vars PROJECT_ID=PROJECT_ID,CARDS_BUCKET=PROJECT_ID-cards,ALLOWED_ORIGINS=https://DOMAIN \
  --set-secrets OPENAI_API_KEY=openai-api-key:1,ANTHROPIC_API_KEY=anthropic-api-key:1 \
  --memory 2Gi --cpu 1 --min-instances 0 --max-instances 3 --concurrency 20 \
  --timeout 900 --port 8080
```

Six notes on the flags:

- `--allow-unauthenticated` opens the URL to the internet. The API checks the Firebase token itself on every request, and a request with no token gets Unauthenticated on Cloud Run. Do not set `ALLOW_DEBUG_USER`.
- `--min-instances 0` bills nothing at idle. The first request after an idle period starts an instance, and the instance loads the newest snapshot from the bucket. Expect a cold start of some seconds.
- `--memory 2Gi` is a starting point. The API holds the whole card index in memory. Read the memory chart after the first week and move the number.
- `--timeout 900` covers a deck build. The default is 300 seconds and the maximum is 3,600. A build with repair passes takes minutes, and the `Chat` RPC streams for that whole time.
- `--set-secrets` pins version `1`. Rotate a key with a new version and a new deploy.
- `PROJECT_ID` must be explicit. Cloud Run sets `K_SERVICE` and not the project id.

Check the service: `curl -s "$(gcloud run services describe mtg-api --region REGION --format='value(status.url)')/healthz"`. It answers 200 once a snapshot exists in the bucket, so run section 12 first when the bucket is empty.

## 12. Deploy the jobs and the schedules

Create the two jobs.

```
gcloud run jobs create mtg-snapshot \
  --image REGION-docker.pkg.dev/PROJECT_ID/mtg/worker:$TAG --args=-once \
  --region REGION --service-account SA_WORKER \
  --set-env-vars PROJECT_ID=PROJECT_ID,CARDS_BUCKET=PROJECT_ID-cards \
  --memory 1Gi --cpu 1 --task-timeout 30m --max-retries 0
gcloud run jobs create mtg-meta \
  --image REGION-docker.pkg.dev/PROJECT_ID/mtg/worker:$TAG --args=-meta \
  --region REGION --service-account SA_WORKER \
  --set-env-vars PROJECT_ID=PROJECT_ID,CARDS_BUCKET=PROJECT_ID-cards \
  --set-secrets TOPDECK_API_KEY=topdeck-api-key:1 \
  --memory 1Gi --cpu 1 --task-timeout 60m --max-retries 0
```

Run the first snapshot by hand and wait for it: `gcloud run jobs execute mtg-snapshot --region REGION --wait`. The largest bulk file is about 80 MB, and the download timeout inside the worker is 25 minutes. Then run `gcloud run jobs execute mtg-meta --region REGION --wait`. The first meta run reads every source and takes about 40 minutes.

Give the scheduler account the invoker role on both jobs:

```
gcloud run jobs add-iam-policy-binding mtg-snapshot --region REGION --member=serviceAccount:SA_SCHEDULER --role=roles/run.invoker
gcloud run jobs add-iam-policy-binding mtg-meta --region REGION --member=serviceAccount:SA_SCHEDULER --role=roles/run.invoker
```

Create the two schedules. D-61 sets the snapshot tick at 15 minutes, and D-492 sets the meta run at 06:00 UTC.

```
gcloud scheduler jobs create http mtg-snapshot-schedule --location REGION \
  --schedule="*/15 * * * *" --time-zone="Etc/UTC" --http-method POST \
  --uri="https://run.googleapis.com/v2/projects/PROJECT_ID/locations/REGION/jobs/mtg-snapshot:run" \
  --oauth-service-account-email SA_SCHEDULER
gcloud scheduler jobs create http mtg-meta-schedule --location REGION \
  --schedule="0 6 * * *" --time-zone="Etc/UTC" --http-method POST \
  --uri="https://run.googleapis.com/v2/projects/PROJECT_ID/locations/REGION/jobs/mtg-meta:run" \
  --oauth-service-account-email SA_SCHEDULER
```

Two jobs stay inside the free tier of three Cloud Scheduler jobs per billing account. The snapshot job exits early on most ticks, because the newest snapshot is under an hour old (D-61).

## 13. Build and deploy the web app

This section needs the two PR-22 additions of section 1: the `hosting` block and the Firebase configuration variables.

1. Build the web app with the API origin and the Firebase configuration. The API origin is the Cloud Run URL of section 11 until section 14 gives it a name.

```
cd web
VITE_API_BASE_URL=https://mtg-api-XXXX-uc.a.run.app VITE_AUTH_EMULATOR_HOST= pnpm --filter web build
cd ..
```

2. Add the `hosting` block to `firebase.json`. This example serves the Vite build as a single-page app.

```
"hosting": {
  "public": "web/apps/web/dist",
  "ignore": ["firebase.json", "**/.*"],
  "rewrites": [
    { "source": "**", "destination": "/index.html" }
  ]
}
```

3. Run `firebase deploy --only hosting`. The site is live on `PROJECT_ID.web.app`.

The web app calls the Cloud Run origin directly (D-544). The API reads `ALLOWED_ORIGINS` for CORS, and the web app reads `VITE_API_BASE_URL`, so step 1 is the whole wiring. No Hosting rewrite carries the RPCs. The reason: Firebase documents a 60-second request timeout for rewrites to Cloud Functions, "Firebase Hosting is subject to a 60-second request timeout". The Cloud Run rewrite page makes no statement about a timeout or about streamed responses. A deck build streams for several minutes over the `Chat` RPC, and no proxy sits in that path.

## 14. Connect the domain

1. Open the Firebase console, then Hosting, then Add custom domain. Type `DOMAIN`.
2. Firebase shows a TXT record. Add it at your DNS host. Keep it there: Firebase reads it again later to prove ownership.
3. Firebase shows the A and AAAA records. Add them at your DNS host. Remove every other A, AAAA, or CNAME record of the name. On Cloudflare, set the records to DNS only.
4. Wait. Google writes: "It may take up to 24 hours after you point your DNS to Firebase Hosting." Most certificates arrive within a few hours.
5. Rebuild the web app with `ALLOWED_ORIGINS=https://DOMAIN` on the API and the same origin in the Firebase Authentication authorized domains (section 5, step 7).

The API keeps its `run.app` URL. Cloud Run domain mappings are a preview feature, and Google writes that they "are not recommended for production services". A global external Application Load Balancer gives the API a name with a managed certificate. It bills by the hour, so this page leaves it out at this scale.

## 15. Invite a user and check the deployment

1. Run `make allow EMAIL=user@example.com`. PR-22 adds the target, and it writes the email into `config/allowlist` (D-420).
2. Open `https://DOMAIN`, create the account with that email, and sign in.
3. Upload a ManaBox export, build a deck, revise it, and export it. This is the PR-22 gate.
4. Sign in with an email that is not on the list. The first RPC must answer `PermissionDenied` with one sentence.
5. Read the Cloud Run logs: `gcloud run services logs read mtg-api --region REGION --limit 50`.

## 16. Cost estimate for five users and three decks a week each

### 16.1 The load

Five users who build three decks a week make 15 decks a week, which is about 65 decks a month. This page assumes one deck holds four question turns, one build with its repair passes, and one revision turn.

### 16.2 The model calls

The provider prices come from `go/internal/llm/prices.json`, verified on 2026-08-24. The roles come from `roles.json`: classify and ask on `gpt-5.6-luna`, generate, repair, and revise on `gpt-5.6-terra`. The judge role runs in the gates alone, and no production build calls it.

| Model | Input, per 1M tokens | Cached input | Output |
|---|---|---|---|
| `gpt-5.6-luna` | $0.20 | $0.02 | $1.20 |
| `gpt-5.6-terra` | $2.00 | $0.20 | $12.00 |

The measured run costs of the gates give the price of each step.

| Step | Measured on | Cost per unit |
|---|---|---|
| One question turn | Question gate runs 38 to 41, 2026-09-05: $0.14 to $0.19 for 108 conversations of 2 to 4 turns | Under $0.002 a turn |
| One build with repairs | Deck gate runs 12 to 14, 2026-09-02 to 2026-09-03: $2.24 to $2.64 for 24 decks | $0.09 to $0.11 a deck |
| One revision turn | Revise gate run 9, 2026-09-04: $1.07 for 11 turns | About $0.10 a turn |

One deck with one revision costs about $0.22. The month costs $8 with no revisions, $14 with one revision a deck, and $20 with two. The per-user cap of D-421 bounds the whole at $25 a month for five users.

### 16.3 Google Cloud

Every line reads the free tier of the Google Cloud free program document, 2026-09-05, against the load above. The paid prices are the fallback when a free tier lapses.

| Product | Load a month | Free tier a month | Cost | Paid price beyond the tier |
|---|---|---|---|---|
| Cloud Run, the API | 65 builds of about 4 busy minutes at 1 vCPU and 2 GiB: 15,600 vCPU-seconds and 31,200 GiB-seconds, plus small requests | 180,000 vCPU-seconds, 360,000 GiB-seconds, 2 million requests | $0 | $0.000024 a vCPU-second, $0.0000025 a GiB-second, $0.40 a million requests |
| Cloud Run, the two jobs | 2,880 snapshot ticks of about 10 seconds plus one download a day, and 30 meta runs of about 10 minutes: about 52,000 vCPU-seconds | 240,000 vCPU-seconds and 450,000 GiB-seconds for instance-based billing | $0 | $0.000018 a vCPU-second, $0.000002 a GiB-second |
| Firestore | About 13,000 writes and 50,000 reads a month, under 100 MiB stored | 20,000 writes, 50,000 reads, and 20,000 deletes a day, 1 GiB stored | $0 | $0.09 a 100,000 writes, $0.03 a 100,000 reads, $0.15 a GiB a month in a US region |
| Cloud Storage | About 0.6 GB: three snapshots of 109 MB and a 263 MB meta store, with a few thousand operations | 5 GB-months in `us-central1`, 5,000 Class A and 50,000 Class B operations | $0 | $0.020 a GB a month, $0.05 a 10,000 Class A, $0.004 a 10,000 Class B |
| Secret Manager | 3 active versions, a few hundred accesses at instance starts | 6 active versions, 10,000 accesses | $0 | $0.06 a version a month, $0.03 a 10,000 accesses |
| Cloud Scheduler | 2 jobs | 3 jobs a billing account | $0 | $0.10 a job a month |
| Artifact Registry | Two images of some tens of MB each, a few tags | 0.5 GB | $0 | $0.10 a GB a month |
| Cloud Build | 0 minutes with local Docker builds | 2,500 build-minutes | $0 | $0.006 a minute |
| Cloud Logging | Well under 1 GiB | 50 GiB a project | $0 | $0.50 a GiB |
| Firebase Hosting | A 2 MB bundle, five users | 10 GB stored, 360 MB a day transferred | $0 | $0.026 a GB stored, $0.15 a GB transferred |
| Firebase Authentication | 5 monthly active users | 50,000 monthly active users | $0 | Google Cloud pricing above the tier |
| Network egress | The model calls move about 0.2 GB | 1 GB from North America for Cloud Run | $0 | $0.12 a GB for the first TiB |
| The domain | One `.com` | None | $0.85 to $1.70 | $10 to $20 a year |

### 16.4 The total

| Line | A month |
|---|---|
| Model calls, 65 decks | $8 to $20 |
| Google Cloud, inside the free tiers | $0 |
| The domain | $1 to $2 |
| Total | $9 to $22 |

An idle month costs the domain alone. Three things move the number. A user who revises a deck five times spends five revision turns, which the $5 cap of D-421 stops. The Always Free tier is per billing account, so a second project on the same account shares it. The free trial credit does not change the estimate: the credit pays the first $300 of whatever the meter reads.

CAUTION: the PR-22 gate records the measured monthly cost at idle. Replace this estimate with that measurement when it exists.

## 17. Sources and dates

Every fact of this page carries a date. The repo facts read the code and the documents of 2026-09-05. The web facts read these pages on 2026-09-05.

| Fact | Source |
|---|---|
| Google Domains sold to Squarespace on 2023-09-07, and the migration finished on 2024-07-10 | https://newsroom.squarespace.com/blog/squarespace-domains-updates and https://docs.cloud.google.com/domains/docs/faq |
| Cloud Domains still registers domains, Squarespace is the registrar of record, and Google bills | https://docs.cloud.google.com/domains/docs/faq |
| Google Domains DNS closed to new registrations on 2023-10-19 | https://docs.cloud.google.com/domains/docs/deprecations/feature-deprecations |
| The Cloud Domains commands | https://docs.cloud.google.com/domains/docs/register-domain |
| Squarespace `.com` at $12 the first year and $20 a year after | https://www.stackscored.com/pricing/domain-registrars/squarespace-domains/ (third party) |
| Cloudflare Registrar at cost, about $10 to $12 a year | https://www.stackscored.com/pricing/domain-registrars/compare/cloudflare-registrar-vs-squarespace-domains/ (third party) |
| The Free Trial terms and every Always Free amount of section 16 | https://docs.cloud.google.com/free/docs/free-cloud-features |
| Cloud Run request-based and instance-based prices | https://preprice.app/ai-costs/gcp_cloud_run (third party, verified 2026-06-14) and https://cloudchipr.com/blog/cloud-run-pricing (third party, 2025-11-14). The official page is https://cloud.google.com/run/pricing. |
| Cloud Run request timeout, default 300 and maximum 3,600 seconds | https://docs.cloud.google.com/run/docs/configuring/request-timeout |
| Cloud Run secrets as environment variables and the pinned version advice | https://docs.cloud.google.com/run/docs/configuring/services/secrets |
| Cloud Run jobs on a schedule, the invoker role, and the run URI | https://docs.cloud.google.com/run/docs/execute/jobs-on-schedule |
| The `gcloud run jobs create` flags | https://docs.cloud.google.com/sdk/gcloud/reference/run/jobs/create |
| The `gcloud firestore databases create` flags | https://docs.cloud.google.com/sdk/gcloud/reference/firestore/databases/create |
| Firestore free quota and the one free database a project | https://firebase.google.com/pricing and https://firebase.google.com/docs/firestore/pricing |
| Firestore US region prices beyond the quota | https://airbyte.com/data-engineering-resources/google-firestore-pricing (third party, 2025-05-05). The official page is https://cloud.google.com/firestore/pricing. |
| Cloud Storage US region prices | https://www.nops.io/blog/google-cloud-storage-pricing/ (third party, 2026-09-01). The official page is https://cloud.google.com/storage/pricing. |
| Secret Manager prices | https://cloud.google.com/secret-manager/pricing, read through search on 2026-09-05 |
| Cloud Scheduler, Artifact Registry, and Cloud Logging prices | https://cloud.google.com/scheduler/pricing, https://cloud.google.com/artifact-registry/pricing, and https://cloud.google.com/logging, read through search on 2026-09-05 |
| Cloud Tasks prices | https://cloud.google.com/tasks/pricing |
| Cloud Build free minutes and the price since 2025-11-01 | https://cloud.google.com/build/pricing and https://cloud.google.com/build/pricing-update |
| Cloud Run egress, 1 GB free and $0.12 a GB | https://egresscost.com/gcp/ (third party) and https://cloud.google.com/run/pricing |
| Firebase Hosting and Authentication prices | https://firebase.google.com/pricing |
| Hosting rewrites to Cloud Run: Blaze, regions, the syntax | https://firebase.google.com/docs/hosting/cloud-run |
| The 60-second Hosting timeout for Cloud Functions rewrites | https://firebase.google.com/docs/hosting/functions |
| Custom domains on Hosting: the TXT record, the A and AAAA records, 24 hours | https://firebase.google.com/docs/hosting/custom-domain |
| Cloud Run custom domain options and the preview status of domain mappings | https://docs.cloud.google.com/run/docs/mapping-custom-domains |
| Budgets send alerts and cap nothing | https://docs.cloud.google.com/billing/docs/how-to/budgets |
| Firebase CLI command forms | `firebase --help` of firebase-tools 14.14.0 on the owner's Mac |
| The model prices and roles | `go/internal/llm/prices.json` and `roles.json`, verified 2026-08-24 |
| The measured gate costs | `docs/reference/pr7-question-gate-run38.md` to `run41.md`, `pr8-deck-gate-run12.md` to `run14.md`, `pr12b-revise-gate-run9.md` |
| The snapshot and meta store sizes, and the three kept versions | `.local/gcs/mtg-local-cards` on 2026-09-05 and `go/internal/cards/refresh.go` |
