# Complete Zone Setup
#
# This example demonstrates a production-ready zone configuration with:
# - Multiple groups with different access levels
# - Zone with TSIG authentication
# - ACL rules for fine-grained access control
# - Various record types

terraform {
  required_providers {
    vinyldns = {
      source = "vinyldns/vinyldns"
    }
  }
}

provider "vinyldns" {}

# =============================================================================
# Groups
# =============================================================================

# Primary zone administrators
resource "vinyldns_group" "zone_admins" {
  name        = "${var.zone_name}-admins"
  email       = var.admin_email
  description = "Administrators for ${var.zone_name}"
  member_ids  = var.admin_user_ids
  admin_ids   = var.admin_user_ids
}

# Read-only access group (e.g., for monitoring)
resource "vinyldns_group" "zone_readers" {
  name        = "${var.zone_name}-readers"
  email       = var.admin_email
  description = "Read-only access to ${var.zone_name}"
  member_ids  = var.reader_user_ids
  admin_ids   = var.admin_user_ids
}

# Web team - can manage web-related records
resource "vinyldns_group" "web_team" {
  name        = "${var.zone_name}-web-team"
  email       = "web-team@example.com"
  description = "Web team - manages www and web-related records"
  member_ids  = var.web_team_user_ids
  admin_ids   = var.admin_user_ids
}

# =============================================================================
# Zone
# =============================================================================

resource "vinyldns_zone" "main" {
  name           = "${var.zone_name}."
  email          = var.admin_email
  admin_group_id = vinyldns_group.zone_admins.id

  # TSIG authentication for zone updates
  zone_connection {
    name           = "${var.zone_name}."
    key_name       = var.tsig_key_name
    key            = var.tsig_key
    primary_server = var.primary_dns_server
  }

  # Separate credentials for zone transfers (if needed)
  dynamic "transfer_connection" {
    for_each = var.transfer_server != "" ? [1] : []
    content {
      name           = "${var.zone_name}."
      key_name       = var.transfer_key_name
      key            = var.transfer_key
      primary_server = var.transfer_server
    }
  }

  # ACL Rules

  # Read access for monitoring team
  acl_rule {
    access_level = "Read"
    group_id     = vinyldns_group.zone_readers.id
    description  = "Read access for monitoring"
  }

  # Web team can manage www.* and web.* A/AAAA/CNAME records
  acl_rule {
    access_level = "Write"
    group_id     = vinyldns_group.web_team.id
    record_mask  = "www.*"
    record_types = ["A", "AAAA", "CNAME"]
    description  = "Web team - www records"
  }

  acl_rule {
    access_level = "Write"
    group_id     = vinyldns_group.web_team.id
    record_mask  = "web.*"
    record_types = ["A", "AAAA", "CNAME"]
    description  = "Web team - web records"
  }
}

# =============================================================================
# DNS Records
# =============================================================================

# Root domain A records
resource "vinyldns_record_set" "root" {
  name             = "${var.zone_name}."
  zone_id          = vinyldns_zone.main.id
  type             = "A"
  ttl              = var.default_ttl
  record_addresses = var.root_ips
}

# WWW CNAME
resource "vinyldns_record_set" "www" {
  name         = "www"
  zone_id      = vinyldns_zone.main.id
  type         = "CNAME"
  ttl          = var.default_ttl
  record_cname = "${var.zone_name}."
}

# Mail server
resource "vinyldns_record_set" "mail" {
  name             = "mail"
  zone_id          = vinyldns_zone.main.id
  type             = "A"
  ttl              = var.default_ttl
  record_addresses = var.mail_server_ips
}

# SPF record
resource "vinyldns_record_set" "spf" {
  name         = "${var.zone_name}."
  zone_id      = vinyldns_zone.main.id
  type         = "TXT"
  ttl          = var.default_ttl
  record_texts = ["v=spf1 mx include:_spf.${var.zone_name} ~all"]
}

# DMARC record
resource "vinyldns_record_set" "dmarc" {
  name         = "_dmarc"
  zone_id      = vinyldns_zone.main.id
  type         = "TXT"
  ttl          = var.default_ttl
  record_texts = ["v=DMARC1; p=quarantine; rua=mailto:dmarc-reports@${var.zone_name}"]
}

# API endpoints with owner group
resource "vinyldns_record_set" "api" {
  name             = "api"
  zone_id          = vinyldns_zone.main.id
  type             = "A"
  ttl              = 60 # Lower TTL for API endpoints
  record_addresses = var.api_server_ips
  owner_group_id   = vinyldns_group.zone_admins.id
}

# =============================================================================
# Variables
# =============================================================================

variable "zone_name" {
  description = "The zone name (without trailing dot)"
  type        = string
}

variable "admin_email" {
  description = "Admin contact email for the zone"
  type        = string
}

variable "admin_user_ids" {
  description = "User IDs for zone administrators"
  type        = list(string)
}

variable "reader_user_ids" {
  description = "User IDs for read-only access"
  type        = list(string)
  default     = []
}

variable "web_team_user_ids" {
  description = "User IDs for the web team"
  type        = list(string)
  default     = []
}

variable "primary_dns_server" {
  description = "Primary DNS server for zone updates"
  type        = string
}

variable "tsig_key_name" {
  description = "TSIG key name for zone authentication"
  type        = string
}

variable "tsig_key" {
  description = "TSIG key (base64 encoded)"
  type        = string
  sensitive   = true
}

variable "transfer_server" {
  description = "DNS server for zone transfers (optional)"
  type        = string
  default     = ""
}

variable "transfer_key_name" {
  description = "TSIG key name for transfers"
  type        = string
  default     = ""
}

variable "transfer_key" {
  description = "TSIG key for transfers (base64 encoded)"
  type        = string
  sensitive   = true
  default     = ""
}

variable "default_ttl" {
  description = "Default TTL for records"
  type        = number
  default     = 300
}

variable "root_ips" {
  description = "IP addresses for the root domain"
  type        = list(string)
}

variable "mail_server_ips" {
  description = "IP addresses for mail servers"
  type        = list(string)
}

variable "api_server_ips" {
  description = "IP addresses for API servers"
  type        = list(string)
}

# =============================================================================
# Outputs
# =============================================================================

output "zone_id" {
  description = "The zone ID"
  value       = vinyldns_zone.main.id
}

output "zone_status" {
  description = "The zone status"
  value       = vinyldns_zone.main.status
}

output "admin_group_id" {
  description = "The admin group ID"
  value       = vinyldns_group.zone_admins.id
}

output "reader_group_id" {
  description = "The reader group ID"
  value       = vinyldns_group.zone_readers.id
}

output "web_team_group_id" {
  description = "The web team group ID"
  value       = vinyldns_group.web_team.id
}
