FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates postgresql-client

WORKDIR /root/

COPY --from=builder /app/api .
COPY --from=builder /app/.env .
COPY --from=builder /app/internal/migrations /root/internal/migrations

EXPOSE ${API_PORT}

CMD ["./api"]
