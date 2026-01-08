# VinylDNS Provider

The VinylDNS provider allows Terraform to manage [VinylDNS](https://www.vinyldns.io/) resources. VinylDNS is a vendor-agnostic DNS front-end for streamlining DNS operations and enabling self-service for DNS infrastructure.

## Example Usage

```hcl
terraform {
  required_providers {
    vinyldns = {
      source = "vinyldns/vinyldns"
    }
  }
}

# Configure the provider
provider "vinyldns" {
  host       = "https://vinyldns.example.com"
  access_key = var.vinyldns_access_key
  secret_key = var.vinyldns_secret_key
}

# Create a group to administer the zone
resource "vinyldns_group" "example" {
  name       = "example-group"
  email      = "dns-team@example.com"
  member_ids = ["user-id-1"]
  admin_ids  = ["user-id-1"]
}

# Create a zone
resource "vinyldns_zone" "example" {
  name           = "example.com."
  email          = "hostmaster@example.com"
  admin_group_id = vinyldns_group.example.id
}

# Create a DNS record
resource "vinyldns_record_set" "www" {
  name             = "www"
  zone_id          = vinyldns_zone.example.id
  type             = "A"
  ttl              = 300
  record_addresses = ["192.0.2.1"]
}
```

## Authentication

The provider supports authentication via explicit configuration or environment variables.

### Environment Variables (Recommended)

```bash
export VINYLDNS_HOST="https://vinyldns.example.com"
export VINYLDNS_ACCESS_KEY="your-access-key"
export VINYLDNS_SECRET_KEY="your-secret-key"
```

Then configure the provider without credentials:

```hcl
provider "vinyldns" {}
```

### Explicit Configuration

```hcl
provider "vinyldns" {
  host       = "https://vinyldns.example.com"
  access_key = var.vinyldns_access_key
  secret_key = var.vinyldns_secret_key
}
```

## Argument Reference

* `host` - (Required) The VinylDNS API endpoint URL. May also be set via the `VINYLDNS_HOST` environment variable.

* `access_key` - (Required) The access key for VinylDNS API authentication. May also be set via the `VINYLDNS_ACCESS_KEY` environment variable.

* `secret_key` - (Required) The secret key for VinylDNS API authentication. May also be set via the `VINYLDNS_SECRET_KEY` environment variable.
