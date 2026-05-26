# Complete Zone Setup

This example demonstrates a production-ready DNS zone configuration with multiple groups, ACL rules, and TSIG authentication.

## What This Creates

### Groups
- **Zone Admins** - Full administrative access to the zone
- **Zone Readers** - Read-only access for monitoring
- **Web Team** - Write access to www.* and web.* records only

### Zone Configuration
- Zone with TSIG authentication for secure updates
- Optional transfer connection for zone transfers from a separate server
- ACL rules providing granular access control

### DNS Records
- Root domain A records
- WWW CNAME pointing to root
- Mail server A record
- SPF and DMARC TXT records for email security
- API endpoint with owner group assignment

## Prerequisites

- VinylDNS server running and accessible
- TSIG keys configured on your DNS server
- User IDs for the various groups

## Usage

1. Set environment variables:

```bash
export VINYLDNS_HOST="https://vinyldns.example.com"
export VINYLDNS_ACCESS_KEY="your-access-key"
export VINYLDNS_SECRET_KEY="your-secret-key"
```

2. Create a `terraform.tfvars` file:

```hcl
zone_name          = "example.com"
admin_email        = "dns-admin@example.com"
admin_user_ids     = ["admin-user-1", "admin-user-2"]
reader_user_ids    = ["monitoring-user"]
web_team_user_ids  = ["web-dev-1", "web-dev-2"]

primary_dns_server = "ns1.example.com"
tsig_key_name      = "vinyldns-key."
tsig_key           = "base64-encoded-key-here"

root_ips        = ["192.0.2.1", "192.0.2.2"]
mail_server_ips = ["192.0.2.10"]
api_server_ips  = ["192.0.2.20", "192.0.2.21"]
```

3. Run Terraform:

```bash
terraform init
terraform plan
terraform apply
```

## ACL Rules Explained

| Group | Access Level | Record Pattern | Record Types | Purpose |
|-------|--------------|----------------|--------------|---------|
| Zone Readers | Read | All | All | Monitoring and auditing |
| Web Team | Write | www.* | A, AAAA, CNAME | Manage www subdomain records |
| Web Team | Write | web.* | A, AAAA, CNAME | Manage web.* subdomain records |

## Notes

- TSIG keys should be stored securely (use environment variables or a secrets manager)
- The `owner_group_id` on records determines which group "owns" a record in shared zones
- ACL `record_mask` supports regex patterns for matching record names
