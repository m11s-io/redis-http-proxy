# syntax=docker/dockerfile:1.28.0@sha256:bb22d9815c728170f72750f4e5b0d672e06176142e1d602c7e66c050100b7e5b

FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o redis-http-proxy .

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /app/redis-http-proxy /redis-http-proxy
EXPOSE 80
ENTRYPOINT ["/redis-http-proxy"]
