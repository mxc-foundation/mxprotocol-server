.PHONY: subdirs all clean dev-requirements
SUBDIRS=m2m

subdirs:
	@for subdir in $(SUBDIRS); \
	do \
	echo "build in $$subdir"; \
	( cd $$subdir && make build) || exit 1; \
	done

all: subdirs

clean:
	@for subdir in $(SUBDIRS); \
	do \
	echo "cleaning in $$subdir"; \
	( cd $$subdir && make clean) || exit 1; \
	done

# shortcuts for development
dev-requirements:
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go install github.com/golang/protobuf/protoc-gen-go
	go install github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs
	go install github.com/jteeuwen/go-bindata/go-bindata


