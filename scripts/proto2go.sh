#!/bin/sh

# Install the following tools for grpc-gateway
# go install \
#     github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
#     github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
#     google.golang.org/protobuf/cmd/protoc-gen-go \
#     google.golang.org/grpc/cmd/protoc-gen-go-grpc

SRC=api/protobuf-spec
GOOGLEAPIS=third_party/googleapis
DEST_RES=./gen/go/resource
DEST_SVC=./gen/go/service

protoc -I $SRC -I $GOOGLEAPIS \
  --go_opt=paths=source_relative --go_out=$DEST_RES \
  $SRC/resource.proto

protoc -I $SRC -I $GOOGLEAPIS \
  --go_opt=paths=source_relative --go_out=$DEST_SVC \
  --go-grpc_opt=paths=source_relative --go-grpc_out=$DEST_SVC \
  --grpc-gateway_opt logtostderr=true --grpc-gateway_opt=paths=source_relative --grpc-gateway_out=$DEST_SVC \
  --openapiv2_opt=logtostderr=true --openapiv2_out=./gen/openapiv2 \
  $SRC/echo_service.proto