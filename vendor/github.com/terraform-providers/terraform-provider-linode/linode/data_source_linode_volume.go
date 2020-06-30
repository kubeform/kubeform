package linode

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/linode/linodego"
)

func dataSourceLinodeVolume() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLinodeVolumeRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeInt,
				Description: "The unique id of this Volume.",
				Required:    true,
			},
			"label": {
				Type:        schema.TypeString,
				Description: "The Volume's label. For display purposes only.",
				Computed:    true,
			},
			"region": {
				Type:        schema.TypeString,
				Description: "The datacenter where this Volume is located.",
				Computed:    true,
			},
			"size": {
				Type:        schema.TypeInt,
				Description: "The size of this Volume in GiB.",
				Computed:    true,
			},
			"linode_id": {
				Type:        schema.TypeInt,
				Description: "If a Volume is attached to a specific Linode, the ID of that Linode will be displayed here.",
				Computed:    true,
			},
			"filesystem_path": {
				Type:        schema.TypeString,
				Description: "The full filesystem path for the Volume based on the Volume's label. Path is /dev/disk/by-id/scsi-0LinodeVolume + Volume label.",
				Computed:    true,
			},
			"tags": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of tags applied to this Volume. Tags are for organizational purposes only.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "The status of the Volume. Can be one of active | creating | resizing | contact_support",
				Computed:    true,
			},
			"created": {
				Type:        schema.TypeString,
				Description: "Datetime string representing when the Volume was created.",
				Computed:    true,
			},
			"updated": {
				Type:        schema.TypeString,
				Description: "Datetime string representing when the Volume was last updated.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLinodeVolumeRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	requestedVolumeID := d.Get("id").(int)

	if requestedVolumeID == 0 {
		return fmt.Errorf("Volume ID is required")
	}

	var volume *linodego.Volume

	volume, err := client.GetVolume(context.Background(), requestedVolumeID)

	if err != nil {
		return fmt.Errorf("Error requesting Volume: %s", err)
	}

	if volume != nil {
		d.SetId(strconv.Itoa(volume.ID))
		d.Set("region", volume.Region)
		d.Set("size", volume.Size)
		d.Set("filesystem_path", volume.FilesystemPath)
		d.Set("label", volume.Label)
		d.Set("linode_id", volume.LinodeID)
		d.Set("status", volume.Status)
		d.Set("created", volume.Created.Format(time.RFC3339))
		d.Set("updated", volume.Updated.Format(time.RFC3339))
		d.Set("tags", volume.Tags)
		return nil
	}

	return fmt.Errorf("Linode Volume %s was not found", string(requestedVolumeID))

}
