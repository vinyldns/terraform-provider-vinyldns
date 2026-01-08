# vinyldns_groups (Data Source)

Use this data source to list VinylDNS groups, optionally filtered by name.

## Example Usage

### List All Groups

```hcl
data "vinyldns_groups" "all" {}

output "group_names" {
  value = [for g in data.vinyldns_groups.all.groups : g.name]
}
```

### List Groups Matching a Filter

```hcl
data "vinyldns_groups" "dns_teams" {
  name_filter = "dns"
}

output "dns_teams" {
  value = [
    for g in data.vinyldns_groups.dns_teams.groups : {
      name         = g.name
      email        = g.email
      member_count = length(g.member_ids)
    }
  ]
}
```

## Argument Reference

* `name_filter` - (Optional) A substring filter for group names. VinylDNS performs a contains search; wildcards and regex are not supported.

## Attribute Reference

* `groups` - A list of groups. Each group has the following attributes:

  * `id` - The unique identifier of the group.

  * `name` - The name of the group.

  * `email` - The email address associated with the group.

  * `description` - The description of the group.

  * `member_ids` - A set of user IDs who are members of the group.

  * `admin_ids` - A set of user IDs who are administrators of the group.
