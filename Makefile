
.PHONY: init
# init env
init:
	go mod tidy
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.3

.PHONY: validatex
# generate validatex
validatex:
	protoc --proto_path=./validatex \
		   --proto_path=./third_party \
		   --go_out=paths=source_relative:./validatex \
		   ./validatex/*.proto

.PHONY: example
# generate example
example:
	go install . && \
	protoc --proto_path=. \
		   --go_out=paths=source_relative:. \
		   --validatex_out=paths=source_relative:. \
		   --validatex_opt=i18n_dir=./example/i18n,i18n_out_relative_dir=./example/i18n,json_name=true \
		   ./example/*.proto

.PHONY: all
# generate all
all:
	make validatex;
	make example;
	go mod tidy;
