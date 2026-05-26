# vinyldns_record_sets

Use this data source to list record sets for a zone, optionally filtered by name.

## Example Usage

```hcl
data "vinyldns_record_sets" "zone_records" {
  zone_id     = "zone-id"
  name_filter = "www"
}
```

## Arguments Reference

* `zone_id` - (Required) The zone ID to query.
* `name_filter` - (Optional) Filter record sets by name.

## Attributes Reference

* `record_sets` - List of matching record sets. Each record set includes:
  * `id` - The record set ID.
  * `name` - The record set name.
  * `fqdn` - The record set FQDN.
  * `type` - The record set type.
  * `ttl` - The record set TTL.
  * `owner_group_id` - The owner group ID.
  * `status` - The record set status.
