# Basic Zone Setup

This example shows the minimum configuration needed to create a functional DNS zone in VinylDNS.

## What This Creates

- A VinylDNS group to administer the zone
- A DNS zone
- Basic DNS records (A record for root, CNAME for www, A record for mail)

## Prerequisites

- VinylDNS server running and accessible
- User IDs for group members and admins

## Usage

1. Set environment variables for authentication:

```bash
export VINYLDNS_HOST="https://vinyldns.example.com"
export VINYLDNS_ACCESS_KEY="your-access-key"
export VINYLDNS_SECRET_KEY="your-secret-key"
```

2. Create a `terraform.tfvars` file:

```hcl
group_member_ids = ["user-id-1", "user-id-2"]
group_admin_ids  = ["user-id-1"]
```

3. Run Terraform:

```bash
terraform init
terraform plan
terraform apply
```

## Notes

- Zone names must end with a trailing dot (e.g., `example.com.`)
- CNAME records must also end with a trailing dot
- The zone's admin group must exist before the zone is created (Terraform handles this automatically through the dependency)
