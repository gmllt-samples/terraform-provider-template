package datasources

import (
	"context"
	"fmt"
	"github.com/gmllt-samples/terraform-provider-template/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceFoo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFooRead,

		Schema: map[string]*schema.Schema{
			"foo_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the Foo resource",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the Foo resource",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the Foo resource",
			},
		},
	}
}

func dataSourceFooRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	fooID := d.Get("foo_id").(string)

	// Simulate retrieving the Foo resource by ID
	foo, err := c.GetFoo(fooID)
	if err != nil {
		if client.IsNotFound(err) {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.SetId(foo.ID)
	d.Set("name", foo.Name)
	d.Set("description", foo.Description)

	tflog.Info(ctx, fmt.Sprintf("Retrieved Foo resource: %s", foo.ID))

	return nil
}
