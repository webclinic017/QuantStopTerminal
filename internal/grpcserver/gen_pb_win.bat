@echo off
echo QSTrader: Generating gRPC, proxy files.
REM You may need to include the go mod package for the annotations file:
REM %GOPATH%\pkg\mod\github.com\grpc-ecosystem\grpc-gateway@v1.16.0\third_party\googleapis

protoc -I=. -I=%GOPATH%\src -I=%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --go_out=paths=source_relative:. rpc.proto
protoc -I=. -I=%GOPATH%\src -I=%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --go-grpc_out=paths=source_relative:. rpc.proto
protoc -I=. -I=%GOPATH%\src -I=%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --grpc-gateway_out=paths=source_relative,logtostderr=true:. rpc.proto
protoc -I=. -I=%GOPATH%\src -I=%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --openapiv2_out=logtostderr=true:. rpc.proto
