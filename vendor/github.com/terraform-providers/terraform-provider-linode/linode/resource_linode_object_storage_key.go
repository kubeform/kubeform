package linode

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/linode/linodego"
)

func resourceLinodeObjectStorageKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceLinodeObjectStorageKeyCreate,
		Read:   resourceLinodeObjectStorageKeyRead,
		Update: resourceLinodeObjectStorageKeyUpdate,
		Delete: resourceLinodeObjectStorageKeyDelete,

		Schema: map[string]*schema.Schema{
			"label": {
				Type:        schema.TypeString,
				Description: "The label given to this key. For display purposes only.",
				Required:    true,
			},
			"access_key": {
				Type:        schema.TypeString,
				Description: "This keypair's access key. This is not secret.",
				Computed:    true,
			},
			"secret_key": {
				Type:        schema.TypeString,
				Description: "This keypair's secret key.",
				Sensitive:   true,
				Computed:    true,
			},
		},
	}
}

func resourceLinodeObjectStorageKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client, ok := meta.(linodego.Client)
	if !ok {
		return fmt.Errorf("Invalid Client when creating Linode Object Storage Key")
	}

	createOpts := linodego.ObjectStorageKeyCreateOptions{
		Label: d.Get("label").(string),
	}
	objectStorageKey, err := client.CreateObjectStorageKey(context.Background(), createOpts)
	if err != nil {
		return fmt.Errorf("Error creating a Linode Object Storage Key: %s", err)
	}
	d.SetId(fmt.Sprintf("%d", objectStorageKey.ID))
	d.Set("label", objectStorageKey.Label)
	d.Set("access_key", objectStorageKey.AccessKey)

	// secret_key only available on creation
	d.Set("secret_key", objectStorageKey.SecretKey)

	return resourceLinodeObjectStorageKeyRead(d, meta)
}

func resourceLinodeObjectStorageKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode Object Storage Key ID %s as int: %s", d.Id(), err)
	}

	objectStorageKey, err := client.GetObjectStorageKey(context.Background(), int(id))

	if err != nil {
		return fmt.Errorf("Error finding the specified Linode Object Storage Key: %s", err)
	}

	d.Set("label", objectStorageKey.Label)
	d.Set("access_key", objectStorageKey.AccessKey)
	return nil
}

func resourceLinodeObjectStorageKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode Object Storage Key id %s as int: %s", d.Id(), err)
	}

	if d.HasChange("label") {
		objectStorageKey, err := client.GetObjectStorageKey(context.Background(), int(id))

		updateOpts := linodego.ObjectStorageKeyUpdateOptions{
			Label: d.Get("label").(string),
		}

		if err != nil {
			return fmt.Errorf("Error fetching data about the current Linode Object Storage Key: %s", err)
		}

		if objectStorageKey, err = client.UpdateObjectStorageKey(context.Background(), int(id), updateOpts); err != nil {
			return err
		}
		d.Set("label", objectStorageKey.Label)
	}

	return resourceLinodeObjectStorageKeyRead(d, meta)
}

func resourceLinodeObjectStorageKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode Object Storage Key id %s as int", d.Id())
	}
	err = client.DeleteObjectStorageKey(context.Background(), int(id))
	if err != nil {
		return fmt.Errorf("Error deleting Linode Object Storage Key %d: %s", id, err)
	}
	return nil
}
