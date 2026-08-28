# Firestore + Auth emulators for the Compose stack.
# Node runs the firebase CLI. Java runs the Firestore emulator.
# Node matches .nvmrc. firebase-tools matches scripts/doctor.sh and docs/setup.md.
# Digest resolved from the registry manifest on 2026-08-28. Bump the tag
# and the digest together.
FROM node:26.7.0-slim@sha256:5758d367d7b4f48b73a9bb3530e687e47efb289f3b43f9c0450a25225ae0db5d
# Java 17 is the major that scripts/doctor.sh and docs/setup.md require.
RUN apt-get update \
    && apt-get install -y --no-install-recommends openjdk-17-jre-headless \
    && rm -rf /var/lib/apt/lists/* \
    && npm install -g firebase-tools@14.14.0
WORKDIR /app
COPY firebase.json firestore.rules firestore.indexes.json ./
# The host copy binds 127.0.0.1 for laptop use. Containers must bind 0.0.0.0
# so the published ports and other services can reach the emulators.
# The ui block has no host key, so the sed adds one. Without it the UI
# binds localhost and the published port 4100 gets no answer.
RUN sed -i 's/"host": "127.0.0.1"/"host": "0.0.0.0"/g' firebase.json \
    && sed -i 's/"ui": { "enabled": true,/"ui": { "enabled": true, "host": "0.0.0.0",/' firebase.json \
    && ! grep -q '127.0.0.1' firebase.json \
    && grep -q '"ui": { "enabled": true, "host": "0.0.0.0"' firebase.json
# /data/firestore is a Compose volume. The emulator imports it at start and
# exports to it at exit, the same as scripts/dev.sh does with .local/firestore.
# An empty directory is fine: the CLI warns and skips the import.
# The node image ships a `node` user (uid 1000). The emulators need no root.
RUN mkdir -p /data/firestore && chown -R node:node /data /app
USER node
EXPOSE 8281 9199 4100
CMD ["firebase", "emulators:start", "--only", "firestore,auth", "--project", "mtg-local", \
     "--import", "/data/firestore", "--export-on-exit", "/data/firestore"]
