FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o cart-api ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates postgresql-client

WORKDIR /root/

COPY --from=builder /app/cart-api .
COPY --from=builder /app/.env .
COPY --from=builder /app/internal/migrations /root/internal/migrations

EXPOSE ${API_PORT}

CMD ["./cart-api"]
