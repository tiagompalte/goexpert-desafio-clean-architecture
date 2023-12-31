run:
	cd cmd/ordersystem/ && go run main.go wire_gen.go

test: 
	go test -v --cover ./...

wire:
	wire ./cmd/ordersystem

protoc:
	protoc --go_out=. --go-grpc_out=. ./internal/infra/grpc/protofiles/order.proto 

gqlgen:
	go run github.com/99designs/gqlgen generate

evans:
	evans -r repl

.PHONY: run test wire protoc gqlgen evans