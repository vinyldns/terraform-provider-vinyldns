# vinyldns\_zone

The zone resource allows VinylDNS zones to be created and managed.

## Example Usage

```hcl
# Create a VinylDNS group
resource "vinyldns_group" "test_group" {
  name = "terraform-provider-test-group"
}

# Create a VinylDNS zone with a zone connection
resource "vinyldns_zone" "test_zone" {
  name           = "system-test."
  email          = "foo@bar.com"
  admin_group_id = "${vinyldns_group.test_group.id}"
  zone_connection {
    name           = "vinyldns."
    key_name       = "vinyldns."
    key            = "123"
    primary_server = "127.0.0.1"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name for the zone created.

* `email` - (Required) The email address to associate with the zone

* `admin_group_id` - (Required) The group ID of the group to make the zone's admin group

* `zone_connection` - (Optional) The connection used to issue DDNS updates to the backend zone.
  See [zone connection](#zone-connection) below for details.

* `transfer_connection` - (Optional) The connection that is used to sync the zone with the DNS backend.
  See [transfer connection](#transfer-connection) below for details.

### Zone Connection

* `name` - (Required) The connection name.

* `key` - (Required) The TSIG secret key used to sign requests when communicating with the primary server.

* `key_name` - (Required) The name of the DNS key that has access to the DNS server and zone.

* `primary_server` - (Required) The IP address or host that is connected to.

### Transfer Connection

* `name` - (Required) The connection name.

* `key` - (Required) The TSIG secret key used to sign requests when communicating with the primary server.

* `key_name` - (Required) The name of the DNS key that has access to the DNS server and zone.

* `primary_server` - (Required) The IP address or host that is connected to.

## Attributes Reference

The following attributes are exported:

* `status` - The zone status.

* `created` - The time when the zone was first created.

* `updated` - The time when the zone was last updated.

* `latest_sync` - The time when the zone was synced.

* `shared` - A boolean flag indicating if the zone is a shared zone.

## Import

`vinyldns_zone` can be imported using the ID of the zone. For example:

```
terraform import vinyldns_zone.example 9cbdd3ac-9752-4d56-9ca0-6a1a14fc5562
```
