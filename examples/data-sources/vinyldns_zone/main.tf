# Look up an existing zone by name
data "vinyldns_zone" "example" {
  name = "example.com."
}

# Use the zone data to create a record
resource "vinyldns_record_set" "in_existing_zone" {
  name             = "new-record"
  zone_id          = data.vinyldns_zone.example.id
  type             = "A"
  ttl              = 300
  record_addresses = ["192.0.2.100"]
}

# Output zone details
output "zone_id" {
  value = data.vinyldns_zone.example.id
}

output "zone_email" {
  value = data.vinyldns_zone.example.email
}

output "zone_admin_group_id" {
  value = data.vinyldns_zone.example.admin_group_id
}
