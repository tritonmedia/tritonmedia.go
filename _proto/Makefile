space := $(empty) $(empty)
comma := ,
empty :=

import_path = github.com/tritonmedia/tritonmedia.go/pkg/api

dst 	= ../../pkg/proto
target  = go

api_ias    = $(subst $(space),$(comma),$(addsuffix =$(import_path)/$(api_pkg),$(addprefix M,$(api_pbs))))
api_pkg    = api

.PHONY: all
all: api

.PHONY: api
api:
	cd api; PATH=../bin:$(PATH) protoc --$(target)_out=$(api_ias):$(dst) *.proto; cd ..


.PHONY: clean
clean:
	@rm -rf $(dst)/hapi 2>/dev/null