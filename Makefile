# Setup Project
PROJECT_NAME := caster
PROJECT_REPO := github.com/bkonkle/go-example-caster-api

PLATFORMS ?= linux_amd64 linux_arm64
-include build/makelib/common.mk

# Setup Output
-include build/makelib/output.mk

# Setup Go
NPROCS ?= 1
GO_TEST_PARALLEL := $(shell echo $$(( $(NPROCS) / 2 )))
GO_STATIC_PACKAGES = $(GO_PROJECT)/cmd/${PROJECT_NAME}
GO_LDFLAGS += -X $(GO_PROJECT)/internal/version.Version=$(VERSION)
GO_SUBDIRS += cmd internal
GO111MODULE = on
-include build/makelib/golang.mk

.PHONY: docker-up
docker-up:
	docker-compose -f docker-compose.yml up -d
