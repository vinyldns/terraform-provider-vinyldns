# terraform-provider-vinyldns

A [Terraform](https://terraform.io) provider for the [VinylDNS](https://github.com/vinyldns/vinyldns) DNS as a service API.

* [Terraform](http://terraform.io)
* [VinylDNS](https://www.vinyldns.io)

See `example.tf` for an example `.tf` file. See `website/docs` for documentation.

## Installation

1. Download the desired release version for your operating system from [GitHub](https://github.com/vinyldns/terraform-provider-vinyldns/releases).
2. Untar the download contents
3. Install the `terraform-provider-vinyldns` anywhere on your system
4. Add `terraform-provider-vinyldns` to your `~/.terraformrc` file:

```
providers {
  "vinyldns" = "path/to/your/terraform-provider-vinyldns"
}
```

### Installing from source

Alternatively, you can install from source:

* install Golang (1.10 currently required)
* establish your `$GOPATH`
* clone `vinyldns/terraform-provider-vinyldns` to `/$GOPATH/src/github.com/vinyldns/terraform-provider-vinyldns`
* `cd /$GOPATH/github.com/vinyldns/terraform-provider-vinyldns && make`
* Add the following to your `~/.terraformrc`:

```
providers {
  "vinyldns" = "path/to/your/terraform-provider-vinyldns"
}
```

## Running acceptance tests

The `terraform-provider-vinyldns` acceptance tests assume a VinylDNS API is running on `localhost:9000`.

To run a local VinylDNS API, you'll need to:

```
git clone git@github.com:vinyldns/vinyldns.git
cd vinyldns
bin/docker-up-vinyldns.sh
```

Note that a `make` convenience task handles this:

```
make start-api
```

Then, to run the `terraform-provider-vinyldns` acceptance tests against the local Dockerized VinylDNS API server:

```
make test
```

To stop the `localhost:9000` VinylDNS:

```
make stop-api
```

## Building

To build `terraform-provider-vinyldns` binaries for all supported platforms:

```
make build
```

### Building in Docker

The project contains a `docker-compose.yml`/`Dockerfile` that will perform a test build in a empty container. To run:

```
docker-compose build
```

## Upgrading Dependencies

`dep` is used to manage dependencies. To require a specific version of `github.com/vinyldns/go-vinyldns`:

To add a dependency:

```
dep ensure -add github.com/pkg/errors
```

## Credits

`terraform-provider-vinyldns` would not be possible without the help of many other pieces of open source software. Thank you open source world!

Given the Apache 2.0 license of `terraform-provider-vinyldns`, we specifically want to call out the following packages and their corresponding licenses:

* [github.com/hashicorp/errwrap](https://github.com/hashicorp/errwrap) - Mozilla Public License 2.0
* [github.com/hashicorp/go-getter](https://github.com/hashicorp/go-getter) - Mozilla Public License 2.0
* [github.com/hashicorp/go-multierror](https://github.com/hashicorp/go-multierror) - Mozilla Public License 2.0
* [github.com/hashicorp/go-plugin](https://github.com/hashicorp/go-plugin) - Mozilla Public License 2.0
* [github.com/hashicorp/go-uuid](https://github.com/hashicorp/go-uuid) - Mozilla Public License 2.0
* [github.com/hashicorp/go-version](https://github.com/hashicorp/go-version) - Mozilla Public License 2.0
* [github.com/hashicorp/hcl](https://github.com/hashicorp/hcl) - Mozilla Public License 2.0
* [github.com/hashicorp/hil](https://github.com/hashicorp/hil) - Mozilla Public License 2.0
* [github.com/hashicorp/logutils](https://github.com/hashicorp/logutils) - Mozilla Public License 2.0
* [github.com/hashicorp/terraform](github.com/hashicorp/terraform) - Mozilla Public License 2.0
* [github.com/hashicorp/yamux](https://github.com/hashicorp/yamux) - Mozilla Public License 2.0
