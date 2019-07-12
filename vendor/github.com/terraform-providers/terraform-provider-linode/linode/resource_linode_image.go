package linode

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linode/linodego"
)

const (
	LinodeImageCreateTimeout = 20 * time.Minute
)

func resourceLinodeImage() *schema.Resource {
	return &schema.Resource{
		Create: resourceLinodeImageCreate,
		Read:   resourceLinodeImageRead,
		Update: resourceLinodeImageUpdate,
		Delete: resourceLinodeImageDelete,
		Exists: resourceLinodeImageExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(LinodeImageCreateTimeout),
		},
		Schema: map[string]*schema.Schema{
			"label": {
				Type:        schema.TypeString,
				Description: "A short description of the Image. Labels cannot contain special characters.",
				Required:    true,
			},
			"disk_id": {
				Type:        schema.TypeInt,
				Description: "The ID of the Linode Disk that this Image will be created from.",
				Required:    true,
				ForceNew:    true,
			},
			"linode_id": {
				Type:        schema.TypeInt,
				Description: "The ID of the Linode that this Image will be created from.",
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "A detailed description of this Image.",
				Optional:    true,
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

func resourceLinodeImageExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(linodego.Client)

	_, err := client.GetImage(context.Background(), d.Id())
	if err != nil {
		if lerr, ok := err.(*linodego.Error); ok && lerr.Code == 404 {
			return false, nil
		}

		return false, fmt.Errorf("Error getting Linode Image ID %s: %s", d.Id(), err)
	}
	return true, nil
}

func resourceLinodeImageRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	image, err := client.GetImage(context.Background(), d.Id())

	found, err := resourceLinodeImageExists(d, meta)
	if err != nil {
		return fmt.Errorf("Error finding the specified Linode Image: %s", err)
	}
	if !found {
		log.Printf("[WARN] removing Image ID %q from state because it no longer exists", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("label", image.Label)
	d.Set("description", image.Description)
	d.Set("type", image.Type)
	d.Set("size", image.Size)
	d.Set("vendor", image.Vendor)
	d.Set("created_by", image.CreatedBy)
	d.Set("deprecated", image.Deprecated)
	d.Set("is_public", image.IsPublic)
	if image.Created != nil {
		d.Set("created", image.Created.Format(time.RFC3339))
	}
	if image.Expiry != nil {
		d.Set("expiry", image.Expiry.Format(time.RFC3339))
	}

	return nil
}

func resourceLinodeImageCreate(d *schema.ResourceData, meta interface{}) error {
	client, ok := meta.(linodego.Client)
	if !ok {
		return fmt.Errorf("Invalid Client when creating Linode Image")
	}
	d.Partial(true)

	linodeID := d.Get("linode_id").(int)
	diskID := d.Get("disk_id").(int)

	if _, err := client.WaitForInstanceDiskStatus(context.Background(), linodeID, diskID, linodego.DiskReady, int(d.Timeout(schema.TimeoutCreate).Seconds())); err != nil {
		return fmt.Errorf("Error waiting for Linode Instance %d Disk %d to become ready for taking an Image", linodeID, diskID)
	}

	createOpts := linodego.ImageCreateOptions{
		DiskID:      diskID,
		Label:       d.Get("label").(string),
		Description: d.Get("description").(string),
	}

	image, err := client.CreateImage(context.Background(), createOpts)
	if err != nil {
		return fmt.Errorf("Error creating a Linode Image: %s", err)
	}

	d.SetId(image.ID)
	d.SetPartial("label")
	d.SetPartial("description")
	d.Partial(false)

	if _, err := client.WaitForInstanceDiskStatus(context.Background(), linodeID, diskID, linodego.DiskReady, int(d.Timeout(schema.TimeoutCreate).Seconds())); err != nil {
		return fmt.Errorf("Error waiting for Linode Instance %d Disk %d to become ready while taking an Image", linodeID, diskID)
	}

	found, err := resourceLinodeImageExists(d, meta)
	if err != nil {
		return fmt.Errorf("Error finding the newly created Linode Image: %s", err)
	}
	if !found {
		return fmt.Errorf("The Linode Image failed to be created")
	}

	return resourceLinodeImageRead(d, meta)
}

func resourceLinodeImageUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	image, err := client.GetImage(context.Background(), d.Id())
	if err != nil {
		return fmt.Errorf("Error fetching data about the current Image: %s", err)
	}

	updateOpts := linodego.ImageUpdateOptions{}

	if d.HasChange("label") {
		updateOpts.Label = d.Get("label").(string)
	}

	if d.HasChange("description") {
		descString := d.Get("description").(string)
		updateOpts.Description = &descString
	}

	image, err = client.UpdateImage(context.Background(), d.Id(), updateOpts)
	if err != nil {
		return err
	}

	d.Set("label", image.Label)
	d.Set("description", image.Description)

	return resourceLinodeImageRead(d, meta)
}

func resourceLinodeImageDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	err := client.DeleteImage(context.Background(), d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting Linode Image %s: %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}
