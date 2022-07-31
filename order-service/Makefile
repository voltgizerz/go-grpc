run:
	go run server.go

gen-proto:
	protoc protos/*.proto --plugin=protoc-gen-go=$(go env GOPATH)/bin/protoc-gen-go.exe --go_out=.

gen-grpc:
	protoc protos/order_service.proto --plugin=protoc-gen-go-grpc=$(go env GOPATH)/bin/protoc-gen-go-grpc.exe --go-grpc_out=.

format:
	gofmt -w ./...

update-grpc:
	go get -u github.com/voltgizerz/public-grpc@latest

build:
	go build