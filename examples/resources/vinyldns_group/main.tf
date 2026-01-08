# Basic group
resource "vinyldns_group" "example" {
  name        = "example-group"
  email       = "group-contact@example.com"
  description = "Example VinylDNS group"
  member_ids  = ["user-id-1", "user-id-2"]
  admin_ids   = ["user-id-1"]
}

# Group with all members as admins
resource "vinyldns_group" "admins_only" {
  name        = "dns-admins"
  email       = "dns-admins@example.com"
  description = "DNS administrators group"
  member_ids  = ["admin-1", "admin-2", "admin-3"]
  admin_ids   = ["admin-1", "admin-2", "admin-3"]
}
