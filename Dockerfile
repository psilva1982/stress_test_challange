FROM golang:alpine AS builder

WORKDIR /app
COPY . .

WORKDIR /app/cmd/cli
RUN go build -o loadtest

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/cmd/cli/loadtest .

ENTRYPOINT ["./loadtest"]