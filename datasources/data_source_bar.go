package datasources

import (
	"context"
	"fmt"
	"github.com/gmllt-samples/terraform-provider-template/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceBar() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBarRead,

		Schema: map[string]*schema.Schema{
			"bar_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the Bar resource",
			},
			"title": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The title of the Bar resource",
			},
			"content": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The content of the Bar resource",
			},
		},
	}
}

func dataSourceBarRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	barID := d.Get("bar_id").(string)

	// Simulate retrieving the Bar resource by ID
	bar, err := c.GetBar(barID)
	if err != nil {
		if client.IsNotFound(err) {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.SetId(bar.ID)
	d.Set("title", bar.Title)
	d.Set("content", bar.Content)

	tflog.Info(ctx, fmt.Sprintf("Retrieved Bar resource: %s", bar.ID))

	return nil
}
