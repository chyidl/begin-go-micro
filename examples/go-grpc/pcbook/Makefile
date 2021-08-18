gen:
	protoc --proto_path=proto \
		--go_out=plugins=grpc:pb \
        proto/*.proto

py-gen:
	python3 -m grpc_tools.protoc --proto_path=proto \
		--python_out=py-pb \
		proto/*.proto


clean:
	rm pb/*.go

run:
	go run main.go

.PHONY: gen py-gen clean run
