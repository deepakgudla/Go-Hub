# Build Stage
FROM golang:1.23.1 AS builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

# Run Stage
FROM alpine:latest
LABEL maintainer="deepakgudla"
LABEL org.opencontainers.image.source="https://github.com/deepakgudla/Go-Hub"
LABEL org.opencontainers.image.description="Basic Go server running on port 1357"
LABEL org.opencontainers.image.licenses="MIT"

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 1357

CMD ["./main"]
