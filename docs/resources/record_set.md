# vinyldns\_record_set

The record set resource allows VinylDNS recordsets to be created and managed.

## Example Usage

```hcl
resource "vinyldns_record_set" "test_record_set" {
  name             = "terraformtestrecordset"
  zone_id          = "123"
  type             = "A"
  ttl              = 6000
  record_addresses = ["127.0.0.1"]
}

resource "vinyldns_record_set" "another_test_record_set" {
  name         = "another-terraformtestrecordset"
  zone_id      = "123"
  type         = "CNAME"
  ttl          = 6000
  record_cname = "foo-bar.com."
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name for the record set.

* `zone_id` - (Required) The ID for the record set's zone.

* `owner_group_id` - (Optional) Record ownership assignment. This is applicable if the record set exists in a shared zone.

* `type` - (Required) The type of DNS record.

* `ttl` - (Optional) The DNS record set's TTL, or time to live.

* `record_addresses` - (Optional) A list of the record set's addresses.

* `record_texts` - (Optional) If the record is a TXT record, a list of the record set's text values.

* `record_nsdnames` - (Optional) If the record is an NS record, a list of the record set's nsdname values.

* `record_ptrdnames` - (Optional) If the record is a PTR record, a list of the record set's ptrdname values.

* `record_cname` - (Optional) If the record is a CNAME, the record's value.

## Import

`vinyldns_record_set` can be imported using a combination of its zone ID and record set ID. For example, run the following command to import a record set with ID `8306cce4-e16a-4579-9b19-4af46dc75853` from a zone with ID `9cbdd3ac-9752-4d56-9ca0-6a1a14fc5562`:

```
terraform import vinyldns_record_set.example 9cbdd3ac-9752-4d56-9ca0-6a1a14fc5562:8306cce4-e16a-4579-9b19-4af46dc75853
```
