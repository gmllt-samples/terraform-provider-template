# `template_resource_foo`

The `template_resource_foo` resource allows you to create and manage a Foo resource.

## Example Usage

```hcl
resource "template_resource_foo" "example" {
name        = "example_foo"
description = "An example Foo resource"
}
```

## Argument Reference

- `name` (Required, String) - The name of the Foo resource.
- `description` (Optional, String) - A description of the Foo resource.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

- `id` - The ID of the Foo resource.
- `name` - The name of the Foo resource.
- `description` - The description of the Foo resource.
