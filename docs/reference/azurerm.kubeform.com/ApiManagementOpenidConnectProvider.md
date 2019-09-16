---
title: ApiManagementOpenidConnectProvider
menu:
  docs_{{ .version }}:
    identifier: apimanagementopenidconnectprovider-azurerm.kubeform.com
    name: ApiManagementOpenidConnectProvider
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiManagementOpenidConnectProvider
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementOpenidConnectProvider` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementOpenidConnectProviderSpec](#ApiManagementOpenidConnectProviderSpec)***||
| `status` | ***[ApiManagementOpenidConnectProviderStatus](#ApiManagementOpenidConnectProviderStatus)***||
## ApiManagementOpenidConnectProviderSpec

Appears on:[ApiManagementOpenidConnectProvider](#ApiManagementOpenidConnectProvider), [ApiManagementOpenidConnectProviderStatus](#ApiManagementOpenidConnectProviderStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `apiManagementName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `displayName` | ***string***||
| `metadataEndpoint` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
## ApiManagementOpenidConnectProviderStatus

Appears on:[ApiManagementOpenidConnectProvider](#ApiManagementOpenidConnectProvider)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementOpenidConnectProviderSpec](#ApiManagementOpenidConnectProviderSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `client_id` | ***string*** ||
| `client_secret` | ***string*** ||
