FROM golang:1.26 AS build
WORKDIR /src
COPY go/go.mod go/go.sum go/
RUN cd go && go mod download
COPY go/ go/
RUN cd go && CGO_ENABLED=0 go build -o /out/api ./cmd/api

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/api /api
EXPOSE 8080
ENTRYPOINT ["/api"]
