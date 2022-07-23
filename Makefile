.PHONY : proto

WORKDIR := ${PWD}


proto:
	protoc --go_out=. --go_opt=Mprotos/calculator.proto=pb \
  --go-grpc_out=. --go-grpc_opt=Mprotos/calculator.proto=pb \
  src/grpc/protos/calculator.proto