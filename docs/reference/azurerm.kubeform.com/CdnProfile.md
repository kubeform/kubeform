---
title: CdnProfile
menu:
  docs_v0.0.1:
    identifier: cdnprofile-azurerm.kubeform.com
    name: CdnProfile
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## CdnProfile
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `CdnProfile` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CdnProfileSpec](#CdnProfileSpec)***||
| `status` | ***[CdnProfileStatus](#CdnProfileStatus)***||
## CdnProfileSpec
##### (Appears on:[CdnProfile](#CdnProfile), [CdnProfileStatus](#CdnProfileStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sku` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## CdnProfileStatus
##### (Appears on:[CdnProfile](#CdnProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CdnProfileSpec](#CdnProfileSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
