## QuantstopTerminal Development Guide - gRPC Server

### Prerequisites:
QuantstopTerminal utilizes gRPC for all external communication with the daemon processes.
A standard gRPC server is provided, along with command line interface client.
In addition, a gRPC REST proxy server is provided, if the web based frontend is enabled. 
To develop on these systems you must install the following pre-requisites:

* [protoc](https://developers.google.com/protocol-buffers/docs/gotutorial)
* [gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/)
* [buf](https://docs.buf.build/installation)




To generate the go code from .proto file, all command are wrapped up in a Task.
Simply go to the project root, and run the following command in a terminal:
```bash
task gen-proto
```


<!--```bash
protoc --go_opt=paths=source_relative --go_out=plugins=grpc:./ rpc.proto
go get \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```-->