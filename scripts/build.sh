#!/bin/sh

docker build -f deploy/Dockerfile --build-arg SERVICE=${SERVICE} --build-arg VERSION=${VERSION} -t shortener/${SERVICE} -t shortener/${SERVICE}:${VERSION} .
