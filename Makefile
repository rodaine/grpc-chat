.PHONY: install
install: protos/chat.pb.go
	go install .
	which grpc-chat

protos/chat.pb.go:
	protoc --go_opt=paths=source_relative --go_out="." --go-grpc_out=module=github.com/rodaine/grpc-chat/protos:protos protos/chat.proto

.PHONY: docker
docker:
	docker build --rm -t rodaine/grpc-chat .

.PHONY: plugins
plugins:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
