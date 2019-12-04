
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
    protoc --proto_path=. --micro_out=. --go_out=. proto/example/example.proto

	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/example/example.proto

.PHONY: build
build: proto

	go build -o microex-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t microex-srv:latest
