FROM golang:1.23-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# Final stage
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/garm.pem .

EXPOSE 8080

CMD ["./main"]
