# Pinned to the go/go.mod toolchain. Bump both together.
# Digests resolved from the registry manifests on 2026-08-28. Bump the
# tag and the digest together.
FROM golang:1.26.6@sha256:0d1d3a794be25f809dd2cb3160d8c73276c4056a9f8242a138e908ddeee7b6b6 AS build
WORKDIR /src
# See docker/api.Dockerfile for why there is no separate `go mod download`.
COPY go/ go/
RUN cd go && CGO_ENABLED=0 go build -o /out/worker ./cmd/worker

# The nonroot tag runs as uid 65532.
FROM gcr.io/distroless/static-debian12:nonroot@sha256:afa5c872c891853ca7fcf1f12c3edb23f7eeef36189728842dd51042ff57f7ab
COPY --from=build /out/worker /worker
ENTRYPOINT ["/worker"]
