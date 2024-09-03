package resources

import (
	"context"
	"fmt"
	"github.com/gmllt/terraform-provider-template/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceBar() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBarCreate,
		ReadContext:   resourceBarRead,
		UpdateContext: resourceBarUpdate,
		DeleteContext: resourceBarDelete,

		Schema: map[string]*schema.Schema{
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The title of the Bar resource",
			},
			"content": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The content of the Bar resource",
			},
		},
	}
}

func resourceBarCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	title := d.Get("title").(string)
	content := d.Get("content").(string)

	// Here we simulate an API call with the client
	bar, err := c.CreateBar(title, content)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(bar.ID)

	tflog.Info(ctx, fmt.Sprintf("Created Bar resource: %s", bar.ID))

	return resourceBarRead(ctx, d, m)
}

func resourceBarRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	barID := d.Id()

	// Simulate retrieving the Bar resource by ID
	bar, err := c.GetBar(barID)
	if err != nil {
		if client.IsNotFound(err) {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.Set("title", bar.Title)
	d.Set("content", bar.Content)

	tflog.Info(ctx, fmt.Sprintf("Read Bar resource: %s", bar.ID))

	return nil
}

func resourceBarUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	barID := d.Id()
	title := d.Get("title").(string)
	content := d.Get("content").(string)

	// Simulate updating the Bar resource
	_, err := c.UpdateBar(barID, title, content)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, fmt.Sprintf("Updated Bar resource: %s", barID))

	return resourceBarRead(ctx, d, m)
}

func resourceBarDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	barID := d.Id()

	// Simulate deleting the Bar resource
	err := c.DeleteBar(barID)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, fmt.Sprintf("Deleted Bar resource: %s", barID))

	d.SetId("")

	return nil
}
