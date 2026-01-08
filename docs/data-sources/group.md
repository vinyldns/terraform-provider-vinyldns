# vinyldns_group (Data Source)

Use this data source to look up an existing VinylDNS group by name.

## Example Usage

### Look Up a Group

```hcl
data "vinyldns_group" "dns_admins" {
  name = "dns-admins"
}
```

### Create a Zone with an Existing Group

```hcl
data "vinyldns_group" "existing" {
  name = "existing-admin-group"
}

resource "vinyldns_zone" "example" {
  name           = "example.com."
  email          = "hostmaster@example.com"
  admin_group_id = data.vinyldns_group.existing.id
}
```

### Output Group Details

```hcl
data "vinyldns_group" "example" {
  name = "my-group"
}

output "group_info" {
  value = {
    id           = data.vinyldns_group.example.id
    email        = data.vinyldns_group.example.email
    member_count = length(data.vinyldns_group.example.member_ids)
  }
}
```

## Argument Reference

* `name` - (Required) The exact name of the group to look up.

## Attribute Reference

* `id` - The unique identifier of the group.

* `email` - The email address associated with the group.

* `description` - The description of the group.

* `member_ids` - A set of user IDs who are members of the group.

* `admin_ids` - A set of user IDs who are administrators of the group.
