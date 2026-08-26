# Pinned to the go/go.mod toolchain. Bump both together.
# Digests resolved from the registry manifests on 2026-08-26. Bump the
# tag and the digest together.
FROM golang:1.26.4@sha256:f96cc555eb8db430159a3aa6797cd5bae561945b7b0fe7d0e284c63a3b291609 AS build
WORKDIR /src
# Copy the whole module before the build. `go build` fetches only the
# modules the api binary imports. A separate `go mod download` would also
# fetch every module in go.sum, and buf is one of them. golangci-lint is
# not in go.mod: `make lint-go` runs it with `go run ...@version`.
COPY go/ go/
RUN cd go && CGO_ENABLED=0 go build -o /out/api ./cmd/api
# The runtime image has no shell and no curl. This static probe calls
# GET /healthz for the Compose healthcheck. It exits 0 on HTTP 200.
RUN mkdir -p /probe && cd /probe && printf '%s\n' \
    'package main' \
    'import ("net/http"; "os")' \
    'func main() {' \
    '	r, err := http.Get("http://127.0.0.1:8080/healthz")' \
    '	if err != nil || r.StatusCode != http.StatusOK { os.Exit(1) }' \
    '}' > main.go && printf 'module probe\n\ngo 1.26\n' > go.mod \
    && CGO_ENABLED=0 go build -o /out/healthcheck .

# The nonroot tag runs as uid 65532. The api binds :8080, so no root port is needed.
FROM gcr.io/distroless/static-debian12:nonroot@sha256:afa5c872c891853ca7fcf1f12c3edb23f7eeef36189728842dd51042ff57f7ab
COPY --from=build /out/api /api
COPY --from=build /out/healthcheck /healthcheck
EXPOSE 8080
ENTRYPOINT ["/api"]
