* Regenerate gRPC code
```
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto

```

* Go module
```
# 创建go.mod 文件
$ go mod init github.com/chyidl/hello

$ go mod tidy
```
