#!/usr/bin/env bash

go get -d -u github.com/grpc-ecosystem/grpc-gateway/...

sed \
  -e 's:runtime:jsonpb:' \
  -e 's:golang/protobuf:gogo/protobuf:g' \
  -e 's:JSONPb:GoGoJSONPb:g' \
  -e 's:DecoderWrapper:GoGoDecoderWrapper:g' \
  -e 's:DisallowUnknownFields:GoGoDisallowUnknownFields:g' \
  -e '/import (/a \
  . "github.com/grpc-ecosystem/grpc-gateway/runtime" // nolint: golint' \
  $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/runtime/marshal_jsonpb.go > marshaling.go

GO111MODULE=on go run golang.org/x/tools/cmd/goimports -w marshaling.go

echo "// Code generated from github.com/grpc-ecosystem/grpc-gateway/runtime by copy.sh

$(cat marshaling.go)" > marshaling.go
