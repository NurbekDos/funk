FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s"  -o /funk

FROM alpine:latest

WORKDIR /app

COPY --from=builder /funk /funk
COPY .env .env
COPY internal/db/migrations /app/internal/db/migrations

EXPOSE 8081

CMD ["/funk"]
# CMD ls -l