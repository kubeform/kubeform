package linode

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/linode/linodego"
)

func dataSourceLinodeInstanceType() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLinodeInstanceTypeRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"label": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Linode Type's label is for display purposes only.",
				Computed:    true,
			},
			"disk": {
				Type:        schema.TypeInt,
				Description: "The Disk size, in MB, of the Linode Type.",
				Computed:    true,
			},
			"class": {
				Type:        schema.TypeString,
				Description: "The class of the Linode Type. There are currently three classes of Linodes: nanode, standard, highmem, dedicated",
				Computed:    true,
			},
			"price": {
				Type:        schema.TypeList,
				Description: "Cost in US dollars, broken down into hourly and monthly charges.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hourly": {
							Type:        schema.TypeFloat,
							Description: "Cost (in US dollars) per hour.",
							Computed:    true,
						},
						"monthly": {
							Type:        schema.TypeFloat,
							Description: "Cost (in US dollars) per month.",
							Computed:    true,
						},
					},
				},
			},
			"addons": {
				Type:        schema.TypeList,
				Description: "",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backups": {
							Type:        schema.TypeList,
							Description: "Information about the optional Backup service offered for Linodes.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"price": {
										Type:        schema.TypeList,
										Description: "Cost of enabling Backups for this Linode Type.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hourly": {
													Type:        schema.TypeFloat,
													Description: "The cost (in US dollars) per hour to add Backups service.",
													Computed:    true,
												},
												"monthly": {
													Type:        schema.TypeFloat,
													Description: "The cost (in US dollars) per month to add Backups service.",
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"network_out": {
				Type:        schema.TypeInt,
				Description: "The Mbits outbound bandwidth allocation.",
				Computed:    true,
			},
			"memory": {
				Type:        schema.TypeInt,
				Description: "Amount of RAM included in this Linode Type.",
				Computed:    true,
			},
			"transfer": {
				Type:        schema.TypeInt,
				Description: "The monthly outbound transfer amount, in MB.",
				Computed:    true,
			},
			"vcpus": {
				Type:        schema.TypeInt,
				Description: "The number of VCPU cores this Linode Type offers.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLinodeInstanceTypeRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	types, err := client.ListTypes(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("Error listing ranges: %s", err)
	}

	reqType := d.Get("id").(string)

	for _, r := range types {
		if r.ID == reqType {
			d.SetId(r.ID)
			d.Set("label", r.Label)
			d.Set("disk", r.Disk)
			d.Set("memory", r.Memory)
			d.Set("vcpus", r.VCPUs)
			d.Set("network_out", r.NetworkOut)
			d.Set("transfer", r.Transfer)
			d.Set("class", r.Class)

			d.Set("price", []map[string]interface{}{{
				"hourly":  r.Price.Hourly,
				"monthly": r.Price.Monthly,
			}})

			d.Set("addons", []map[string]interface{}{{
				"backups": []map[string]interface{}{{
					"price": []map[string]interface{}{{
						"hourly":  r.Addons.Backups.Price.Hourly,
						"monthly": r.Addons.Backups.Price.Monthly,
					}},
				}},
			}})
			return nil
		}
	}

	d.SetId("")

	return fmt.Errorf("Instance Type %s was not found", reqType)
}
