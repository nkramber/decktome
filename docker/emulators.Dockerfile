# Firestore + Auth emulators for the Compose stack.
# Node runs the firebase CLI. Java runs the Firestore emulator.
FROM node:20-slim
RUN apt-get update \
    && apt-get install -y --no-install-recommends default-jre-headless \
    && rm -rf /var/lib/apt/lists/* \
    && npm install -g firebase-tools@14.14.0
WORKDIR /app
COPY firebase.json firestore.rules firestore.indexes.json ./
# The host copy binds 127.0.0.1 for laptop use. Containers must bind 0.0.0.0
# so the published ports and other services can reach the emulators.
RUN sed -i 's/"host": "127.0.0.1"/"host": "0.0.0.0"/g' firebase.json
EXPOSE 8281 9199 4100
CMD ["firebase", "emulators:start", "--only", "firestore,auth", "--project", "mtg-local"]
