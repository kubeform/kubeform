package linode

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/linode/linodego"
)

func dataSourceLinodeObjectStorageCluster() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLinodeObjectStorageClusterRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "The unique ID of this Cluster.",
				Required:    true,
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The base URL for this cluster.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This cluster's status.",
				Computed:    true,
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The region this cluster is located in.",
				Computed:    true,
			},
			"static_site_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The base URL for this cluster used when hosting static sites.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLinodeObjectStorageClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(linodego.Client)

	reqObjectStorageCluster := d.Get("id").(string)

	if reqObjectStorageCluster == "" {
		return fmt.Errorf("Error object storage cluster id is required")
	}

	objectStorageCluster, err := client.GetObjectStorageCluster(context.Background(), reqObjectStorageCluster)
	if err != nil {
		return fmt.Errorf("Error listing object storage clusters: %s", err)
	}

	if objectStorageCluster != nil {
		d.SetId(objectStorageCluster.ID)
		d.Set("domain", objectStorageCluster.Domain)
		d.Set("status", objectStorageCluster.Status)
		d.Set("region", objectStorageCluster.Region)
		d.Set("static_site_domain", objectStorageCluster.StaticSiteDomain)

		return nil
	}

	return fmt.Errorf("Linode object storage cluster %s was not found", reqObjectStorageCluster)
}
