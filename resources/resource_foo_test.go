package resources_test

import (
	"fmt"
	"github.com/gmllt/terraform-provider-template/client"
	"testing"

	"github.com/gmllt/terraform-provider-template/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceFoo(t *testing.T) {
	t.Parallel()

	client := client.NewMockClient()

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"template": provider.Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFooConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFooExists("template_resource_foo.example", client),
				),
			},
		},
	})
}

func testAccResourceFooConfig_basic() string {
	return `
provider "template" {
  api_key = "test_api_key"
}

resource "template_resource_foo" "example" {
  name        = "test_foo"
  description = "A test Foo resource"
}
`
}

func testAccCheckFooExists(resourceName string, client *client.MockClient) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		_, err := client.GetFoo(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("Error fetching Foo resource: %s", err)
		}

		return nil
	}
}
