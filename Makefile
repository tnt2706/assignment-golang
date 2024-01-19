.PHONY:	start gen-cal gen-grpc

start:
	air /cmd/main.go

gen-cal:
    protoc \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/calculator/proto/calculator.proto

gen-grpc:
    protoc --go_out=. --go_opt=Mprotos/calculator.proto=pb \
    --go-grpc_out=. --go-grpc_opt=Mprotos/calculator.proto=pb \
    pkg/calculator/proto/calculator.proto

gen-graphql:
    go env -w GOFLAGS=-mod=mod \
    go run github.com/99designs/gqlgen generate
