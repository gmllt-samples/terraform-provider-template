resource "template_resource_bar" "example" {
  title   = "example_bar"
  content = "An example Bar resource"
}

output "bar_id" {
  value = template_resource_bar.example.id
}

output "bar_title" {
  value = template_resource_bar.example.title
}

output "bar_content" {
  value = template_resource_bar.example.content
}
