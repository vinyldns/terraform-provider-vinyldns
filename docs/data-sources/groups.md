# vinyldns_groups

Use this data source to list groups, optionally filtered by name.

## Example Usage

```hcl
data "vinyldns_groups" "all" {}

data "vinyldns_groups" "filtered" {
  name_filter = "admins"
}
```

## Arguments Reference

* `name_filter` - (Optional) Filter groups by name.

## Attributes Reference

* `groups` - List of matching groups. Each group includes:
  * `id` - The group ID.
  * `name` - The group name.
  * `email` - The email associated with the group.
  * `description` - The group description.
  * `member_ids` - The member user IDs.
  * `admin_ids` - The admin user IDs.
