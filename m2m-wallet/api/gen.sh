#!/usr/bin/env bash
GRPC_GW_PATH=`go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`
GRPC_GW_PATH="${GRPC_GW_PATH}/../third_party/googleapis"

# generate the gRPC code
protoc -I. -I${GRPC_GW_PATH} --go_out=plugins=grpc:. \
    internal.proto \
    money.proto \
    super_node.proto \
    topup.proto \
    wallet.proto \
    withdraw.proto

# generate the JSON interface code
protoc -I. -I${GRPC_GW_PATH} --grpc-gateway_out=logtostderr=true:. \
    internal.proto \
    money.proto \
    super_node.proto \
    topup.proto \
    wallet.proto \
    withdraw.proto

# generate the swagger definitions
protoc -I. -I${GRPC_GW_PATH} --swagger_out=json_names_for_fields=true:./swagger \
    internal.proto \
    money.proto \
    super_node.proto \
    topup.proto \
    wallet.proto \
    withdraw.proto