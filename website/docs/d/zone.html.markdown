---
page_title: "VinylDNS: vinyldns_zone"
sidebar_current: "docs-vinyldns-datasource-zone"
description: |-
  Get the zone details for a given zone name.
---

# vinyldns_zone

Use this data source to retrieve the `id`, `email`, and `admin_group_id` for a zone.

## Example Usage

```hcl
data "vinyldns_zone" "test" {
  name = "foo"
}
```

## Arguments Reference

* `name` - (Required) The name of the zone.

## Attributes Reference

* `id` - The ID of the zone

* `admin_group_id` - The ID of the zone's admin group

* `email` - The email associated with the zone
