#!/bin/sh

cd ./api
rm -rf *.go
cd ./src
for FILE in $(ls *.proto); do
  protoc \
    --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go \
    --proto_path=${GOPATH}/src:. \
    --go_out=../ \
    ${FILE}
done
