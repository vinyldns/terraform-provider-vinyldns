FROM golang:1.10.3
COPY . /go/src/github.com/vinyldns/terraform-provider-vinyldns
WORKDIR /go/src/github.com/vinyldns/terraform-provider-vinyldns
RUN make build
