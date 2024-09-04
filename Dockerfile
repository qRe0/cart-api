FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o cart-api ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/cart-api .
COPY --from=builder /app/.env .
COPY --from=builder /app/internal/migrations ./internal/migrations

EXPOSE ${API_PORT}

CMD ["./cart-api"]
