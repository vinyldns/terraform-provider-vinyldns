# vinyldns_group

Use this data source to retrieve the `name`, `email`, and `description` for a group.

## Example Usage

```hcl
data "vinyldns_group" "test" {
  id = "foo"
}
```

## Arguments Reference

* `id` - (Required) The id of the group.

## Attributes Reference

* `name` - The name of the group

* `desription` - The description of the group

* `email` - The email associated with the group
