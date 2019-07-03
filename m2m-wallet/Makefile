.PHONY: build clean api pkg/static statics/swagger/api.swagger.json dev-requirements
VERSION := "test-run"

build: ui/build pkg/static
	mkdir -p build
	go build $(GO_EXTRA_BUILD_ARGS) -ldflags "-s -w -X main.version=$(VERSION)" -o build/m2m-wallet command/main.go

clean:
	@echo "Cleaning up workspace"
	@rm -rf build pkg/static/static_gen.go ui/build static/static
	@rm -f static/index.html static/icon.png static/manifest.json static/asset-manifest.json static/service-worker.js
	@rm -rf static/logo

ui/build:
	@echo "Building ui"
	@cd ui && npm run build
	@mv ui/build/* static
api:
	@echo "Generating API code from .proto files"
	@go mod vendor
	@go generate api/api.go
	@rm -rf vendor/

pkg/static: static/swagger/api.swagger.json
	@echo "Generating static files"
	@go generate pkg/static/static.go

static/swagger/api.swagger.json:
	@echo "Generating combined Swagger JSON"
	@GOOS="" GOARCH="" go run api/swagger/main.go api/swagger > static/swagger/api.swagger.json
	@cp api/swagger/*.json static/swagger

ui-requirements:
	@echo "Installing UI requirements"
	@cd ui && npm install
