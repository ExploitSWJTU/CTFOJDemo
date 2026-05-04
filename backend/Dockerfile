FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/server ./cmd/server

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /out/server /usr/local/bin/server

EXPOSE 8080

CMD ["server"]
