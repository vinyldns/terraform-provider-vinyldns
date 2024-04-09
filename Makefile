SHELL=bash
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
GO_TARGET=$(ROOT_DIR)/...
PKG_NAME=vinyldns
NAME=terraform-provider-vinyldns
WEBSITE_REPO=github.com/hashicorp/terraform-website
VINYLDNS_VERSION=0.20.0

# Check that the required version of make is being used
REQ_MAKE_VER:=3.82
ifneq ($(REQ_MAKE_VER),$(firstword $(sort $(MAKE_VERSION) $(REQ_MAKE_VER))))
   $(error The version of MAKE $(REQ_MAKE_VER) or higher is required; you are running $(MAKE_VERSION))
endif

.ONESHELL:

all: test build


.PHONY: start-api
start-api: stop-api
	@set -euo pipefail
	echo "Starting VinylDNS API.."
	docker run -d --name vinyldns-terraform-api -p "9000:9000" -p "19001:19001" -e RUN_SERVICES="all tail-logs" vinyldns/build:base-test-integration-v$(VINYLDNS_VERSION)
	echo "Waiting for VinylDNS API to start.."
	{ timeout "20s" grep -q 'STARTED SUCCESSFULLY' <(timeout 20s docker logs -f vinyldns-terraform-api); } || { echo "VinylDNS API failed to start"; exit 1; }

.PHONY: stop-api
stop-api:
	@set -euo pipefail
	if docker ps | grep -q "vinyldns-terraform-api"; then
		docker kill vinyldns-terraform-api &> /dev/null || true
		docker rm vinyldns-terraform-api &> /dev/null || true
	fi

.PHONY: test
test: start-api execute-tests stop-api

.PHONY: execute-tests
execute-tests:
	set -euo pipefail
	go vet "$(GO_TARGET)"
	go test "$(GO_TARGET)" -cover
	VINYLDNS_ACCESS_KEY=okAccessKey \
	VINYLDNS_SECRET_KEY=okSecretKey \
	VINYLDNS_HOST=http://localhost:9000 \
	TF_LOG=DEBUG \
	TF_ACC=1 \
	go test "$(GO_TARGET)" -v

.PHONY: install
install:
	@set -euo pipefail
	go install
	mkdir -p "$${HOME}/.terraform.d/plugins/local/vinyldns-provider/vinyldns/0.0.1/$$(go env GOOS)_$$(go env GOARCH)/"
	cp "$$(go env GOPATH)/bin/terraform-provider-vinyldns" "$${HOME}/.terraform.d/plugins/local/vinyldns-provider/vinyldns/0.0.1/$$(go env GOOS)_$$(go env GOARCH)/"

.PHONY: clean
clean:
	@set -euo pipefail
	if [ -d "$(ROOT_DIR)/build/" ]; then
		echo -n "Cleaning existing build.." && \
		rm -rf "$(ROOT_DIR)/build" && echo "done."
	fi

.PHONY: build
build: clean
	@set -euo pipefail
	echo -n "Building provider.." && \
	mkdir -p "$(ROOT_DIR)/build/$$(go env GOOS)_$$(go env GOARCH)/" && \
	go build  -o "$(ROOT_DIR)/build/$$(go env GOOS)_$$(go env GOARCH)/" "$(GO_TARGET)" && \
	echo "done." && \
	echo "Compiled binary: $(ROOT_DIR)/build/$$(go env GOOS)_$$(go env GOARCH)/$(NAME)"

.PHONY: version
version:
	echo $(VERSION)

.PHONY: website
website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(ROOT_DIR) PROVIDER_NAME=$(PKG_NAME)

.PHONY: website-test
website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(ROOT_DIR) PROVIDER_NAME=$(PKG_NAME)


