# List all groups
data "vinyldns_groups" "all" {}

# List groups matching a name filter
data "vinyldns_groups" "dns_teams" {
  name_filter = "dns"
}

# Output all group names
output "all_group_names" {
  value = [for g in data.vinyldns_groups.all.groups : g.name]
}

# Output DNS team groups with member counts
output "dns_teams" {
  value = [
    for g in data.vinyldns_groups.dns_teams.groups : {
      name         = g.name
      email        = g.email
      member_count = length(g.member_ids)
      admin_count  = length(g.admin_ids)
    }
  ]
}
