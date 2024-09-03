# `template_data_source_foo`

The `template_data_source_foo` data source allows you to retrieve information about a Foo resource.

## Example Usage

```hcl
data "template_data_source_foo" "example" {
foo_id = "example_foo_id"
}
```

## Argument Reference

- `foo_id` (Required, String) - The ID of the Foo resource.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

- `id` - The ID of the Foo resource.
- `name` - The name of the Foo resource.
- `description` - The description of the Foo resource.
