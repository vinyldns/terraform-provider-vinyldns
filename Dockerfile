FROM golang:1.17
COPY . /go/src/github.com/vinyldns/terraform-provider-vinyldns
WORKDIR /go/src/github.com/vinyldns/terraform-provider-vinyldns
CMD ["make", "package"]
