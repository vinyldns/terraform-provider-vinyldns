# vinyldns_record_set

Use this data source to retrieve the list of record sets present in a zone

## Example Usage

```hcl
data "vinyldns_record_set" "test" {
  zone_id   = "fbf7a440-891c-441a-ad09-e1cbc861sda2q"
}
```

## Arguments Reference

* `zone_id` - (Required) The id of the zone.

## Attributes Reference

* `recordset` - List of records present for the particular zone

