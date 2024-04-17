SOURCE?=./...
PKG_NAME=vinyldns
NAME=terraform-provider-vinyldns
WEBSITE_REPO=github.com/hashicorp/terraform-website
VINYLDNS_REPO=github.com/vinyldns/vinyldns
VINYLDNS_DIR="$(GOPATH)/src/$(VINYLDNS_REPO)"

ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

all: check-fmt test build integration stop-api validate-version install

fmt:
	gofmt -s -w vinyldns

check-fmt:
	test -z "$(shell gofmt -s -l vinyldns | tee /dev/stderr)"

test:
	go vet $(SOURCE)
	GO111MODULE=on go test $(SOURCE) -cover

integration: start-api
	VINYLDNS_ACCESS_KEY=okAccessKey \
	VINYLDNS_SECRET_KEY=okSecretKey \
	VINYLDNS_HOST=http://localhost:9000 \
	TF_LOG=ERROR \
	TF_ACC=1 \
	go test "$(SOURCE)" -v -cover

validate-version:
	cat vinyldns/version.go | grep 'var Version = "$(VERSION)"'

clonevinyl:
	if [ ! -d  $(VINYLDNS_DIR) ]; then \
		echo "$(VINYLDNS_REPO) not found in your GOPATH (necessary for acceptance tests), getting..."; \
		git clone \
			https://$(VINYLDNS_REPO) \
			$(VINYLDNS_DIR); \
	else \
		git -C $(VINYLDNS_DIR) pull ; \
	fi

start-api: clonevinyl stop-api
	$(GOPATH)/src/$(VINYLDNS_REPO)/quickstart/quickstart-vinyldns.sh \
		--api

stop-api:
	$(GOPATH)/src/$(VINYLDNS_REPO)/quickstart/quickstart-vinyldns.sh \
		--clean

build:
	GO111MODULE=on go build -ldflags "-X main.version=$(VERSION)" $(SOURCE)

install:
	GO111MODULE=on go install $(SOURCE)
	mkdir -p "$${HOME}/.terraform.d/plugins/local/vinyldns-provider/vinyldns/$(VERSION)/$$(go env GOOS)_$$(go env GOARCH)/"
	cp "$$(go env GOPATH)/bin/terraform-provider-vinyldns" "$${HOME}/.terraform.d/plugins/local/vinyldns-provider/vinyldns/$(VERSION)/$$(go env GOOS)_$$(go env GOARCH)/"

release: test validate-version
	go get github.com/aktau/github-release
	github-release release \
		--user vinyldns \
		--repo terraform-provider-vinyldns \
		--tag $(VERSION) \
		--name "$(VERSION)" \
		--description "terraform-provider-vinyldns version $(VERSION)"


website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(ROOT_DIR) PROVIDER_NAME=$(PKG_NAME)

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(ROOT_DIR) PROVIDER_NAME=$(PKG_NAME)