.PHONY: compile


compile:
	# protoc-gen-go 不支持多包同时编译
	# proto_path或者-I 指定编译源码; proto 文件中使用的import关键字导入的路径一定要基于proto_path
	# generate gRPC code
	protoc --proto_path=. \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
        proto/**/*.proto