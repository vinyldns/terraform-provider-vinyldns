# First, look up the zone
data "vinyldns_zone" "example" {
  name = "example.com."
}

# List all record sets in a zone
data "vinyldns_record_sets" "all" {
  zone_id = data.vinyldns_zone.example.id
}

# List record sets matching a name filter
data "vinyldns_record_sets" "api_records" {
  zone_id     = data.vinyldns_zone.example.id
  name_filter = "api"
}

# Output all record names and types
output "all_records" {
  value = [
    for rs in data.vinyldns_record_sets.all.record_sets : {
      name = rs.name
      fqdn = rs.fqdn
      type = rs.type
      ttl  = rs.ttl
    }
  ]
}

# Filter by record type in Terraform
output "a_records" {
  value = [
    for rs in data.vinyldns_record_sets.all.record_sets : rs.fqdn
    if rs.type == "A"
  ]
}

output "cname_records" {
  value = [
    for rs in data.vinyldns_record_sets.all.record_sets : rs.fqdn
    if rs.type == "CNAME"
  ]
}

# Output API record details
output "api_records" {
  value = data.vinyldns_record_sets.api_records.record_sets
}
