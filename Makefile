SHELL := /bin/bash

.PHONY: update-protos
update-protos:	
	@./hack/update-proto.sh

.PHONY: protoc
protoc:
	$(MAKE) -C _proto/ all