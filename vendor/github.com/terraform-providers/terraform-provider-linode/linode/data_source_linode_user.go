package linode

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/linode/linodego"
)

func dataSourceLinodeUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLinodeUserRead,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Description: "This User's username. This is used for logging in, and may also be displayed alongside actions the User performs (for example, in Events or public StackScripts).",
				Required:    true,
			},
			"ssh_keys": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "A list of SSH Key labels added by this User. These are the keys that will be deployed if this User is included in the authorized_users field of a create Linode, rebuild Linode, or create Disk request.",
				Computed:    true,
			},
			"email": {
				Type:        schema.TypeString,
				Description: "The email address for this User, for account management communications, and may be used for other communications as configured.",
				Computed:    true,
			},
			"restricted": {
				Type:        schema.TypeBool,
				Description: "If true, this User must be granted access to perform actions or access entities on this Account.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLinodeUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	reqUsername := d.Get("username").(string)

	if reqUsername == "" {
		return fmt.Errorf("Error User username is required")
	}

	users, err := client.ListUsers(context.Background(), nil)
	var user linodego.User
	if err != nil {
		return fmt.Errorf("Error listing user: %s", err)
	}

	for _, testuser := range users {
		if testuser.Username == reqUsername {
			user = testuser
			break
		}
	}

	if user.Username != "" {
		d.SetId(fmt.Sprintf("%s", user.Username))
		d.Set("username", user.Username)
		d.Set("email", user.Email)
		d.Set("ssh_keys", user.SSHKeys)
		d.Set("restricted", user.Restricted)

		return nil
	}

	return fmt.Errorf("Linode User with username %s was not found", reqUsername)
}
