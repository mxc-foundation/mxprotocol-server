.PHONY: build clean api pkg/static statics/swagger/api.swagger.json dev-requirements pkg/migrations
PKGS := $(shell go list ./... | grep -v /vendor |grep -v m2m-wallet/api | grep -v m2m-wallet/migrations \
            | grep -v m2m-wallet/static | grep -v m2m-wallet/ui| grep -v grpc_api-paymemt_service )
VERSION := $(shell git describe --always |sed -e "s/^v//")

build: ui/build pkg/static pkg/migrations
	mkdir -p build
	go build $(GO_EXTRA_BUILD_ARGS) -ldflags "-s -w -X main.version=$(VERSION)" -o build/m2m command/main.go

clean:
	@echo "Cleaning up workspace"
	@rm -rf build pkg/static/static_gen.go pkg/migrations/migrations_gen.go ui/build static/static
	@rm -f static/index.html static/icon.png static/manifest.json static/asset-manifest.json static/service-worker.js
	@rm -f static/swagger/*.json
	@rm -rf static/logo

test: pkg/static pkg/migrations
	@echo "Running tests"
	@rm -f coverage.out
	@for pkg in $(PKGS) ; do \
		golint $$pkg ; \
	done
	@go vet $(PKGS)
	@go test -p 1 -v $(PKGS) -cover -coverprofile coverage.out

ui/build:
	@echo "Building ui"
	@cd ui && npm run build
	@mv ui/build/* static
api:
	@echo "Generating API code from .proto files"
	@go mod vendor
	@go generate api/m2m/api.go 
	@go generate api/appserver/api.go
	@rm -rf vendor/

pkg/static pkg/migrations: static/swagger/api.swagger.json
	@echo "Generating static files"
	@go generate pkg/static/static.go
	@go generate pkg/migrations/migration.go

static/swagger/api.swagger.json:
	@echo "Generating combined Swagger JSON"
	@GOOS="" GOARCH="" go run api/swagger/main.go api/swagger > static/swagger/api.swagger.json
	@cp api/swagger/*.json static/swagger

ui-requirements:
	@echo "Installing UI requirements"
	@cd ui && npm install
