# `template_data_source_bar`

The `template_data_source_bar` data source allows you to retrieve information about a Bar resource.

## Example Usage

```hcl
data "template_data_source_bar" "example" {
bar_id = "example_bar_id"
}
```

## Argument Reference

- `bar_id` (Required, String) - The ID of the Bar resource.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

- `id` - The ID of the Bar resource.
- `title` - The title of the Bar resource.
- `content` - The content of the Bar resource.
