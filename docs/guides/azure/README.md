---
title: Azure
menu:
  docs_{{ .version }}:
    identifier: readme-azure
    name: Overview
    parent: azure-guides
    weight: 10
menu_name: docs_{{ .version }}
section_menu_id: guides
url: /docs/{{ .version }}/guides/azure/
aliases:
  - /docs/{{ .version }}/guides/azure/README/
---

# Azure

This guide will show you how to provision a Azure Redis Cache using Kubeform.

> Examples used in this guide can be found [here](https://github.com/kubeform/docs/tree/{{< param "info.version" >}}/docs/examples/azure).

Look at the `Terraform` configuration below:

```
provider "azurerm" {
  subscription_id = "Subscription ID"
  client_id       = "Client ID"
  client_secret   = "Client Secret"
  tenant_id       = "Tenant ID"
}

resource "azurerm_redis_cache" "example" {
  name                = "example-cache"
  resource_group_name = "dev"
  location            = "East US"
  capacity            = 2
  family              = "C"
  sku_name            = "Standard"
  enable_non_ssl_port = false
  minimum_tls_version = "1.2"
}
```

This config creates an Azure Redis Cache. We'll create the exact configuration using `kubeform`. The steps are given below:

## 1. Create CRD:

At first, create the CRD of Azure Redis Cache using the following yaml:

```yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: rediscaches.azurerm.kubeform.com
spec:
  group: azurerm.kubeform.com
  version: v1alpha1
  names:
    kind: RedisCache
    plural: rediscaches
  scope: Namespaced
```

Save it in a file (eg. `crd.yaml`) then apply it using kubectl.

```bash
$ kubectl apply -f crd.yaml
```

## 2. Create Azure Provider Secret

Then create the secret which is necessary for provisioning the Redis Cache in Azure.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: azure
type: "kubeform.com/azurerm"
data:
  subscription_id: <Base64 encoded value of Azure Subscription ID>
  client_id: <Base64 encoded value of Azure Client ID>
  client_secret: <Base64 encoded value of Azure Client Secret>
  tenant_id: <Base64 encoded value of Azure Tenant ID>

```

Here we can see that, the data of the secret is same as the field of the provider part in the terraform config file. Save it in a file (eg. `secret.yaml`) then apply it using kubectl.

```bash
$ kubectl apply -f secret.yaml
```
> **Note:** here, data key (eg. `client_id`, `tenant_id` etc.) must be in snake case format (same as the tf configuration file)


## 3. Create Azure Redis Cache

Now, we'll create the Azure Redis Cache CRD. The yaml is given below:

```yaml
apiVersion: azurerm.kubeform.com/v1alpha1
kind: RedisCache
metadata:
  name: redis-test-1
spec:
  name: example-cache
  resourceGroupName: dev
  location: East US
  capacity: 2
  family: C
  skuName: Standard
  enableNonSslPort: false
  minimumTlsVersion: "1.2"
  providerRef:
    name: azure
```

Here, we can see that the provider secret is referenced using a field called `providerRef`.

Save it in a file (eg. `redis.yaml`) then apply it using kubectl.

```bash
$ kubectl apply -f redis.yaml
```

After that, an Azure Redis Cache will be created!

## Delete Azure Redis Cache

To delete the Azure Redis Cache just run:

```bash
kubectl delete -f redis.yaml
```
