PKG_NAME=vinyldns
NAME=terraform-provider-vinyldns
WEBSITE_REPO=github.com/hashicorp/terraform-website
VINYLDNS_REPO=github.com/vinyldns/vinyldns
SOURCE=./...
VERSION=0.9.5
VINYLDNS_VERSION=0.9.3

all: start-api test build stop-api

start-api:
	if [ ! -d "$(GOPATH)/src/$(VINYLDNS_REPO)-$(VINYLDNS_VERSION)" ]; then \
		echo "$(VINYLDNS_REPO)-$(VINYLDNS_VERSION) not found in your GOPATH (necessary for acceptance tests), getting..."; \
		git clone \
			--branch v$(VINYLDNS_VERSION) \
			https://$(VINYLDNS_REPO) \
			$(GOPATH)/src/$(VINYLDNS_REPO)-$(VINYLDNS_VERSION); \
	fi
	$(GOPATH)/src/$(VINYLDNS_REPO)-$(VINYLDNS_VERSION)/bin/docker-up-vinyldns.sh \
		--api-only \
		--version $(VINYLDNS_VERSION)

stop-api:
	$(GOPATH)/src/$(VINYLDNS_REPO)-$(VINYLDNS_VERSION)/bin/remove-vinyl-containers.sh

# NOTE: acceptance tests assume a VinylDNS instance is running on localhost:9000 using the
# technique here: https://github.com/vinyldns/vinyldns/blob/master/bin/docker-up-vinyldns.sh
# See `start-api` for a convenience task in doing so.
test:
	GO111MODULE=on go vet "${SOURCE}"
	GO111MODULE=on go test ${SOURCE} -cover
	VINYLDNS_ACCESS_KEY=okAccessKey \
		VINYLDNS_SECRET_KEY=okSecretKey \
		VINYLDNS_HOST=http://localhost:9000 \
		TF_LOG=DEBUG \
		TF_ACC=1 \
		go test ${SOURCE} -v

install:
	GO111MODULE=on go install

build:
	go get github.com/mitchellh/gox
	GO111MODULE=on CGO_ENABLED=0 \
		gox \
			-ldflags "-X main.version=${VERSION}" \
			-osarch "darwin/amd64 freebsd/386 freebsd/amd64 freebsd/arm linux/386 linux/amd64 linux/arm linux/arm64 openbsd/386 openbsd/amd64 solaris/amd64 windows/386 windows/amd64" \
			-output "build/{{.OS}}_{{.Arch}}/terraform-provider-vinyldns-$(VERSION)"

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

# NOTE: this currently uses a mikedball@gmail.com GPG key,
# as that key has been uploaded to registry.terraform.io
# TODO: this should ideally use a dedicated key
package: build
	rm -rf release
	mkdir release
	for f in build/*; do \
		g=`basename $$f`; \
		zip --junk-paths release/$(NAME)-$(VERSION)-$${g}.zip build/$${g}/$(NAME)*; \
		tar -zcf release/$(NAME)-$${g}-$(VERSION).tgz -C build/$${g} .; \
	done
	cd release && shasum -a 256 *.zip > $(NAME)-$(VERSION)-SHASUMS
	cd release && gpg \
		--armor \
		--detach-sign $(NAME)-$(VERSION)-SHASUMS

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

.PHONY: run-api stop-api test cover install build version website website-test
