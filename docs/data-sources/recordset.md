# vinyldns_group

Use this data source to retrieve the `name`, `email`, and `description` for a group.

## Example Usage

```hcl
data "vinyldns_record_set" "test" {
  zone_id   = "foo"
  record_id = "abc"
}
```

## Arguments Reference

* `zone_id` - (Required) The id of the zone.
* `record_id` - (Required) The id of the record.

## Attributes Reference

* `record_name` - The name of the record

* `record_type` - The type of the record

* `TTL` - The Time to Live (TTL) of the record

* `record_data` - The data associated with the record

* `zone_name` - The zone associated with the record

* `zone_access_type` - The access type of the zone associated with the record

* `owner_group_name` - The owner group name associated with the record

