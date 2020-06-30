package linode

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/linode/linodego"
)

func dataSourceLinodeImage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLinodeImageRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"label": {
				Type:        schema.TypeString,
				Description: "A short description of the Image. Labels cannot contain special characters.",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "A detailed description of this Image.",
				Computed:    true,
			},
			"created": {
				Type:        schema.TypeString,
				Description: "When this Image was created.",
				Computed:    true,
			},
			"created_by": {
				Type:        schema.TypeString,
				Description: "The name of the User who created this Image.",
				Computed:    true,
			},
			"deprecated": {
				Type:        schema.TypeBool,
				Description: "Whether or not this Image is deprecated. Will only be True for deprecated public Images.",
				Computed:    true,
			},
			"is_public": {
				Type:        schema.TypeBool,
				Description: "True if the Image is public.",
				Computed:    true,
			},
			"size": {
				Type:        schema.TypeInt,
				Description: "The minimum size this Image needs to deploy. Size is in MB.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "How the Image was created. 'Manual' Images can be created at any time. 'Automatic' images are created automatically from a deleted Linode.",
				Computed:    true,
			},
			"expiry": {
				Type:        schema.TypeString,
				Description: "Only Images created automatically (from a deleted Linode; type=automatic) will expire.",
				Computed:    true,
			},
			"vendor": {
				Type:        schema.TypeString,
				Description: "The upstream distribution vendor. Nil for private Images.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLinodeImageRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	reqImage := d.Get("id").(string)

	if reqImage == "" {
		return fmt.Errorf("Image id is required")
	}

	image, err := client.GetImage(context.Background(), reqImage)
	if err != nil {
		return fmt.Errorf("Error listing images: %s", err)
	}

	if image != nil {
		d.SetId(image.ID)
		d.Set("label", image.Label)
		d.Set("description", image.Description)
		if image.Created != nil {
			d.Set("created", image.Created.Format(time.RFC3339))
		}
		if image.Expiry != nil {
			d.Set("expiry", image.Expiry.Format(time.RFC3339))
		}
		d.Set("created_by", image.CreatedBy)
		d.Set("deprecated", image.Deprecated)
		d.Set("is_public", image.IsPublic)
		d.Set("size", image.Size)
		d.Set("type", image.Type)
		d.Set("vendor", image.Vendor)
		return nil
	}

	d.SetId("")

	return fmt.Errorf("Image %s was not found", reqImage)
}
