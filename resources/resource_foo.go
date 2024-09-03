package resources

import (
	"context"
	"fmt"
	"github.com/gmllt/terraform-provider-template/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceFoo() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFooCreate,
		ReadContext:   resourceFooRead,
		UpdateContext: resourceFooUpdate,
		DeleteContext: resourceFooDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the Foo resource",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the Foo resource",
			},
		},
	}
}

func resourceFooCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	// Here we simulate an API call with the client
	foo, err := c.CreateFoo(name, description)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(foo.ID)

	tflog.Info(ctx, fmt.Sprintf("Created Foo resource: %s", foo.ID))

	return resourceFooRead(ctx, d, m)
}

func resourceFooRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	fooID := d.Id()

	// Simulate retrieving the Foo resource by ID
	foo, err := c.GetFoo(fooID)
	if err != nil {
		if client.IsNotFound(err) {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.Set("name", foo.Name)
	d.Set("description", foo.Description)

	tflog.Info(ctx, fmt.Sprintf("Read Foo resource: %s", foo.ID))

	return nil
}

func resourceFooUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	fooID := d.Id()
	name := d.Get("name").(string)
	description := d.Get("description").(string)

	// Simulate updating the Foo resource
	_, err := c.UpdateFoo(fooID, name, description)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, fmt.Sprintf("Updated Foo resource: %s", fooID))

	return resourceFooRead(ctx, d, m)
}

func resourceFooDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	fooID := d.Id()

	// Simulate deleting the Foo resource
	err := c.DeleteFoo(fooID)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, fmt.Sprintf("Deleted Foo resource: %s", fooID))

	d.SetId("")

	return nil
}
