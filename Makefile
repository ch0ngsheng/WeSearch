# Image URL to use all building/pushing image targets
IMGManager ?= yuchsh/wesearch-manager:latest
IMGUserDoc ?= yuchsh/wesearch-userdoc:latest
IMGRetrieve ?= yuchsh/wesearch-retrieve:latest

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

##@ Build

.PHONY: build
build: fmt ## Build manager binary.
	go build -o manager services/wxmanager/api/manager.go

.PHONY: run-manager
run: fmt vet ## Run a controller from your host.
	go run services/wxmanager/api/manager.go

.PHONY: docker-build-manager
docker-build-manager: ## Build docker image with the manager.
	docker build -t ${IMGManager} -f services/wxmanager/api/Dockerfile .

.PHONY: docker-build-userdoc
docker-build-userdoc: ## Build docker image with the manager.
	docker build -t ${IMGUserDoc} -f services/userdocument/rpc/Dockerfile .

.PHONY: docker-build-retrieve
docker-build-retrieve: ## Build docker image with the manager.
	docker build -t ${IMGRetrieve} -f services/retrieve/rpc/Dockerfile .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	docker push ${IMGManager}
	docker push ${IMGUserDoc}
	docker push ${IMGRetrieve}