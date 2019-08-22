.PHONY: subdirs all clean ui-subdirs ui-requirements dev-requirements
SUBDIRS=m2m-wallet

subdirs:
	@for subdir in $(SUBDIRS); \
	do \
	echo "build in $$subdir"; \
	( cd $$subdir && make build) || exit 1; \
	done

all: subdirs

ui-subdirs:
	@for subdir in $(SUBDIRS); \
	do \
	echo "build ui in $$subdir"; \
	( cd $$subdir && make ui-requirements) || exit 1; \
	done

clean:
	@for subdir in $(SUBDIRS); \
	do \
	echo "cleaning in $$subdir"; \
	( cd $$subdir && make clean) || exit 1; \
	done

# shortcuts for development
ui-requirements: ui-subdirs

dev-requirements:
	go install golang.org/x/lint/golint
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go install github.com/golang/protobuf/protoc-gen-go
	go install github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs
	go install github.com/jteeuwen/go-bindata/go-bindata
	go install golang.org/x/tools/cmd/stringer
	go install github.com/goreleaser/nfpm

