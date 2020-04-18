#!/bin/sh

docker build -f deploy/Dockerfile --build-arg SERVICE=${SERVICE} --build-arg VERSION=${VERSION} -t link/${SERVICE} -t link/${SERVICE}:${VERSION} .
