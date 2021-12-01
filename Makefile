VERSION = $(shell (git describe --tags --abbrev=0 2>/dev/null || echo '0.0.0') | tr -d v)
COMMIT = $(shell git rev-parse --short HEAD)

OUT = build/
# export GOPROXY = direct
# export CGO_ENABLED = 0

.DEFAULT_GOAL := build

.PHONY: info
info:
	@echo "VERSION: $(VERSION)"
	@echo "COMMIT:  $(COMMIT)"

.PHONY: codegen
codegen: grpc

.PHONY: grpc
grpc: grpc-go grpc-python

.PHONY: grpc-go
grpc-go:
	@docker run --rm -v $$(pwd):/oomstore -w /oomstore rvolosatovs/protoc \
		--experimental_allow_proto3_optional \
		-I=proto \
		--go_out=oomagent \
		--go-grpc_out=oomagent \
		proto/oomagent.proto

.PHONY: grpc-python
grpc-python:
	@docker run --rm -v $$(pwd):/oomstore -w /oomstore rvolosatovs/protoc \
		--experimental_allow_proto3_optional \
		-I=proto \
		--python_out=sdk/python/src/oomstore/codegen \
		--grpc-python_out=sdk/python/src/oomstore/codegen \
		proto/oomagent.proto
	@perl -pi -e s,"import status_pb2","from . import status_pb2",g sdk/python/src/oomstore/codegen/oomagent_pb2.py
	@perl -pi -e s,"import oomagent_pb2","from . import oomagent_pb2",g sdk/python/src/oomstore/codegen/oomagent_pb2_grpc.py

.PHONY: build
build: oomcli oomagent

.PHONY: oomcli
oomcli:
	$(MAKE) -C oomcli build

.PHONY: oomagent
oomagent: codegen
	$(MAKE) -C oomagent build

.PHONY: test
test: codegen
	@go test -race -coverprofile=coverage.out -covermode=atomic ./...

.PHONY: integration-test
integration-test: codegen
	$(MAKE) -C oomcli   integration-test
	$(MAKE) -C oomagent integration-test

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: start_playground
start_playground:
	@if [ "$$(docker ps -q -f name=oomstore_playtround)" ]; then \
		echo "already running"; \
	else \
		docker run --rm -d --name oomstore_playtround \
			-e POSTGRES_PASSWORD=postgres \
			-e POSTGRES_USER=postgres \
			-p 5432:5432 \
			postgres:14.0-alpine; \
		docker exec oomstore_playtround sh -c 'while ! pg_isready; do sleep 1; done'; \
	fi

.PHONY: stop_playground
stop_playground:
	@docker stop oomstore_playtround 2>/dev/null || true

.PHONY: restart_playground
restart_playground: stop_playground start_playground

.PHONY: clean
clean:
	$(MAKE) -C oomcli clean
