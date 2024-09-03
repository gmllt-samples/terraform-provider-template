data "template_data_source_foo" "example" {
  foo_id = template_resource_foo.example.id
}

output "foo_data_name" {
  value = data.template_data_source_foo.example.name
}

output "foo_data_description" {
  value = data.template_data_source_foo.example.description
}
