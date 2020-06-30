package linode

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/linode/linodego"
)

func dataSourceLinodeSSHKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLinodeSSHKeyRead,

		Schema: map[string]*schema.Schema{
			"label": {
				Type:        schema.TypeString,
				Description: "The label of the Linode SSH Key.",
				Required:    true,
			},
			"ssh_key": {
				Type:        schema.TypeString,
				Description: "The public SSH Key, which is used to authenticate to the root user of the Linodes you deploy.",
				Computed:    true,
			},
			"created": {
				Type:        schema.TypeString,
				Description: "The date this key was added.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLinodeSSHKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	reqLabel := d.Get("label").(string)

	if reqLabel == "" {
		return fmt.Errorf("Error SSH Key label is required")
	}

	sshkeys, err := client.ListSSHKeys(context.Background(), nil)
	var sshkey linodego.SSHKey
	if err != nil {
		return fmt.Errorf("Error listing sshkey: %s", err)
	}

	for _, testkey := range sshkeys {
		if testkey.Label == reqLabel {
			sshkey = testkey
			break
		}
	}

	if sshkey.ID != 0 {
		d.SetId(fmt.Sprintf("%d", sshkey.ID))
		d.Set("label", sshkey.Label)
		d.Set("ssh_key", sshkey.SSHKey)
		if sshkey.Created != nil {
			d.Set("created", sshkey.Created.Format(time.RFC3339))
		}

		return nil
	}

	return fmt.Errorf("Linode SSH Key with label %s was not found", reqLabel)
}
