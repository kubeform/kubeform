package linode

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/httpclient"
	"github.com/linode/linodego"
	"github.com/terraform-providers/terraform-provider-linode/version"
	"golang.org/x/oauth2"
)

// DefaultLinodeURL is the Linode APIv4 URL to use
const DefaultLinodeURL = "https://api.linode.com/v4"

// Config represents the Linode provider configuration
type Config struct {
	AccessToken string
	APIURL      string
	APIVersion  string
	UAPrefix    string

	terraformVersion string
}

// Client returns a fully initialized Linode client
func (c *Config) Client() linodego.Client {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: c.AccessToken})
	oauthTransport := &oauth2.Transport{
		Source: tokenSource,
	}
	loggingTransport := logging.NewTransport("Linode", oauthTransport)

	oauth2Client := &http.Client{
		Transport: loggingTransport,
	}
	client := linodego.NewClient(oauth2Client)

	tfUserAgent := httpclient.TerraformUserAgent(c.terraformVersion)
	userAgent := strings.TrimSpace(fmt.Sprintf("%s terraform-provider-linode/%s linodego/%s",
		tfUserAgent, version.ProviderVersion, linodego.Version))
	if c.UAPrefix != "" {
		userAgent = c.UAPrefix + " " + userAgent
	}
	client.SetUserAgent(userAgent)

	if c.APIURL != "" {
		client.SetBaseURL(c.APIURL)
	} else if len(c.APIVersion) > 0 {
		client.SetAPIVersion(c.APIVersion)
	} else {
		client.SetBaseURL(DefaultLinodeURL)
	}

	return client
}
