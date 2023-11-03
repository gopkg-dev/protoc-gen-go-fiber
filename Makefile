install:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/favadi/protoc-go-inject-tag@latest

gen_example:
	go install
	cd ./example && buf generate
	protoc-go-inject-tag -input="./example/gen/*/*.pb.go"