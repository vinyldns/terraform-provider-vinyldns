FROM golang:1.13
COPY . /go/src/github.com/vinyldns/terraform-provider-vinyldns
WORKDIR /go/src/github.com/vinyldns/terraform-provider-vinyldns
CMD ["make", "build-deps", "package"]
