.PHONY: build clean api dev-requirements pkg/migrations scripts
PKGS := $(shell go list ./... | grep -v /vendor |grep -v m2m/api | grep -v m2m/migrations \
            | grep -v m2m/static | grep -v m2m/ui| grep -v grpc_api-paymemt_service )
VERSION := $(shell git describe --tags |sed -e "s/^v//")

build: pkg/migrations
	mkdir -p build
	go build $(GO_EXTRA_BUILD_ARGS) -ldflags "-s -w -X main.version=$(VERSION)" -o build/m2m command/main.go

clean:
	@echo "Cleaning up workspace"
	@rm -rf build  pkg/migrations/migrations_gen.go

test: pkg/migrations
	@echo "Running tests"
	@rm -f coverage.out
	@for pkg in $(PKGS) ; do \
		golint $$pkg ; \
	done
	@go vet $(PKGS)
	@go test -p 1 -v $(PKGS) -cover -coverprofile coverage.out

api:
	@echo "Generating API code from .proto files"
	@go mod vendor
	@go generate api/appserver/api.go
	@go generate api/networkserver/api.go
	@rm -rf vendor/

pkg/migrations:
	@echo "Generating static files"
	@go generate pkg/migrations/migration.go
