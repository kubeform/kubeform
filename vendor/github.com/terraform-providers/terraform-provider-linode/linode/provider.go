package linode

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/logging"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/httpclient"
	"github.com/hashicorp/terraform/terraform"
	"github.com/linode/linodego"
	"golang.org/x/oauth2"
)

// DefaultLinodeURL is the Linode APIv4 URL to use
const DefaultLinodeURL = "https://api.linode.com/v4"

// Provider creates and manages the resources in a Linode configuration.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LINODE_TOKEN", nil),
				Description: "The token that allows you access to your Linode account",
			},
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LINODE_URL", nil),
				Description: "The HTTP(S) API address of the Linode API to use.",
			},
			"ua_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LINODE_UA_PREFIX", nil),
				Description: "An HTTP User-Agent Prefix to prepend in API requests.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"linode_account":       dataSourceLinodeAccount(),
			"linode_domain":        dataSourceLinodeDomain(),
			"linode_image":         dataSourceLinodeImage(),
			"linode_instance_type": dataSourceLinodeInstanceType(),
			"linode_networking_ip": dataSourceLinodeNetworkingIP(),
			"linode_profile":       dataSourceLinodeProfile(),
			"linode_region":        dataSourceLinodeRegion(),
			"linode_sshkey":        dataSourceLinodeSSHKey(),
			"linode_user":          dataSourceLinodeUser(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"linode_image":               resourceLinodeImage(),
			"linode_instance":            resourceLinodeInstance(),
			"linode_domain":              resourceLinodeDomain(),
			"linode_domain_record":       resourceLinodeDomainRecord(),
			"linode_nodebalancer":        resourceLinodeNodeBalancer(),
			"linode_nodebalancer_config": resourceLinodeNodeBalancerConfig(),
			"linode_nodebalancer_node":   resourceLinodeNodeBalancerNode(),
			"linode_rdns":                resourceLinodeRDNS(),
			"linode_sshkey":              resourceLinodeSSHKey(),
			"linode_stackscript":         resourceLinodeStackscript(),
			"linode_token":               resourceLinodeToken(),
			"linode_volume":              resourceLinodeVolume(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	token, ok := d.Get("token").(string)
	if !ok {
		return nil, fmt.Errorf("The Linode API Token was not valid")
	}

	url, ok := d.Get("url").(string)
	if !ok {
		return nil, fmt.Errorf("The Linode API URL was not valid")
	}

	uaPrefix, ok := d.Get("ua_prefix").(string)
	if !ok {
		return nil, fmt.Errorf("The Linode UA Prefix was not valid")
	}

	client := getLinodeClient(token, url, uaPrefix)
	// Ping the API for an empty response to verify the configuration works
	_, err := client.ListTypes(context.Background(), linodego.NewListOptions(100, ""))
	if err != nil {
		return nil, fmt.Errorf("Error connecting to the Linode API: %s", err)
	}

	return client, nil
}

func getLinodeClient(token, url, uaPrefix string) linodego.Client {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})

	oauthTransport := &oauth2.Transport{
		Source: tokenSource,
	}
	loggingTransport := logging.NewTransport("Linode", oauthTransport)
	oauth2Client := &http.Client{
		Transport: loggingTransport,
	}

	client := linodego.NewClient(oauth2Client)

	terraformVersion := httpclient.UserAgentString()
	projectURL := "(+https://www.terraform.io)"
	sdkVersion := fmt.Sprintf("linodego/%s", linodego.Version)
	userAgent := fmt.Sprintf("%s %s %s", terraformVersion, projectURL, sdkVersion)

	if len(uaPrefix) > 0 {
		userAgent = uaPrefix + " " + userAgent
	}

	var baseURL = DefaultLinodeURL
	if len(url) > 0 {
		baseURL = url
	}

	client.SetBaseURL(baseURL)
	client.SetUserAgent(userAgent)
	return client
}
