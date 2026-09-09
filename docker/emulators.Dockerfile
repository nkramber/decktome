# Firestore + Auth emulators for the Compose stack.
# Node runs the firebase CLI. Java runs the Firestore emulator.
# Node matches .nvmrc. firebase-tools matches scripts/doctor.sh and docs/setup.md.
# Digest resolved from the registry manifest on 2026-08-28. Bump the tag
# and the digest together.
FROM node:22.23.2-slim@sha256:83f487e0a63425e5b4d146fb5e5be574bcbe1b7b843d3ebafdd95eaf7767a7e5
# default-jre-headless is the Debian release's own Java: 17 on bookworm,
# which is the release behind the node:22 digest above. The Firestore
# emulator needs Java 11 or newer, so the major moves with the image.
# Java 17 in scripts/doctor.sh is the laptop pin. The apt version is not
# pinned by choice: Debian removes an old package version from the
# archive, so a pinned build breaks at the next security release. The
# image digest fixes the apt index the build starts from.
# verified_at: 2026-08-29, source: `docker run --rm node:22.23.2-slim@<digest> cat /etc/os-release`
# and `apt-cache policy default-jre-headless` (candidate 2:1.17-74).
RUN apt-get update \
    && apt-get install -y --no-install-recommends default-jre-headless \
    && rm -rf /var/lib/apt/lists/* \
    && npm install -g firebase-tools@14.14.0
WORKDIR /app
COPY firebase.json firestore.rules firestore.indexes.json ./
# The host copy binds 127.0.0.1 for laptop use. Containers must bind
# 0.0.0.0 so the published ports and other services can reach the
# emulators. The ui block carries no host key, so this adds one. Without
# it the UI binds localhost and the published port 4100 gets no answer.
#
# Node reads the file as JSON and writes it back. A sed read the shape of
# the text instead, and it broke the moment the file took one key per
# line: the ui edit matched nothing and the build failed on its own
# check (D-641).
RUN node -e "const f='firebase.json',fs=require('fs'); \
    const j=JSON.parse(fs.readFileSync(f,'utf8')); \
    for (const e of Object.values(j.emulators||{})) if (e && typeof e==='object') e.host='0.0.0.0'; \
    fs.writeFileSync(f, JSON.stringify(j,null,2));" \
    && ! grep -q '127.0.0.1' firebase.json \
    && node -e "const j=require('./firebase.json'); \
    for (const [k,e] of Object.entries(j.emulators||{})) { \
      if (e && typeof e==='object' && e.host!=='0.0.0.0') { \
        console.error('emulator '+k+' binds '+e.host); process.exit(1); } }"
# /data/firestore is a Compose volume. The emulator imports it at start and
# exports to it at exit, the same as scripts/dev.sh does with .local/firestore.
# An empty directory is fine: the CLI warns and skips the import.
# The node image ships a `node` user (uid 1000). The emulators need no root.
RUN mkdir -p /data/firestore && chown -R node:node /data /app
USER node
EXPOSE 8281 9199 4100
CMD ["firebase", "emulators:start", "--only", "firestore,auth", "--project", "mtg-local", \
     "--import", "/data/firestore", "--export-on-exit", "/data/firestore"]
