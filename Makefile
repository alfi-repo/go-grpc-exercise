generate:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

run-server:
	go run server/main.go

run-client:
	go run client/main.go