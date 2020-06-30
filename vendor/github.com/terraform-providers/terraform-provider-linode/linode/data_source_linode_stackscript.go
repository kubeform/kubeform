package linode

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/linode/linodego"
)

func dataSourceLinodeStackscript() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLinodeStackscriptRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"label": {
				Type:        schema.TypeString,
				Description: "The StackScript's label is for display purposes only.",
				Computed:    true,
			},
			"script": {
				Type:        schema.TypeString,
				Description: "The script to execute when provisioning a new Linode with this StackScript.",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "A description for the StackScript.",
				Computed:    true,
			},
			"rev_note": {
				Type:        schema.TypeString,
				Description: "This field allows you to add notes for the set of revisions made to this StackScript.",
				Computed:    true,
			},
			"is_public": {
				Type:        schema.TypeBool,
				Description: "This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private.",
				Computed:    true,
			},
			"images": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of Image IDs representing the Images that this StackScript is compatible for deploying with.",
				Computed:    true,
			},

			"deployments_active": {
				Type:        schema.TypeInt,
				Description: "Count of currently active, deployed Linodes created from this StackScript.",
				Computed:    true,
			},
			"user_gravatar_id": {
				Type:        schema.TypeString,
				Description: "The Gravatar ID for the User who created the StackScript.",
				Computed:    true,
			},
			"deployments_total": {
				Type:        schema.TypeInt,
				Description: "The total number of times this StackScript has been deployed.",
				Computed:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "The User who created the StackScript.",
				Computed:    true,
			},
			"created": {
				Type:        schema.TypeString,
				Description: "The date this StackScript was created.",
				Computed:    true,
			},
			"updated": {
				Type:        schema.TypeString,
				Description: "The date this StackScript was updated.",
				Computed:    true,
			},
			"user_defined_fields": {
				Description: "This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment.",
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ConfigMode:  schema.SchemaConfigModeAttr,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"example": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"one_of": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"many_of": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"default": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceLinodeStackscriptRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)
	id := d.Get("id").(int)

	ss, err := client.GetStackscript(context.Background(), id)
	if err != nil {
		return fmt.Errorf("Error getting Staakscript: %s", err)
	}

	if ss != nil {
		d.SetId(strconv.Itoa(id))
		d.Set("label", ss.Label)
		d.Set("script", ss.Script)
		d.Set("description", ss.Description)
		d.Set("rev_note", ss.RevNote)
		d.Set("is_public", ss.IsPublic)
		d.Set("images", ss.Images)
		d.Set("user_gravatar_id", ss.UserGravatarID)
		d.Set("deployments_active", ss.DeploymentsActive)
		d.Set("deployments_total", ss.DeploymentsTotal)
		d.Set("username", ss.Username)
		d.Set("created", ss.Created.Format(time.RFC3339))
		d.Set("updated", ss.Created.Format(time.RFC3339))
		setStackScriptUserDefinedFields(d, ss)
		return nil
	}

	return fmt.Errorf("StackScript %d not found", id)
}
