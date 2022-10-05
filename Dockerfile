FROM golang:1.19-alpine as builder

WORKDIR /tmp/grpc-chat

COPY . .

ENV CGO_ENABLED=0

RUN go build -ldflags="-d -s -w" -tags timetzdata -trimpath -o app .

# --- Execution Stage

FROM scratch

COPY --from=builder /tmp/grpc-chat/app /app
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 6262/tcp

ENTRYPOINT ["/app"]
