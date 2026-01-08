# List all available backend IDs
# Backend IDs identify the DNS backends configured in VinylDNS
data "vinyldns_backend_ids" "available" {}

# Output the available backends
output "available_backends" {
  value       = data.vinyldns_backend_ids.available.backend_ids
  description = "List of DNS backend IDs configured in VinylDNS"
}

# Check if a specific backend exists
output "has_default_backend" {
  value = contains(data.vinyldns_backend_ids.available.backend_ids, "default")
}
