# Go - Docker

### Folder structure
```
└── 📁Docker
    └── .env
    └── .gitignore
    └── Dockerfile
    └── go.mod
    └── go.sum
    └── main_test.go
    └── main.go
    └── Readme.md
```

### Instructions

- **step-1** clone the repo using the following command
    - git clone `https://github.com/deepakgudla/Go-Hub.git`
- **step-2** change directory to Docker folder using the following command
    - `cd Docker`
- **step-3**Install the dependencies using the following command
    - `go mod tidy`
- **step-4** Build the docker image using the following command
    - `docker build -t deepakgudla/go-basic-server:latest .`
- **step-5** Run the docker image using the following command
    - `docker run -p 1357:1357 deepakgudla/go-basic-server:v1`
    - `docker run --env-file .env -p 1357:1357 deepakgudla/go-basic-server:v1`
- **step-6** To publish the image into Docker hub, please login using `docker login` 
- **step-7** Push the image into Docker Hub using the following command
    - `docker push deepakgudla/go-basic-server:latest` 

### Working of Application 
- If you want to test this application directly, pull the image using the following command
    - `docker pull deepakgudla/go-basic-server`
    - open `http://localhost:1357' on your browser
    - You will be able to see `Jai Mahishmathi ✊`

- **NOTE**:
    - If you want to push the Docker image into your own Docker repository, login with your own docker credentials and replace `deepakgudla` with your docker username..
  
### The Go web server

- `main.go`
    - This is a simple web server that listens on a port and responds with a text upon accessing the URL
- `main_test.go`
    - simple unit test to make sure that the web server is working fine 

### Dockerfile

- What is a Dockerfile ?
    - A `Dockerfile` defines how to build and run a lightweight, production ready container

#### Dockerfile explanation

- **FROM golang:1.23.1 AS builder**
    - `FROM` - it is a keyword
    - Uses the official Go (Golang) image version 1.23.1, which includes everything needed to compile Go code
    - `AS builder` - gives this stage a name `builder`
- `WORKDIR /app`
    - Sets the working directory inside the container to `/app`
    - All the commands after this will run inside the `/app`
- `COPY go.mod ./`
    - Copies the Go module files into the container. 
- `COPY go.sum ./`
    - This file describes the app dependencies
- `RUN go mod download`
    - Downloads the dependencies described in the go.mod and go.sum files.
- `COPY . ./`
    - Copies the entire project (all source code) into the container’s /app directory.
- `RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main`
    - Compiles the Go app into a binary named main
    - **CGO_ENABLED=0** - disables C bindings for a fully static binary
    - **GOOS=linux** - Go OS, targets linux system
    - **GOARCH** - Go Architecture, targets 64 bit architecture
- `FROM alpine:latest`
    - Starts a second, clean image using Alpine Linux which is lightweight (has less features compared to ubuntu, efficient for single web server)
- `Metadata`
    -  LABEL maintainer="deepakgudla"
    - LABEL org.opencontainers.image.source="https://github.com/deepakgudla/Go-Hub"
    - LABEL org.opencontainers.image.description="Basic Go server running on port 1357"
    - LABEL org.opencontainers.image.licenses="MIT"
    - These labels are metadata for the Docker image. The metadata can be found in Docker Hub or container registries
- `RUN apk --no-cache add ca-certificates`
    -  Installs root SSL certificates needed for HTTPS connections.
- `WORKDIR /root/`
    - Sets the working directory to `/root/` inside the container
- `COPY --from=builder /app/main .`
    - Copies the compiled main binary from the builder stage  into the current stage
- `EXPOSE 1357` 
    - The docker image is set to run on port 1357
    - It just exposes the port but does not actually publishes the port 
- `CMD ["./main"]`
    - This is the command that tells Docker to run the main binary when the container starts
    - This launches the Go web server 

#### Layers in the Docker image

- What is a Docker layer ?
    - Each step in a Dockerfile is considered as a layer. 
    - Each layer is built on top of previous one
    - When the Dockerfile gets updated, Docker only rebuilds the layers that has been modified and the ones after it 

- Layers in the current Dockerfile
  ![image](https://github.com/user-attachments/assets/f61c3dd1-6b2b-424e-893c-ed240280ccc5)
