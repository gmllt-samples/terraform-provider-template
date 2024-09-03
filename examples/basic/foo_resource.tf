resource "template_resource_foo" "example" {
  name        = "example_foo"
  description = "An example Foo resource"
}

output "foo_id" {
  value = template_resource_foo.example.id
}

output "foo_name" {
  value = template_resource_foo.example.name
}

output "foo_description" {
  value = template_resource_foo.example.description
}
