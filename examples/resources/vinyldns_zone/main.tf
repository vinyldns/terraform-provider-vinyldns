# Basic zone (requires a group to own it)
resource "vinyldns_group" "zone_admin" {
  name       = "zone-admin-group"
  email      = "dns-team@example.com"
  member_ids = ["user-id-1"]
  admin_ids  = ["user-id-1"]
}

resource "vinyldns_zone" "example" {
  name           = "example.com."
  email          = "dns-admin@example.com"
  admin_group_id = vinyldns_group.zone_admin.id
}

# Zone with connection details (for zones requiring TSIG authentication)
resource "vinyldns_zone" "with_connection" {
  name           = "secure.example.com."
  email          = "dns-admin@example.com"
  admin_group_id = vinyldns_group.zone_admin.id

  zone_connection {
    name           = "secure.example.com."
    key_name       = "tsig-key."
    key            = "base64-encoded-tsig-key"
    primary_server = "ns1.example.com"
  }
}

# Zone with transfer connection (for zone transfers from a different server)
resource "vinyldns_zone" "with_transfer" {
  name           = "transferred.example.com."
  email          = "dns-admin@example.com"
  admin_group_id = vinyldns_group.zone_admin.id

  zone_connection {
    name           = "transferred.example.com."
    key_name       = "update-key."
    key            = "base64-encoded-update-key"
    primary_server = "ns1.example.com"
  }

  transfer_connection {
    name           = "transferred.example.com."
    key_name       = "transfer-key."
    key            = "base64-encoded-transfer-key"
    primary_server = "ns2.example.com"
  }
}

# Zone with ACL rules
resource "vinyldns_zone" "with_acl" {
  name           = "restricted.example.com."
  email          = "dns-admin@example.com"
  admin_group_id = vinyldns_group.zone_admin.id

  # Grant read access to a specific group for all record types
  acl_rule {
    access_level = "Read"
    group_id     = "reader-group-id"
    description  = "Read access for monitoring team"
  }

  # Grant write access for A and AAAA records only
  acl_rule {
    access_level = "Write"
    group_id     = "web-team-group-id"
    record_types = ["A", "AAAA"]
    description  = "Web team can manage A/AAAA records"
  }

  # Grant write access for specific record name pattern
  acl_rule {
    access_level = "Write"
    group_id     = "api-team-group-id"
    record_mask  = "api-.*"
    record_types = ["A", "CNAME"]
    description  = "API team can manage api-* records"
  }

  # Grant access to a specific user
  acl_rule {
    access_level = "Delete"
    user_id      = "specific-user-id"
    description  = "Full access for specific user"
  }
}
