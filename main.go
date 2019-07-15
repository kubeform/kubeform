package main

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm"
	"github.com/terraform-providers/terraform-provider-digitalocean/digitalocean"
	"github.com/terraform-providers/terraform-provider-google/google"
	"github.com/terraform-providers/terraform-provider-linode/linode"
	"kubeform.dev/kubeform/util"
)

func main() {
	var providersMap map[string]terraform.ResourceProvider

	version := "v1alpha1"

	providersMap = map[string]terraform.ResourceProvider{
		"linode":       linode.Provider(),
		"digitalocean": digitalocean.Provider(),
		"aws":          aws.Provider(),
		"google":       google.Provider(),
		"azurerm":      azurerm.Provider(),
	}

	for key, provider := range providersMap {
		p, ok := provider.(*schema.Provider)
		if ok {
			var structNames []string
			var schemas []map[string]*schema.Schema

			for key, values := range p.ResourcesMap {
				providerPrefix := strings.Split(key, "_")[0]
				key = strings.Replace(key, providerPrefix+"_", "", 1)
				structNames = append(structNames, util.SnakeCaseToCamelCase(key))
				schemas = append(schemas, values.Schema)
			}

			err := util.GenerateProviderAPIS(key, version, schemas, structNames)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
