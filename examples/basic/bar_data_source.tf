data "template_data_source_bar" "example" {
  bar_id = template_resource_bar.example.id
}

output "bar_data_title" {
  value = data.template_data_source_bar.example.title
}

output "bar_data_content" {
  value = data.template_data_source_bar.example.content
}
