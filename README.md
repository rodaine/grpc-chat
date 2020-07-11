# grpc-chat [![](https://img.shields.io/docker/automated/rodaine/teamspeak3-alpine.svg)](https://hub.docker.com/r/rodaine/grpc-chat/)

A simple chat server/client implemented with [gRPC](https://grpc.io) in Go. Built with :heart: for the [Orange County Gopher's Meetup](https://www.meetup.com/Orange-County-Gophers/).

> This project is for demonstrating some of the features of gRPC in Go and should not be used in production.

## Installation

Installation requires the Go toolchain.

```bash
go install github.com/rodaine/grpc-chat
```

## Usage

```bash
$ grpc-chat --help
Usage of grpc-chat:
  -h string
      the chat server's host (default "0.0.0.0:6262")
  -n string
      the username for the client
  -p string
      the chat server's password
  -s run as the server
  -v enable debug logging
```

### Server

```bash
grpc-chat -s -p "super-secret"
```

### Client

```bash
grpc-chat -h "chat.example.com:6262" -p "super-secret" -n "Rodaine"
```

## Docker

A Dockerfile is included with this project and the image is [hosted on the Docker Hub](https://hub.docker.com/r/rodaine/grpc-chat).

### Run as Server

```bash
docker run --rm \
  -p 6262:6262 \
  rodaine/grpc-chat \
  -s -p "super-secret"
```

### Run as Client

```bash
docker run --rm -i \
  rodaine/grpc-chat \
  -h "chat.example.com" \
  -p "super-secret" \
  -n "Rodaine"
```
