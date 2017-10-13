.PHONY: install
install: protos/chat.pb.go
	go install .
	which grpc-chat

protos/chat.pb.go:
	protoc --go_out="plugins=grpc:." protos/chat.proto

.PHONY: docker
docker:
	docker build --rm -t rodaine/grpc-chat .
