---
page_title: "Provider: VinylDNS"
sidebar_current: "docs-vinyldns-index"
description: |-
  The VinylDNS provider configures VinylDNS resources.
---

# VinylDNS Provider

The VinylDNS provider configures [VinylDNS](https://www.vinyldns.io/) resources.
VinylDNS is a vendor-agnostic DNS front-end for streamlining DNS operations and
enabling self-service for your DNS infrastructure.

The provider configuration block accepts the following arguments:

* ``host`` - (Required) The root URL of a VinylDNS API server. May alternatively be
  set via the ``VINYLDNS_HOST`` environment variable.

* ``access_key`` - (Required) The access key required to authenticate to the
	VinylDNS server. May alternatively be set via the ``VINYLDNS_ACCESS_KEY``
	environment variable.

* ``secret_key`` - (Required) The secret key required to authenticate to the
	VinylDNS server. May alternatively be set via the ``VINYLDNS_SECRET_KEY``
	environment variable.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
provider "vinyldns" {
  host        = "http://vinyldns.example.com"
  access_key  = "123"
  secret_key  = "123"
}

resource "vinyldns_group" "test_group" {
  name = "terraform-provider-test-group"
}

resource "vinyldns_zone" "test_zone" {
  name = "system-test."
  email = "foo@bar.com"
  admin_group_id = "${vinyldns_group.test_group.id}"
  zone_connection {
    name = "vinyldns."
    key_name = "vinyldns."
    key = "123"
    primary_server = "127.0.0.1"
  }
}

resource "vinyldns_record_set" "test_record_set" {
  name = "terraformtestrecordset"
  zone_id = "${vinyldns_zone.test_zone.id}"
  type = "A"
  ttl = 6000
  record_addresses = ["127.0.0.1"]
}

resource "vinyldns_record_set" "another_test_record_set" {
  name = "another-terraformtestrecordset"
  zone_id = "${vinyldns_zone.test_zone.id}"
  type = "CNAME"
  ttl = 6000
  record_cname = "foo-bar.com."
}
```
