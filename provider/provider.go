package provider

import (
	"github.com/gmllt/terraform-provider-template/client"
	datasources2 "github.com/gmllt/terraform-provider-template/datasources"
	resources2 "github.com/gmllt/terraform-provider-template/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns the terraform provider definition
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			// Configuration schema for the provider (e.g., API keys, endpoints)
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_KEY", nil),
				Description: "API key for authenticating with the service",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"template_resource_foo": resources2.ResourceFoo(),
			"template_resource_bar": resources2.ResourceBar(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"template_data_source_foo": datasources2.DataSourceFoo(),
			"template_data_source_bar": datasources2.DataSourceBar(),
		},
		ConfigureFunc: providerConfigure,
	}
}

// providerConfigure sets up the provider client
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &client.Config{
		APIKey: d.Get("api_key").(string),
	}

	return client.NewClient(config), nil
}
