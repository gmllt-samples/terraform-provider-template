# `template_resource_bar`

The `template_resource_bar` resource allows you to create and manage a Bar resource.

## Example Usage

```hcl
resource "template_resource_bar" "example" {
title   = "example_bar"
content = "An example Bar resource"
}
```

## Argument Reference

- `title` (Required, String) - The title of the Bar resource.
- `content` (Optional, String) - The content of the Bar resource.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

- `id` - The ID of the Bar resource.
- `title` - The title of the Bar resource.
- `content` - The content of the Bar resource.
