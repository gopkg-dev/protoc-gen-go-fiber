install:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install .

gen_example:
	go install
	cd ./example && buf generate
	protoc-gen-go-fiber add-tag "./example/gen"