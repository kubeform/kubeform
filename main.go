/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"strings"

	"kubeform.dev/kubeform/util"

	"github.com/gobuffalo/flect"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm"
	"github.com/terraform-providers/terraform-provider-digitalocean/digitalocean"
	"github.com/terraform-providers/terraform-provider-google/google"
	"github.com/terraform-providers/terraform-provider-linode/linode"
)

func main() {
	version := "v1alpha1"

	providersMap := map[string]terraform.ResourceProvider{
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
			providerPrefix := key
			for key, values := range p.ResourcesMap {
				key = strings.TrimPrefix(key, providerPrefix+"_")
				structNames = append(structNames, flect.Capitalize(flect.Camelize(key)))
				schemas = append(schemas, values.Schema)
			}

			err := util.GenerateProviderAPIS(key, version, schemas, structNames)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

	err := util.GenerateProviderAPIS("modules", version, nil, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
