[![Terraform Registry](https://img.shields.io/github/v/release/vinyldns/terraform-provider-vinyldns?color=834FB9&label=registry&logo=terraform)](https://registry.terraform.io/providers/vinyldns/vinyldns/latest)
[![Build Status](https://github.com/vinyldns/terraform-provider-vinyldns/actions/workflows/release.yml/badge.svg)](https://github.com/vinyldns/terraform-provider-vinyldns/actions/workflows/release.yml)
[![GitHub](https://img.shields.io/github/license/vinyldns/terraform-provider-vinyldns)](https://github.com/vinyldns/vinyldns/blob/master/LICENSE)

# terraform-provider-vinyldns

A [Terraform](https://terraform.io) provider for the [VinylDNS](https://github.com/vinyldns/vinyldns) DNS as a service
API.

* [Terraform](http://terraform.io)
* [VinylDNS](https://www.vinyldns.io)

See [example.tf](https://github.com/vinyldns/terraform-provider-vinyldns/blob/master/example.tf) for an example `.tf`
file.

See https://vinyldns.github.io/terraform-provider-vinyldns for documentation.

## Installation

1. Create a `providers.tf` file and add the `vinyldns` provider

```hcl
terraform {
  required_providers {
    vinyldns = {
      source  = "vinyldns/vinyldns"
      version = "0.10.0"
    }
  }
}
```

### Installing from source

Alternatively, you can install from source:

```shell script
$ git clone https://github.com/vinyldns/terraform-provider-vinyldns.git
$ cd terraform-provider-vinyldns
$ make install
```

Add the VinylDNS provider to `providers.tf` using the local path. Note that the locally installed version will always
be `0.0.1` so as not to confuse it with the version released to
the [Terraform Registry](https://registry.terraform.io/providers/vinyldns/vinyldns/latest).

```hcl
terraform {
  required_providers {
    vinyldns = {
      source  = "local/vinyldns-provider/vinyldns"
      version = "0.0.1"
    }
  }
}
```

## Running acceptance tests

The `terraform-provider-vinyldns` acceptance tests assume a VinylDNS API is running on `localhost:9000`.

This will be done automatically for you via `make test`. Note that you must have `Docker` installed and running.

```shell script
$ git clone https://github.com/vinyldns/terraform-provider-vinyldns.git
$ cd vinyldns
$ make test
```

## Building

To build `terraform-provider-vinyldns` binaries for your current platform:

```shell script
$ make build
```

## Credits

`terraform-provider-vinyldns` would not be possible without the help of many other pieces of open source software. Thank
you open source world!

Given the Apache 2.0 license of `terraform-provider-vinyldns`, we specifically want to call out the following packages
and their corresponding licenses:

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
