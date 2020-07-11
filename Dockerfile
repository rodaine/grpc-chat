FROM golang:1.14-alpine as builder
RUN apk --no-cache add git

WORKDIR /go/src/github.com/rodaine/grpc-chat
COPY go.* ./

RUN go mod download

COPY . .
RUN go build -o app .

# --- Execution Stage

FROM alpine:latest
EXPOSE 6262/tcp

WORKDIR /root/
COPY --from=builder /go/src/github.com/rodaine/grpc-chat/app .

ENTRYPOINT ["./app"]
