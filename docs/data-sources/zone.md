# vinyldns_zone (Data Source)

Use this data source to look up an existing VinylDNS zone by name.

## Example Usage

### Look Up a Zone

```hcl
data "vinyldns_zone" "example" {
  name = "example.com."
}
```

### Create a Record in an Existing Zone

```hcl
data "vinyldns_zone" "example" {
  name = "example.com."
}

resource "vinyldns_record_set" "new_record" {
  name             = "new-host"
  zone_id          = data.vinyldns_zone.example.id
  type             = "A"
  ttl              = 300
  record_addresses = ["192.0.2.100"]
}
```

## Argument Reference

* `name` - (Required) The name of the zone to look up.

## Attribute Reference

* `id` - The unique identifier of the zone.

* `email` - The email address associated with the zone.

* `admin_group_id` - The ID of the group that administers the zone.
