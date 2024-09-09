# terraform-provider-template

The `terraform-provider-template` is a Terraform provider that allows you to manage Foo and Bar resources via Terraform.

## Features

- **Resources:**
   - `template_resource_foo`: Manages a Foo resource.
   - `template_resource_bar`: Manages a Bar resource.
- **Data Sources:**
   - `template_data_source_foo`: Retrieves information about a Foo resource.
   - `template_data_source_bar`: Retrieves information about a Bar resource.

## Installation

Clone the repository and install the dependencies:

```bash
git clone https://github.com/gmllt-samples/terraform-provider-template.git
cd terraform-provider-template
go mod tidy
```

## Usage

Here's an example of how to use the provider in your Terraform configuration:

```hcl
provider "template" {
api_key = "your_api_key_here"
}

resource "template_resource_foo" "example" {
name        = "example_foo"
description = "An example Foo resource"
}

resource "template_resource_bar" "example" {
title   = "example_bar"
content = "An example Bar resource"
}

data "template_data_source_foo" "example" {
foo_id = template_resource_foo.example.id
}

data "template_data_source_bar" "example" {
bar_id = template_resource_bar.example.id
}
```

## Contributing

See the [CONTRIBUTING.md](./CONTRIBUTING.md) file for details.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
