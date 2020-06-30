package linode

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/linode/linodego"
)

// Provider creates and manages the resources in a Linode configuration.
func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LINODE_TOKEN", nil),
				Description: "The token that allows you access to your Linode account",
			},
			"url": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("LINODE_URL", nil),
				Description:  "The HTTP(S) API address of the Linode API to use.",
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"ua_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LINODE_UA_PREFIX", nil),
				Description: "An HTTP User-Agent Prefix to prepend in API requests.",
			},
			"api_version": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LINODE_API_VERSION", nil),
				Description: "An HTTP User-Agent Prefix to prepend in API requests.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"linode_account":                dataSourceLinodeAccount(),
			"linode_domain":                 dataSourceLinodeDomain(),
			"linode_domain_record":          dataSourceLinodeDomainRecord(),
			"linode_image":                  dataSourceLinodeImage(),
			"linode_instance_type":          dataSourceLinodeInstanceType(),
			"linode_networking_ip":          dataSourceLinodeNetworkingIP(),
			"linode_object_storage_cluster": dataSourceLinodeObjectStorageCluster(),
			"linode_profile":                dataSourceLinodeProfile(),
			"linode_region":                 dataSourceLinodeRegion(),
			"linode_sshkey":                 dataSourceLinodeSSHKey(),
			"linode_stackscript":            dataSourceLinodeStackscript(),
			"linode_user":                   dataSourceLinodeUser(),
			"linode_volume":                 dataSourceLinodeVolume(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"linode_domain":                resourceLinodeDomain(),
			"linode_domain_record":         resourceLinodeDomainRecord(),
			"linode_firewall":              resourceLinodeFirewall(),
			"linode_image":                 resourceLinodeImage(),
			"linode_instance":              resourceLinodeInstance(),
			"linode_lke_cluster":           resourceLinodeLKECluster(),
			"linode_nodebalancer":          resourceLinodeNodeBalancer(),
			"linode_nodebalancer_config":   resourceLinodeNodeBalancerConfig(),
			"linode_nodebalancer_node":     resourceLinodeNodeBalancerNode(),
			"linode_object_storage_bucket": resourceLinodeObjectStorageBucket(),
			"linode_object_storage_key":    resourceLinodeObjectStorageKey(),
			"linode_rdns":                  resourceLinodeRDNS(),
			"linode_sshkey":                resourceLinodeSSHKey(),
			"linode_stackscript":           resourceLinodeStackscript(),
			"linode_token":                 resourceLinodeToken(),
			"linode_volume":                resourceLinodeVolume(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}
	return provider
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	config := &Config{
		AccessToken: d.Get("token").(string),
		APIURL:      d.Get("url").(string),
		APIVersion:  d.Get("api_version").(string),
		UAPrefix:    d.Get("ua_prefix").(string),
	}
	config.terraformVersion = terraformVersion
	client := config.Client()

	// Ping the API for an empty response to verify the configuration works
	if _, err := client.ListTypes(context.Background(), linodego.NewListOptions(100, "")); err != nil {
		return nil, fmt.Errorf("Error connecting to the Linode API: %s", err)
	}
	return config.Client(), nil
}
