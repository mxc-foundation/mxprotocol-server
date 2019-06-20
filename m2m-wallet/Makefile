.PHONY: build clean api pkg/static statics/swagger/api.swagger.json dev-requirements
VERSION := "test-run"

build: pkg/static
	mkdir -p build
	go build -ldflags "-s -w -X main.version=$(VERSION)" -o build/mxp-server command/main.go

clean:
	@echo "Cleaning up workspace"
	@rm -rf build pkg/static/static_gen.go

api:
	@echo "Generating API code from .proto files"
	@go generate api/api.go

pkg/static: statics/swagger/api.swagger.json
	@echo "Generating static files"
	@go generate pkg/static/static.go

statics/swagger/api.swagger.json:
	@echo "Generating combined Swagger JSON"
	@GOOS="" GOARCH="" go run api/swagger/main.go api/swagger > statics/swagger/api.swagger.json
	@cp api/swagger/*.json statics/swagger

# shortcuts for development

dev-requirements:
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go install github.com/golang/protobuf/protoc-gen-go
	go install github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs
	go install github.com/jteeuwen/go-bindata/go-bindata
	go install golang.org/x/tools/cmd/stringer
