---
title: RelayNamespace
menu:
  docs_{{ .version }}:
    identifier: relaynamespace-azurerm.kubeform.com
    name: RelayNamespace
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RelayNamespace
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `RelayNamespace` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RelayNamespaceSpec](#RelayNamespaceSpec)***||
| `status` | ***[RelayNamespaceStatus](#RelayNamespaceStatus)***||
## RelayNamespaceSpec

Appears on:[RelayNamespace](#RelayNamespace), [RelayNamespaceStatus](#RelayNamespaceStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `location` | ***string***||
| `metricID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sku` | ***[[]RelayNamespaceSpecSku](#RelayNamespaceSpecSku)***| ***(Optional)*** Deprecated|
| `skuName` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## RelayNamespaceSpecSku

Appears on:[RelayNamespaceSpec](#RelayNamespaceSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
## RelayNamespaceStatus

Appears on:[RelayNamespace](#RelayNamespace)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RelayNamespaceSpec](#RelayNamespaceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_connection_string` | ***string*** ||
| `primary_key` | ***string*** ||
| `secondary_connection_string` | ***string*** ||
| `secondary_key` | ***string*** ||
