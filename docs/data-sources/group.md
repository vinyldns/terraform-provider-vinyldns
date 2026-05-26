# vinyldns_group

Use this data source to retrieve group details by name.

## Example Usage

```hcl
data "vinyldns_group" "admins" {
  name = "my-admins"
}
```

## Arguments Reference

* `name` - (Required) The name of the group.

## Attributes Reference

* `id` - The ID of the group.
* `email` - The email associated with the group.
* `description` - The group description.
* `member_ids` - The member user IDs.
* `admin_ids` - The admin user IDs.
