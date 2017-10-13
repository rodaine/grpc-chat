FROM golang:1.9-alpine as builder
RUN apk --no-cache add git

RUN go get -d \
    github.com/pkg/errors \
    golang.org/x/net/context \
    google.golang.org/grpc \
    github.com/golang/protobuf/ptypes

WORKDIR /go/src/github.com/rodaine/grpc-chat
COPY . .

RUN go build -o app .

# --- Execution Stage

FROM alpine:latest
EXPOSE 6262/tcp

WORKDIR /root/
COPY --from=builder /go/src/github.com/rodaine/grpc-chat/app .

ENTRYPOINT ["./app"]
