package linode

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/linode/linodego"
)

func resourceLinodeSSHKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceLinodeSSHKeyCreate,
		Read:   resourceLinodeSSHKeyRead,
		Update: resourceLinodeSSHKeyUpdate,
		Delete: resourceLinodeSSHKeyDelete,
		Exists: resourceLinodeSSHKeyExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"label": {
				Type:        schema.TypeString,
				Description: "The label of the Linode SSH Key.",
				Required:    true,
			},
			"ssh_key": {
				Type:        schema.TypeString,
				Description: "The public SSH Key, which is used to authenticate to the root user of the Linodes you deploy.",
				Required:    true,
				ForceNew:    true,
			},
			"created": {
				Type:        schema.TypeString,
				Description: "The date this key was added.",
				Computed:    true,
			},
		},
	}
}

func resourceLinodeSSHKeyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return false, fmt.Errorf("Error parsing Linode SSH Key ID %s as int: %s", d.Id(), err)
	}

	_, err = client.GetSSHKey(context.Background(), int(id))
	if err != nil {
		if lerr, ok := err.(*linodego.Error); ok && lerr.Code == 404 {
			d.SetId("")
			return false, nil
		}

		return false, fmt.Errorf("Error getting Linode SSH Key ID %s: %s", d.Id(), err)
	}
	return true, nil
}

func resourceLinodeSSHKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode SSH Key ID %s as int: %s", d.Id(), err)
	}

	sshkey, err := client.GetSSHKey(context.Background(), int(id))

	if err != nil {
		return fmt.Errorf("Error finding the specified Linode SSH Key: %s", err)
	}

	d.Set("label", sshkey.Label)
	d.Set("ssh_key", sshkey.SSHKey)
	if sshkey.Created != nil {
		d.Set("created", sshkey.Created.Format(time.RFC3339))
	}

	return nil
}

func resourceLinodeSSHKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client, ok := meta.(linodego.Client)
	if !ok {
		return fmt.Errorf("Invalid Client when creating Linode SSH Key")
	}

	createOpts := linodego.SSHKeyCreateOptions{
		Label:  d.Get("label").(string),
		SSHKey: d.Get("ssh_key").(string),
	}
	sshkey, err := client.CreateSSHKey(context.Background(), createOpts)
	if err != nil {
		return fmt.Errorf("Error creating a Linode SSH Key: %s", err)
	}
	d.SetId(fmt.Sprintf("%d", sshkey.ID))
	d.Set("label", sshkey.Label)
	d.Set("ssh_key", sshkey.SSHKey)
	if sshkey.Created != nil {
		d.Set("created", sshkey.Created.Format(time.RFC3339))
	}

	return resourceLinodeSSHKeyRead(d, meta)
}

func resourceLinodeSSHKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode SSH Key id %s as int: %s", d.Id(), err)
	}

	if d.HasChange("label") {
		sshkey, err := client.GetSSHKey(context.Background(), int(id))

		updateOpts := sshkey.GetUpdateOptions()
		updateOpts.Label = d.Get("label").(string)

		if err != nil {
			return fmt.Errorf("Error fetching data about the current Linode SSH Key: %s", err)
		}

		if sshkey, err = client.UpdateSSHKey(context.Background(), int(id), updateOpts); err != nil {
			return err
		}
		d.Set("label", sshkey.Label)
	}

	return resourceLinodeSSHKeyRead(d, meta)
}

func resourceLinodeSSHKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return fmt.Errorf("Error parsing Linode SSH Key id %s as int", d.Id())
	}
	err = client.DeleteSSHKey(context.Background(), int(id))
	if err != nil {
		return fmt.Errorf("Error deleting Linode SSH Key %d: %s", id, err)
	}
	return nil
}
