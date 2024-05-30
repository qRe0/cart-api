FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o start ./cmd/main.go
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:latest

RUN apk --no-cache add ca-certificates postgresql-client

WORKDIR /root/

COPY --from=builder /app/start .
COPY --from=builder /app/.env .
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY internal/migrations /migrations

EXPOSE ${API_PORT}

CMD ["/bin/sh", "-c", "goose -dir /migrations postgres \"postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DB_PORT}/${DATABASE_NAME}?sslmode=disable\" up && ./start"]
