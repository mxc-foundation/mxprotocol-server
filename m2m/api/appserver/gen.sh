#!/usr/bin/env bash

GRPC_GW_PATH=`go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`
GRPC_GW_PATH="${GRPC_GW_PATH}/../third_party/googleapis"

LS_PATH=`go list -f '{{ .Dir }}' ./...`
LS_PATH="$(LS_PATH)/../m2m"

# generate the gRPC code
protoc -I. -I${GRPC_GW_PATH} -I$(LS_PATH) --go_out=plugins=grpc:. \
    inner_gateway.proto \
    inner_device.proto