SHELL := /bin/bash

.PHONY: update-protos
update-protos:	
	@./hack/update-proto.sh

.PHONY: protoc
protoc:
	$(MAKE) -C _proto/ all

.PHONY: test
test:
	go test -race

.PHONY: render-circle
render-circle:
	@if [[ ! -e /tmp/jsonnet-libs ]]; then git clone git@github.com:tritonmedia/jsonnet-libs /tmp/jsonnet-libs; else cd /tmp/jsonnet-libs; git pull; fi
	JSONNET_PATH=/tmp/jsonnet-libs jsonnet .circleci/circle.jsonnet | yq . -y > .circleci/config.yml