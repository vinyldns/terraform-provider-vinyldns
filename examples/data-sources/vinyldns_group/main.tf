# Look up an existing group by name
data "vinyldns_group" "dns_admins" {
  name = "dns-admins"
}

# Use the group to create a zone
resource "vinyldns_zone" "managed_by_existing_group" {
  name           = "managed.example.com."
  email          = "dns@example.com"
  admin_group_id = data.vinyldns_group.dns_admins.id
}

# Output group details
output "group_id" {
  value = data.vinyldns_group.dns_admins.id
}

output "group_email" {
  value = data.vinyldns_group.dns_admins.email
}

output "group_member_count" {
  value = length(data.vinyldns_group.dns_admins.member_ids)
}

output "group_admin_ids" {
  value = data.vinyldns_group.dns_admins.admin_ids
}
