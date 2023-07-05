.PHONY:	start gen-cal gen-grpc

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
