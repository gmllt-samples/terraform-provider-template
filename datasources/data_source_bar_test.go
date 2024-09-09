package datasources_test

import (
	"fmt"
	"github.com/gmllt-samples/terraform-provider-template/client"
	"testing"

	"github.com/gmllt-samples/terraform-provider-template/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestDataSourceBar(t *testing.T) {
	t.Parallel()

	client := client.NewMockClient()

	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"template": provider.Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceBarConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBarDataSource("template_data_source_bar.example", client),
				),
			},
		},
	})
}

func testAccDataSourceBarConfig_basic() string {
	return `
provider "template" {
  api_key = "test_api_key"
}

data "template_data_source_bar" "example" {
  bar_id = "test_bar"
}
`
}

func testAccCheckBarDataSource(resourceName string, client *client.MockClient) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		_, err := client.GetBar(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("Error fetching Bar data source: %s", err)
		}

		return nil
	}
}
