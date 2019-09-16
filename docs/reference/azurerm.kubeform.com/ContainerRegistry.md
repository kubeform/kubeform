---
title: ContainerRegistry
menu:
  docs_{{ .version }}:
    identifier: containerregistry-azurerm.kubeform.com
    name: ContainerRegistry
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ContainerRegistry
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ContainerRegistry` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ContainerRegistrySpec](#ContainerRegistrySpec)***||
| `status` | ***[ContainerRegistryStatus](#ContainerRegistryStatus)***||
## ContainerRegistrySpec

Appears on:[ContainerRegistry](#ContainerRegistry), [ContainerRegistryStatus](#ContainerRegistryStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `adminEnabled` | ***bool***| ***(Optional)*** |
| `adminUsername` | ***string***| ***(Optional)*** |
| `georeplicationLocations` | ***[]string***| ***(Optional)*** |
| `location` | ***string***||
| `loginServer` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `networkRuleSet` | ***[[]ContainerRegistrySpecNetworkRuleSet](#ContainerRegistrySpecNetworkRuleSet)***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `sku` | ***string***| ***(Optional)*** |
| `storageAccount` | ***[[]ContainerRegistrySpecStorageAccount](#ContainerRegistrySpecStorageAccount)***| ***(Optional)*** Deprecated|
| `storageAccountID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## ContainerRegistrySpecNetworkRuleSet

Appears on:[ContainerRegistrySpec](#ContainerRegistrySpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultAction` | ***string***| ***(Optional)*** |
| `ipRule` | ***[[]ContainerRegistrySpecNetworkRuleSetIpRule](#ContainerRegistrySpecNetworkRuleSetIpRule)***| ***(Optional)*** |
## ContainerRegistrySpecNetworkRuleSetIpRule

Appears on:[ContainerRegistrySpecNetworkRuleSet](#ContainerRegistrySpecNetworkRuleSet)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***||
| `ipRange` | ***string***||
## ContainerRegistrySpecStorageAccount

Appears on:[ContainerRegistrySpec](#ContainerRegistrySpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
## ContainerRegistryStatus

Appears on:[ContainerRegistry](#ContainerRegistry)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ContainerRegistrySpec](#ContainerRegistrySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `admin_password` | ***string*** ||
| `storage_account.<index>.access_key` | ***string*** ||
