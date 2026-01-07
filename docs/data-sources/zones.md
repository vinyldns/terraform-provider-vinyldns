# vinyldns_zones

Use this data source to list zones, optionally filtered by name.

## Example Usage

```hcl
data "vinyldns_zones" "all" {}

data "vinyldns_zones" "filtered" {
  name_filter = "prod."
}
```

## Arguments Reference

* `name_filter` - (Optional) Filter zones by name.

## Attributes Reference

* `zones` - List of matching zones. Each zone includes:
  * `id` - The zone ID.
  * `name` - The zone name.
  * `email` - The zone email.
  * `admin_group_id` - The admin group ID.
  * `status` - The zone status.
  * `shared` - Whether the zone is shared.
  * `backend_id` - The backend ID.
