# Basic Zone Setup
#
# This example demonstrates the minimum configuration needed to:
# 1. Create a group to administer the zone
# 2. Create a zone
# 3. Add DNS records to the zone

terraform {
  required_providers {
    vinyldns = {
      source = "vinyldns/vinyldns"
    }
  }
}

provider "vinyldns" {
  # Configure via environment variables:
  #   VINYLDNS_HOST
  #   VINYLDNS_ACCESS_KEY
  #   VINYLDNS_SECRET_KEY
}

# Step 1: Create a group to own the zone
resource "vinyldns_group" "zone_owners" {
  name        = "example-zone-owners"
  email       = "dns-team@example.com"
  description = "Owners of the example.com zone"
  member_ids  = var.group_member_ids
  admin_ids   = var.group_admin_ids
}

# Step 2: Create the zone
resource "vinyldns_zone" "example" {
  name           = "example.com."
  email          = "hostmaster@example.com"
  admin_group_id = vinyldns_group.zone_owners.id
}

# Step 3: Add DNS records
resource "vinyldns_record_set" "root_a" {
  name             = "example.com."
  zone_id          = vinyldns_zone.example.id
  type             = "A"
  ttl              = 300
  record_addresses = ["192.0.2.1"]
}

resource "vinyldns_record_set" "www" {
  name         = "www"
  zone_id      = vinyldns_zone.example.id
  type         = "CNAME"
  ttl          = 300
  record_cname = "example.com."
}

resource "vinyldns_record_set" "mail" {
  name             = "mail"
  zone_id          = vinyldns_zone.example.id
  type             = "A"
  ttl              = 300
  record_addresses = ["192.0.2.10"]
}

# Variables
variable "group_member_ids" {
  description = "List of user IDs to be members of the zone owner group"
  type        = list(string)
}

variable "group_admin_ids" {
  description = "List of user IDs to be admins of the zone owner group"
  type        = list(string)
}

# Outputs
output "zone_id" {
  description = "The ID of the created zone"
  value       = vinyldns_zone.example.id
}

output "zone_status" {
  description = "The status of the created zone"
  value       = vinyldns_zone.example.status
}

output "group_id" {
  description = "The ID of the zone owner group"
  value       = vinyldns_group.zone_owners.id
}
