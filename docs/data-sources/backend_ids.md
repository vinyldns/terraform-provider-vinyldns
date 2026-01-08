# vinyldns_backend_ids (Data Source)

Use this data source to list the available DNS backend IDs known to VinylDNS. These IDs identify DNS backends that zones can reference; this provider does not manage backends themselves.

## Example Usage

### List Available Backends

```hcl
data "vinyldns_backend_ids" "available" {}

output "available_backends" {
  value = data.vinyldns_backend_ids.available.backend_ids
}
```

### Check for a Specific Backend

```hcl
data "vinyldns_backend_ids" "available" {}

output "has_default_backend" {
  value = contains(data.vinyldns_backend_ids.available.backend_ids, "default")
}
```

## Argument Reference

This data source has no arguments.

## Attribute Reference

* `backend_ids` - A list of backend ID strings representing the DNS backends known to VinylDNS (backends are managed outside Terraform).
