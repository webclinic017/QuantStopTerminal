
To generate the go code from .proto file

```bash
protoc --go_opt=paths=source_relative --go_out=plugins=grpc:./ rpc.proto
```


```bash
go get \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```