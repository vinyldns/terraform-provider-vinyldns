---
layout: "vinyldns"
page_title: "VinylDNS: vinyldns_record_set"
sidebar_current: "docs-vinyldns-resource-record-set"
description: |-
  The vinyldns_record_set resource allows a VinylDNS record set to be created and managed.
---

# vinyldns\_record_set

The record set resource allows VinylDNS record sets to be created and managed.

## Example Usage

```hcl
resource "vinyldns_record_set" "test_record_set" {
  name = "terraformtestrecordset"
  zone_id = "123"
  type = "A"
  ttl = 6000
  record_addresses = ["127.0.0.1"]
}

resource "vinyldns_record_set" "another_test_record_set" {
  name = "another-terraformtestrecordset"
  zone_id = "123"
  type = "CNAME"
  ttl = 6000
  record_cname = "foo-bar.com."
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name for the record set.

* `zone_id` - (Required) The ID for the record set's zone.

* `owner_group_id` - (Optional) Record ownership assignment. This is applicable if the recordset exists in a shared zone.

* `type` - (Required) The type of DNS record.

* `ttl` - (Optional) The DNS record set's TTL, or time to live.

* `record_addresses` - (Optional) A list of the record set's addresses.

* `record_texts` - (Optional) If the record is a TXT record, a list of the record set's text values.

* `record_nsdnames` - (Optional) If the record is an NS record, a list of the record set's nsdname values.

* `record_ptrdnames` - (Optional) If the record is a PTR record, a list of the record set's ptrdname values.

* `record_cname` - (Optional) If the record is a CNAME, the record's value.

## Import

`vinyldns_record_set` can be imported using a combination of the ID of the zone and the ID of the record set, e.g.

```
terraform import vinyldns_record_set.example zone_id:record_set_id
```
