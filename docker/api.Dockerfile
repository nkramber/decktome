# Pinned to the go/go.mod toolchain. Bump both together.
FROM golang:1.26.4 AS build
WORKDIR /src
# Copy the whole module before the build. `go build` fetches only the
# modules the api binary imports. A separate `go mod download` would also
# fetch every tool dependency (buf, golangci-lint) that the binary never uses.
COPY go/ go/
RUN cd go && CGO_ENABLED=0 go build -o /out/api ./cmd/api

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/api /api
EXPOSE 8080
ENTRYPOINT ["/api"]
