# vinyldns_zones (Data Source)

Use this data source to list VinylDNS zones, optionally filtered by name.

## Example Usage

### List All Zones

```hcl
data "vinyldns_zones" "all" {}

output "zone_names" {
  value = [for z in data.vinyldns_zones.all.zones : z.name]
}
```

### List Zones Matching a Filter

```hcl
data "vinyldns_zones" "production" {
  name_filter = "prod"
}

output "production_zones" {
  value = data.vinyldns_zones.production.zones
}
```

### Find Active Zones

```hcl
data "vinyldns_zones" "all" {}

output "active_zones" {
  value = [
    for z in data.vinyldns_zones.all.zones : z.name
    if z.status == "Active"
  ]
}
```

## Argument Reference

* `name_filter` - (Optional) A zone name filter. VinylDNS supports `*` wildcards (e.g., `prod*`, `*corp*`). Without a wildcard, VinylDNS treats this as a prefix match unless you provide the full zone name with a trailing dot for an exact match.

## Attribute Reference

* `zones` - A list of zones. Each zone has the following attributes:

  * `id` - The unique identifier of the zone.

  * `name` - The name of the zone.

  * `email` - The email address associated with the zone.

  * `admin_group_id` - The ID of the group that administers the zone.

  * `status` - The zone status (e.g., `Active`, `Syncing`).

  * `shared` - Whether the zone is a shared zone.

  * `backend_id` - The ID of the DNS backend for this zone.
