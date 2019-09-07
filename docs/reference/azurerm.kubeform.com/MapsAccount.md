---
title: MapsAccount
menu:
  docs_v0.0.1:
    identifier: mapsaccount-azurerm.kubeform.com
    name: MapsAccount
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## MapsAccount
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MapsAccount` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MapsAccountSpec](#MapsAccountSpec)***||
| `status` | ***[MapsAccountStatus](#MapsAccountStatus)***||
## MapsAccountSpec
##### (Appears on:[MapsAccount](#MapsAccount), [MapsAccountStatus](#MapsAccountStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `skuName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `xMsClientID` | ***string***| ***(Optional)*** |
## MapsAccountStatus
##### (Appears on:[MapsAccount](#MapsAccount))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MapsAccountSpec](#MapsAccountSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_access_key` | ***string*** ||
| `secondary_access_key` | ***string*** ||
