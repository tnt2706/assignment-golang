.PHONY:	start gen-cal gen-grpc run-grpc-server gen-graphql

start:
	air

gen-cal:
    protoc \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/calculator/proto/calculator.proto


# setup url go_package
gen-grpc:
    protoc --go_out=. --go_opt=Mprotos/calculator.proto=pb \
    --go-grpc_out=. --go-grpc_opt=Mprotos/calculator.proto=pb \
    pkg/calculator/proto/calculator.proto


gen-graphql:
    go env -w GOFLAGS=-mod=mod \
    go run github.com/99designs/gqlgen generate.


# run-grpc-server:
#     go run ./pkg/calculator/calculatorserver.go
