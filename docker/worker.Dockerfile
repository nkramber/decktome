# Pinned to the go line of go/go.mod. make go-version-check holds the
# two together (D-1216). The golang digest was resolved from the registry
# index on 2026-10-08. Bump the tag and the digest together.
FROM golang:1.27.2@sha256:5bc7f572bbaa98885a3a1fd9c0aa76b59e3e14e8628bfc316bbfd0c701e4818c AS build
WORKDIR /src
# See docker/api.Dockerfile for why there is no separate `go mod download`.
COPY go/ go/
RUN cd go && CGO_ENABLED=0 go build -o /out/worker ./cmd/worker

# The nonroot tag runs as uid 65532.
FROM gcr.io/distroless/static-debian12:nonroot@sha256:afa5c872c891853ca7fcf1f12c3edb23f7eeef36189728842dd51042ff57f7ab
COPY --from=build /out/worker /worker
ENTRYPOINT ["/worker"]
