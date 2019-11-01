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

## Import

`vinyldns_group` can be imported using the ID of the group. For example:

```
terraform import vinyldns_group.example group_id 6f8afcda-7529-4cad-9f2d-76903f4b1aca
```
