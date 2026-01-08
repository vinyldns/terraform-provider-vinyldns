# vinyldns_record_sets (Data Source)

Use this data source to list DNS record sets in a zone, optionally filtered by name.

## Example Usage

### List All Record Sets in a Zone

```hcl
data "vinyldns_zone" "example" {
  name = "example.com."
}

data "vinyldns_record_sets" "all" {
  zone_id = data.vinyldns_zone.example.id
}

output "all_records" {
  value = [
    for rs in data.vinyldns_record_sets.all.record_sets : {
      name = rs.name
      type = rs.type
      fqdn = rs.fqdn
    }
  ]
}
```

### Filter Record Sets by Name

```hcl
data "vinyldns_record_sets" "api_records" {
  zone_id     = data.vinyldns_zone.example.id
  name_filter = "api"
}

output "api_records" {
  value = data.vinyldns_record_sets.api_records.record_sets
}
```

### Filter by Record Type in Terraform

```hcl
data "vinyldns_record_sets" "all" {
  zone_id = data.vinyldns_zone.example.id
}

output "a_records" {
  value = [
    for rs in data.vinyldns_record_sets.all.record_sets : rs.fqdn
    if rs.type == "A"
  ]
}

output "cname_records" {
  value = [
    for rs in data.vinyldns_record_sets.all.record_sets : rs.fqdn
    if rs.type == "CNAME"
  ]
}
```

## Argument Reference

* `zone_id` - (Required) The ID of the zone to list record sets from.

* `name_filter` - (Optional) A record name filter within the zone. VinylDNS supports `*` wildcards (e.g., `api*`). Without a wildcard, VinylDNS treats this as an exact record name match.

## Attribute Reference

* `record_sets` - A list of record sets. Each record set has the following attributes:

  * `id` - The unique identifier of the record set.

  * `name` - The name of the record set (relative to the zone).

  * `fqdn` - The fully qualified domain name of the record set.

  * `type` - The DNS record type (e.g., `A`, `AAAA`, `CNAME`).

  * `ttl` - The time-to-live in seconds.

  * `owner_group_id` - The ID of the group that owns this record (for shared zones).

  * `status` - The record set status.
