# vinyldns_group

Manages a VinylDNS group. Groups are used to control access to zones and can be assigned as zone administrators.

## Example Usage

### Basic Group

```hcl
resource "vinyldns_group" "example" {
  name        = "example-group"
  email       = "dns-team@example.com"
  description = "Example VinylDNS group"
  member_ids  = ["user-id-1", "user-id-2"]
  admin_ids   = ["user-id-1"]
}
```

### Group Where All Members Are Admins

```hcl
resource "vinyldns_group" "admins" {
  name        = "dns-admins"
  email       = "dns-admins@example.com"
  description = "DNS administrators"
  member_ids  = ["admin-1", "admin-2", "admin-3"]
  admin_ids   = ["admin-1", "admin-2", "admin-3"]
}
```

## Argument Reference

* `name` - (Required) The name of the group.

* `email` - (Required) The email address associated with the group.

* `description` - (Optional) A description of the group. Defaults to "Managed by Terraform".

* `member_ids` - (Required) A set of user IDs who are members of the group.

* `admin_ids` - (Required) A set of user IDs who are administrators of the group. Admin IDs should also be included in `member_ids`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The unique identifier of the group.

## Import

Groups can be imported using their ID:

```shell
terraform import vinyldns_group.example 6f8afcda-7529-4cad-9f2d-76903f4b1aca
```
