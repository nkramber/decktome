# Pinned to the go/go.mod toolchain. Bump both together.
FROM golang:1.26.4 AS build
WORKDIR /src
# See docker/api.Dockerfile for why there is no separate `go mod download`.
COPY go/ go/
RUN cd go && CGO_ENABLED=0 go build -o /out/worker ./cmd/worker

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/worker /worker
ENTRYPOINT ["/worker"]
