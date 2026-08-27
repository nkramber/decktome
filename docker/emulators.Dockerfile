# Firestore + Auth emulators for the Compose stack.
# Node runs the firebase CLI. Java runs the Firestore emulator.
# Node matches .nvmrc. firebase-tools matches scripts/doctor.sh and docs/setup.md.
FROM node:22.23.2-slim
RUN apt-get update \
    && apt-get install -y --no-install-recommends default-jre-headless \
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
    && grep -q '"ui": { "enabled": true, "host": "0.0.0.0"' firebase.json
# /data/firestore is a Compose volume. The emulator imports it at start and
# exports to it at exit, the same as scripts/dev.sh does with .local/firestore.
# An empty directory is fine: the CLI warns and skips the import.
RUN mkdir -p /data/firestore
EXPOSE 8281 9199 4100
CMD ["firebase", "emulators:start", "--only", "firestore,auth", "--project", "mtg-local", \
     "--import", "/data/firestore", "--export-on-exit", "/data/firestore"]
