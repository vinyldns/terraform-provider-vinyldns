# List all zones
data "vinyldns_zones" "all" {}

# List zones matching a name filter
data "vinyldns_zones" "production" {
  name_filter = "prod"
}

# Output all zone names
output "all_zone_names" {
  value = [for z in data.vinyldns_zones.all.zones : z.name]
}

# Output production zones with their details
output "production_zones" {
  value = [
    for z in data.vinyldns_zones.production.zones : {
      name      = z.name
      status    = z.status
      shared    = z.shared
    }
  ]
}

# Find zones by status
output "active_zones" {
  value = [
    for z in data.vinyldns_zones.all.zones : z.name
    if z.status == "Active"
  ]
}
