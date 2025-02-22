FROM golang:latest AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o task-service main.go

# Minimal runtime image
FROM alpine:latest

WORKDIR /app
RUN apk add --no-cache ca-certificates libc6-compat sqlite-libs
COPY --from=builder /app/task-service /app/task-service

RUN chmod +x /app/task-service

EXPOSE 3030
CMD ["/app/task-service"]
