provider "template" {
  api_key = "test_api_key"
}

resource "template_resource_foo" "example" {
  name        = "example_foo"
  description = "An example Foo resource"
}

resource "template_resource_bar" "example" {
  title   = "example_bar"
  content = "An example Bar resource"
}

output "foo_id" {
  value = template_resource_foo.example.id
}

output "bar_id" {
  value = template_resource_bar.example.id
}
