PKG_NAME=vinyldns
NAME=terraform-provider-vinyldns
WEBSITE_REPO=github.com/hashicorp/terraform-website
VINYLDNS_REPO=github.com/vinyldns/vinyldns
SOURCE=./...
VERSION=0.9.3

all: deps start-api test build stop-api
deps-build: deps build

deps:
	go get -u github.com/golang/dep/cmd/dep
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/mitchellh/gox
	dep ensure

start-api:
	if [ ! -d "$(GOPATH)/src/$(VINYLDNS_REPO)" ]; then \
		echo "$(VINYLDNS_REPO) not found in your GOPATH (necessary for acceptance tests), getting..."; \
		git clone https://$(VINYLDNS_REPO) $(GOPATH)/src/$(VINYLDNS_REPO); \
	fi
	$(GOPATH)/src/$(VINYLDNS_REPO)/bin/docker-up-vinyldns.sh \
		--version 0.9.1 \
		--api-only

stop-api:
	./../vinyldns/bin/remove-vinyl-containers.sh

# NOTE: acceptance tests assume a VinylDNS instance is running on localhost:9000 using the
# technique here: https://github.com/vinyldns/vinyldns/blob/master/bin/docker-up-vinyldns.sh
# See `start-api` for a convenience task in doing so.
test:
	go vet
	go test ${SOURCE} -v -cover
	VINYLDNS_ACCESS_KEY=okAccessKey \
		VINYLDNS_SECRET_KEY=okSecretKey \
		VINYLDNS_HOST=http://localhost:9000 \
		TF_LOG=DEBUG \
		TF_ACC=1 \
		go test ${SOURCE} -v

cover:
	go test $(TEST) -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

install: deps
	go install

build: deps
	export CGO_ENABLED=0; gox -ldflags "-X main.version=${VERSION}" -os "linux darwin windows" -arch "386 amd64" -output "build/{{.OS}}_{{.Arch}}/terraform-provider-vinyldns"

version:
	echo ${VERSION}

website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

package: build
	rm -rf release
	mkdir release
	for f in build/*; do \
		g=`basename $$f`; \
		tar -zcf release/$(NAME)-$${g}-$(VERSION).tgz -C build/$${g} .; \
	done

release: package
	go get github.com/aktau/github-release
	github-release release \
		--user vinyldns \
		--repo "${NAME}" \
		--target "$(shell git rev-parse --abbrev-ref HEAD)" \
		--tag "${VERSION}" \
		--name "${VERSION}"
	ls release/*.tgz | xargs -I FILE github-release upload \
		--user vinyldns \
		--repo "${NAME}" \
		--tag "${VERSION}" \
		--name FILE \
		--file FILE

.PHONY: deps run-api stop-api test cover install build version website website-test
