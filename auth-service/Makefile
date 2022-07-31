run:
	go run server.go

gen-proto:
	protoc protos/*.proto --plugin=protoc-gen-go=$(go env GOPATH)/bin/protoc-gen-go.exe --go_out=.

gen-grpc:
	protoc protos/auth_service.proto --plugin=protoc-gen-go-grpc=$(go env GOPATH)/bin/protoc-gen-go-grpc.exe --go-grpc_out=.

format:
	gofmt -w ./...

grpcui:
	grpcui -plaintext localhost:4000

build:
	go build