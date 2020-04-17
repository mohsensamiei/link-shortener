#!/bin/sh

cd ./api
rm -rf *.pb.go
cd ./src
for FILE in $(ls *.proto); do
  protoc -I . ${FILE} --go_out=plugins=grpc:../
done
