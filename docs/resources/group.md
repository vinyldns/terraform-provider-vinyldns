# vinyldns\_group

The group resource allows VinylDNS groups to be created and managed.

## Example Usage

```hcl
# Create a VinylDNS group
resource "vinyldns_group" "test_group" {
  name = "terraform-provider-test-group"
  member_ids = ["123"]
  admin_ids = ["123"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name for the group.

* `email` - (Required) The email address for the group.

* `description` - (Optional) A description of the group.

* `member_ids` - (Required) A list of member IDs to associate with the group.

* `admin_ids` - (Required) A list of admin IDs to associate with the group.
