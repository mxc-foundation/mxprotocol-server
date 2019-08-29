#!/usr/bin/env bash
GRPC_GW_PATH=`go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`
GRPC_GW_PATH="${GRPC_GW_PATH}/../third_party/googleapis"

# generate the gRPC code
protoc -I. -I${GRPC_GW_PATH} -I../../api --go_out=plugins=grpc:. \
    inner_gateway.proto