SHELL=bash
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
GO_TARGET=$(ROOT_DIR)/...
PKG_NAME=vinyldns
NAME=terraform-provider-vinyldns
WEBSITE_REPO=github.com/hashicorp/terraform-website
VINYLDNS_REPO=github.com/vinyldns/vinyldns
VERSION=0.10.3
VINYLDNS_VERSION=0.10.3

# Check that the required version of make is being used
REQ_MAKE_VER:=3.82
ifneq ($(REQ_MAKE_VER),$(firstword $(sort $(MAKE_VERSION) $(REQ_MAKE_VER))))
   $(error The version of MAKE $(REQ_MAKE_VER) or higher is required; you are running $(MAKE_VERSION))
endif

.ONESHELL:

all: start-api test build stop-api


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
	go install

.PHONY: build
build:
	go get github.com/mitchellh/gox
	GO111MODULE=on CGO_ENABLED=0 \
		gox \
			-ldflags "-X main.version=${VERSION}" \
			-osarch "darwin/amd64 freebsd/386 freebsd/amd64 freebsd/arm linux/386 linux/amd64 linux/arm linux/arm64 openbsd/386 openbsd/amd64 solaris/amd64 windows/386 windows/amd64" \
			-output "build/{{.OS}}_{{.Arch}}/terraform-provider-vinyldns_$(VERSION)"

.PHONY: version
version:
	echo $(VERSION)

.PHONY: website
website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

.PHONY: website-test
website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

.PHONY: package
package: build
	find release -not -name release -not -name '.dockerignore' -not -name '.gitignore' -print
	find release -not -name release -not -name '.dockerignore' -not -name '.gitignore' -delete
	for f in build/*; do \
		g=`basename $$f`; \
		zip --junk-paths release/$(NAME)_$(VERSION)_$${g}.zip build/$${g}/$(NAME)*; \
	done
	cd release && shasum -a 256 *.zip > $(NAME)_$(VERSION)_SHA256SUMS
	cd release && gpg \
		--detach-sign $(NAME)_$(VERSION)_SHA256SUMS

.PHONY: release
release: package
	go get github.com/aktau/github-release
	github-release release \
		--user vinyldns \
		--repo "${NAME}" \
		--target "$(shell git rev-parse --abbrev-ref HEAD)" \
		--tag "${VERSION}" \
		--name "${VERSION}"
	cd release && ls | xargs -I FILE github-release upload \
		--user vinyldns \
		--repo "${NAME}" \
		--tag "${VERSION}" \
		--name FILE \
		--file FILE

