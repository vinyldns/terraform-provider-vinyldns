# vinyldns_zone

Manages a VinylDNS zone. A zone represents a DNS zone that VinylDNS will manage.

## Example Usage

### Basic Zone

```hcl
resource "vinyldns_group" "example" {
  name       = "zone-admin-group"
  email      = "dns-team@example.com"
  member_ids = ["user-id-1"]
  admin_ids  = ["user-id-1"]
}

resource "vinyldns_zone" "example" {
  name           = "example.com."
  email          = "hostmaster@example.com"
  admin_group_id = vinyldns_group.example.id
}
```

### Zone with TSIG Connection

```hcl
resource "vinyldns_zone" "with_connection" {
  name           = "secure.example.com."
  email          = "hostmaster@example.com"
  admin_group_id = vinyldns_group.example.id

  zone_connection {
    name           = "secure.example.com."
    key_name       = "tsig-key."
    key            = "base64-encoded-tsig-key"
    primary_server = "ns1.example.com"
  }
}
```

### Zone with Transfer Connection

Use a transfer connection when zone data should be synced from a different server than where updates are sent:

```hcl
resource "vinyldns_zone" "with_transfer" {
  name           = "transferred.example.com."
  email          = "hostmaster@example.com"
  admin_group_id = vinyldns_group.example.id

  zone_connection {
    name           = "transferred.example.com."
    key_name       = "update-key."
    key            = "base64-encoded-update-key"
    primary_server = "ns1.example.com"
  }

  transfer_connection {
    name           = "transferred.example.com."
    key_name       = "transfer-key."
    key            = "base64-encoded-transfer-key"
    primary_server = "ns2.example.com"
  }
}
```

### Zone with ACL Rules

```hcl
resource "vinyldns_zone" "with_acl" {
  name           = "restricted.example.com."
  email          = "hostmaster@example.com"
  admin_group_id = vinyldns_group.example.id

  # Grant read access to a group
  acl_rule {
    access_level = "Read"
    group_id     = "reader-group-id"
    description  = "Read access for monitoring team"
  }

  # Grant write access for specific record types
  acl_rule {
    access_level = "Write"
    group_id     = "web-team-group-id"
    record_types = ["A", "AAAA", "CNAME"]
    description  = "Web team can manage A/AAAA/CNAME records"
  }

  # Grant write access for records matching a pattern
  acl_rule {
    access_level = "Write"
    group_id     = "api-team-group-id"
    record_mask  = "api-.*"
    record_types = ["A", "CNAME"]
    description  = "API team can manage api-* records"
  }

  # Grant access to a specific user
  acl_rule {
    access_level = "Delete"
    user_id      = "specific-user-id"
    description  = "Full access for specific user"
  }
}
```

## Argument Reference

* `name` - (Required) The name of the zone. Must end with a trailing dot (e.g., `example.com.`).

* `email` - (Required) The email address associated with the zone (typically hostmaster or admin contact).

* `admin_group_id` - (Required) The ID of the group that will administer this zone.

* `zone_connection` - (Optional) Connection details for issuing DDNS updates to the backend zone. See [Zone Connection](#zone-connection) below.

* `transfer_connection` - (Optional) Connection details for syncing zone data from a DNS backend. See [Transfer Connection](#transfer-connection) below.

* `acl_rule` - (Optional) Access control rules for the zone. Multiple rules can be specified. See [ACL Rule](#acl-rule) below.

### Zone Connection

The `zone_connection` block supports:

* `name` - (Required) The connection name (typically matches the zone name).

* `key_name` - (Required) The name of the TSIG key configured on the DNS server.

* `key` - (Required) The TSIG secret key (base64 encoded).

* `primary_server` - (Required) The IP address or hostname of the DNS server.

### Transfer Connection

The `transfer_connection` block supports the same arguments as `zone_connection`. Use this when zone transfers should come from a different server than where updates are sent.

### ACL Rule

The `acl_rule` block supports:

* `access_level` - (Required) The access level to grant. Valid values: `Read`, `Write`, `Delete`.

* `group_id` - (Optional) The ID of a group to grant access to. Either `group_id` or `user_id` should be specified.

* `user_id` - (Optional) The ID of a user to grant access to. Either `group_id` or `user_id` should be specified.

* `record_types` - (Optional) A set of record types this rule applies to (e.g., `["A", "AAAA", "CNAME"]`). If not specified, the rule applies to all record types.

* `record_mask` - (Optional) A regex pattern to match record names. If not specified, the rule applies to all records.

* `description` - (Optional) A description of the ACL rule. Defaults to "Managed by Terraform".

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The unique identifier of the zone.

* `status` - The zone status (e.g., `Active`, `Syncing`).

* `shared` - Whether the zone is a shared zone.

* `created` - The timestamp when the zone was created.

* `updated` - The timestamp when the zone was last updated.

* `latest_sync` - The timestamp of the last zone sync.

## Import

Zones can be imported using their ID:

```shell
terraform import vinyldns_zone.example 9cbdd3ac-9752-4d56-9ca0-6a1a14fc5562
```
