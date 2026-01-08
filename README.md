[![Terraform Registry](https://img.shields.io/github/v/release/vinyldns/terraform-provider-vinyldns?color=834FB9&label=registry&logo=terraform)](https://registry.terraform.io/providers/vinyldns/vinyldns/latest)
[![Build Status](https://github.com/vinyldns/terraform-provider-vinyldns/actions/workflows/release.yml/badge.svg)](https://github.com/vinyldns/terraform-provider-vinyldns/actions/workflows/release.yml)
[![GitHub](https://img.shields.io/github/license/vinyldns/terraform-provider-vinyldns)](https://github.com/vinyldns/vinyldns/blob/master/LICENSE)

# terraform-provider-vinyldns

A [Terraform](https://terraform.io) provider for the [VinylDNS](https://github.com/vinyldns/vinyldns) DNS as a service API.

## Documentation

Full documentation is available at https://vinyldns.github.io/terraform-provider-vinyldns

## Quick Start

### Installation

Add the VinylDNS provider to your Terraform configuration:

```hcl
terraform {
  required_providers {
    vinyldns = {
      source  = "vinyldns/vinyldns"
    }
  }
}

provider "vinyldns" {
  # Configure via environment variables (recommended):
  #   VINYLDNS_HOST
  #   VINYLDNS_ACCESS_KEY
  #   VINYLDNS_SECRET_KEY
}
```

### Basic Usage

```hcl
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

## Resources

- `vinyldns_group` - Manage VinylDNS groups
- `vinyldns_zone` - Manage DNS zones
- `vinyldns_record_set` - Manage DNS records

## Data Sources

- `vinyldns_zone` - Look up a zone by name
- `vinyldns_zones` - List zones with optional filtering
- `vinyldns_group` - Look up a group by name
- `vinyldns_groups` - List groups with optional filtering
- `vinyldns_record_sets` - List record sets in a zone
- `vinyldns_backend_ids` - List available DNS backend IDs

## Examples

See the [examples](./examples) directory for complete working examples:

- [Provider configuration](./examples/provider)
- [Resource examples](./examples/resources)
- [Data source examples](./examples/data-sources)
- [Scenario: Basic zone setup](./examples/scenarios/basic-zone-setup)
- [Scenario: Complete zone setup](./examples/scenarios/complete-zone-setup)

## Development

### Installing from Source

```shell
git clone https://github.com/vinyldns/terraform-provider-vinyldns.git
cd terraform-provider-vinyldns
make install
```

Use the local provider in your configuration:

```hcl
terraform {
  required_providers {
    vinyldns = {
      source  = "local/vinyldns-provider/vinyldns"
      version = "0.0.1"
    }
  }
}
```

### Running Acceptance Tests

Tests require a VinylDNS API running on `localhost:9000`. This is handled automatically via Docker:

```shell
make test
```

### Building

```shell
make build
```

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) for contribution guidelines.

## License

Apache 2.0 - See [LICENSE](./LICENSE) for details.

## Credits

`terraform-provider-vinyldns` builds upon many open source packages. See the full list in the source code.
