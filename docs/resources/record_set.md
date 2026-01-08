# vinyldns_record_set

Manages a VinylDNS record set. A record set represents one or more DNS records of the same type for a given name.

## Example Usage

### A Record

```hcl
resource "vinyldns_record_set" "web" {
  name             = "www"
  zone_id          = vinyldns_zone.example.id
  type             = "A"
  ttl              = 300
  record_addresses = ["192.0.2.1"]
}
```

### A Record with Multiple Addresses

```hcl
resource "vinyldns_record_set" "api" {
  name             = "api"
  zone_id          = vinyldns_zone.example.id
  type             = "A"
  ttl              = 60
  record_addresses = ["192.0.2.10", "192.0.2.11", "192.0.2.12"]
}
```

### AAAA Record (IPv6)

```hcl
resource "vinyldns_record_set" "web_ipv6" {
  name             = "www"
  zone_id          = vinyldns_zone.example.id
  type             = "AAAA"
  ttl              = 300
  record_addresses = ["2001:db8::1"]
}
```

### CNAME Record

```hcl
resource "vinyldns_record_set" "alias" {
  name         = "blog"
  zone_id      = vinyldns_zone.example.id
  type         = "CNAME"
  ttl          = 3600
  record_cname = "www.example.com."
}
```

### TXT Record

```hcl
resource "vinyldns_record_set" "spf" {
  name         = "example.com."
  zone_id      = vinyldns_zone.example.id
  type         = "TXT"
  ttl          = 3600
  record_texts = ["v=spf1 include:_spf.example.com ~all"]
}
```

### NS Record

```hcl
resource "vinyldns_record_set" "subdomain_ns" {
  name            = "subdomain"
  zone_id         = vinyldns_zone.example.id
  type            = "NS"
  ttl             = 86400
  record_nsdnames = ["ns1.subdomain.example.com.", "ns2.subdomain.example.com."]
}
```

### PTR Record

```hcl
resource "vinyldns_record_set" "ptr" {
  name             = "1"
  zone_id          = vinyldns_zone.reverse.id
  type             = "PTR"
  ttl              = 3600
  record_ptrdnames = ["www.example.com."]
}
```

### Record with Owner Group

In shared zones, records can be assigned to an owner group:

```hcl
resource "vinyldns_record_set" "owned" {
  name             = "app"
  zone_id          = vinyldns_zone.shared.id
  type             = "A"
  ttl              = 300
  record_addresses = ["192.0.2.50"]
  owner_group_id   = vinyldns_group.app_team.id
}
```

## Argument Reference

* `name` - (Required, Forces new resource) The name of the record set. Use the zone name (with trailing dot) for apex records, or just the subdomain name for other records.

* `zone_id` - (Required, Forces new resource) The ID of the zone this record set belongs to.

* `type` - (Required, Forces new resource) The DNS record type. Supported types: `A`, `AAAA`, `CNAME`, `TXT`, `NS`, `PTR`.

* `ttl` - (Optional) The time-to-live in seconds.

* `owner_group_id` - (Optional) The ID of the group that owns this record. Used in shared zones for record ownership.

* `record_addresses` - (Optional) A set of IP addresses. Used for `A` and `AAAA` record types.

* `record_cname` - (Optional) The canonical name. Used for `CNAME` record type. Must end with a trailing dot.

* `record_texts` - (Optional) A set of text values. Used for `TXT` record type.

* `record_nsdnames` - (Optional) A set of nameserver names. Used for `NS` record type.

* `record_ptrdnames` - (Optional) A set of pointer domain names. Used for `PTR` record type. Must end with a trailing dot.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The unique identifier of the record set (format: `zone_id:record_set_id`).

## Import

Record sets can be imported using a combination of zone ID and record set ID:

```shell
terraform import vinyldns_record_set.example 9cbdd3ac-9752-4d56-9ca0-6a1a14fc5562:8306cce4-e16a-4579-9b19-4af46dc75853
```

## Notes

* CNAME and PTR record values must end with a trailing dot (e.g., `www.example.com.`)
* SOA records are read-only and cannot be managed through this provider
* Changing `name`, `zone_id`, or `type` will force creation of a new resource
