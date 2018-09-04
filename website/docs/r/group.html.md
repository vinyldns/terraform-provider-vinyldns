---
page_title: "VinylDNS: vinyldns_group"
sidebar_current: "docs-vinyldns-resource-group"
description: |-
  The vinyldns_group resource allows a VinylDNS group to be created and managed.
---

# vinyldns\_group

The group resource allows VinylDNS groups to be created and managed.

## Example Usage

```hcl
# Create a VinylDNS group
resource "vinyldns_group" "test_group" {
  name = "terraform-provider-test-group"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name for the group.

* `email` - (Required) The email address for the group.

* `description` - (Optional) A description of the group.

* `member` - (Optional) A member to associate with the group.
  See [member](#member) below for details.

* `admin` - (Optional) An admin to associate with the group.
  See [admin](#admin) below for details.

### Member

* `username` - (Optional) The member's username.

* `first_name` - (Optional) The member's first name.

* `last_name` - (Optional) The member's last name.

* `email` - (Optional) The member's email address.

* `id` - (Required) The member's UUID.

### Admin

* `username` - (Optional) The member's username.

* `first_name` - (Optional) The member's first name.

* `last_name` - (Optional) The member's last name.

* `email` - (Optional) The member's email address.

* `id` - (Required) The member's UUID.
