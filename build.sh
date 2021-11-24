#!/bin/bash

set -e

GROUPS_PROJECT_ID="19810936"
GROUPS_VERSION="develop"

curl -o proto/groups.proto --header "PRIVATE-TOKEN: $GITLAB_TOKEN" "https://gitlab.com/api/v4/projects/$GROUPS_PROJECT_ID/repository/files/proto%2Fgroups.proto/raw?ref=$GROUPS_VERSION"

GEN_PATH="grpc"
GO_LIB_PATH=$(go env GOPATH)/src
GOPATH=$(go env GOPATH)

mkdir -p proto/google/api 
curl -o proto/google/api/annotations.proto "https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto"
curl -o proto/google/api/http.proto "https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto"

mkdir -p proto/google/protobuf
curl -o proto/google/protobuf/descriptor.proto "https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf/descriptor.proto"
curl -o proto/google/protobuf/empty.proto "https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf/empty.proto"
curl -o proto/google/protobuf/struct.proto "https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf/struct.proto"
curl -o proto/google/protobuf/timestamp.proto "https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf/timestamp.proto"
curl -o proto/google/protobuf/any.proto "https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf/any.proto"

echo "gen with protoc"
protoc \
        -I proto \
        -I $GO_LIB_PATH/include \
        --go_out=$GEN_PATH      --go_opt=paths=source_relative \
        --go-grpc_out=$GEN_PATH --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=logtostderr=true,paths=source_relative:$GEN_PATH \
        --openapiv2_out=logtostderr=true:$GEN_PATH \
        proto/robots.proto

protoc-go-inject-tag -input=./$GEN_PATH/robots.pb.go

echo "tidy"
go mod tidy
go get -u ./...
