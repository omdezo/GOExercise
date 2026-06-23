# Build stage
FROM golang:1.18 AS builder

WORKDIR /app

COPY go.mod ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o exercise-go .

# Runtime stage
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/exercise-go .

CMD ["./exercise-go"]
