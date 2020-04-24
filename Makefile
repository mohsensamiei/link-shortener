GOPATH=${HOME}/go

.PHONY: build, test

proto:
	@echo "generating protobuf codes"
	env GOPATH=${GOPATH} scripts/proto.sh
fmt:
	@echo "formating all codes"
	scripts/fmt.sh
test:
	@echo "executing all tests"
	scripts/test.sh
run:
	@echo "running service ${SERVICE}"
	scripts/fmt.sh
	scripts/run.sh
build:
	@echo "building service ${SERVICE}:${VERSION}"
	scripts/build.sh
up:
	@echo "uping services in ${MODE} mode"
	scripts/up.sh
down:
	@echo "downing services in ${MODE} mode"
	scripts/down.sh

