FROM golang:1.21.1-bookworm AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY main.go ./

RUN go build -o /my-app

FROM seladb/ubuntu2204

COPY --from=build /my-app /my-app

ENTRYPOINT ["/my-app"]
