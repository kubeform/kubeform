package main

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-digitalocean/digitalocean"

	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-linode/linode"

	"kubeform.dev/kubeform/util"

	"github.com/hashicorp/terraform/helper/schema"
)

func main() {
	var providersMap map[string]terraform.ResourceProvider

	version := "v1alpha1"

	providersMap = map[string]terraform.ResourceProvider{
		"linode":       linode.Provider(),
		"digitalocean": digitalocean.Provider(),
	}

	for key, provider := range providersMap {
		p, ok := provider.(*schema.Provider)
		if ok {
			var structNames []string
			var schemas []map[string]*schema.Schema

			for key, values := range p.ResourcesMap {
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
