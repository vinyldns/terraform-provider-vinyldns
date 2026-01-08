# Prerequisites: group and zone
resource "vinyldns_group" "example" {
  name       = "example-group"
  email      = "dns@example.com"
  member_ids = ["user-id-1"]
  admin_ids  = ["user-id-1"]
}

resource "vinyldns_zone" "example" {
  name           = "example.com."
  email          = "dns@example.com"
  admin_group_id = vinyldns_group.example.id
}

# A record (single address)
resource "vinyldns_record_set" "web" {
  name             = "www"
  zone_id          = vinyldns_zone.example.id
  type             = "A"
  ttl              = 300
  record_addresses = ["192.0.2.1"]
}

# A record (multiple addresses for load balancing)
resource "vinyldns_record_set" "api" {
  name             = "api"
  zone_id          = vinyldns_zone.example.id
  type             = "A"
  ttl              = 60
  record_addresses = ["192.0.2.10", "192.0.2.11", "192.0.2.12"]
}

# AAAA record (IPv6)
resource "vinyldns_record_set" "web_ipv6" {
  name             = "www"
  zone_id          = vinyldns_zone.example.id
  type             = "AAAA"
  ttl              = 300
  record_addresses = ["2001:db8::1"]
}

# CNAME record
resource "vinyldns_record_set" "alias" {
  name         = "blog"
  zone_id      = vinyldns_zone.example.id
  type         = "CNAME"
  ttl          = 3600
  record_cname = "www.example.com."
}

# TXT record (single value)
resource "vinyldns_record_set" "spf" {
  name         = "@"
  zone_id      = vinyldns_zone.example.id
  type         = "TXT"
  ttl          = 3600
  record_texts = ["v=spf1 include:_spf.example.com ~all"]
}

# TXT record (multiple values, e.g., for domain verification)
resource "vinyldns_record_set" "verification" {
  name         = "_dmarc"
  zone_id      = vinyldns_zone.example.id
  type         = "TXT"
  ttl          = 3600
  record_texts = ["v=DMARC1; p=reject; rua=mailto:dmarc@example.com"]
}

# NS record (delegation)
resource "vinyldns_record_set" "subdomain_ns" {
  name            = "subdomain"
  zone_id         = vinyldns_zone.example.id
  type            = "NS"
  ttl             = 86400
  record_nsdnames = ["ns1.subdomain.example.com.", "ns2.subdomain.example.com."]
}

# PTR record (reverse DNS)
# Note: This would be in a reverse zone like "2.0.192.in-addr.arpa."
resource "vinyldns_record_set" "ptr" {
  name             = "1"
  zone_id          = "reverse-zone-id"
  type             = "PTR"
  ttl              = 3600
  record_ptrdnames = ["www.example.com."]
}

# Record with owner group (for shared zones)
resource "vinyldns_record_set" "owned_record" {
  name             = "app"
  zone_id          = vinyldns_zone.example.id
  type             = "A"
  ttl              = 300
  record_addresses = ["192.0.2.50"]
  owner_group_id   = vinyldns_group.example.id
}
