---
title: AzureadServicePrincipal
menu:
  docs_v0.0.1:
    identifier: azureadserviceprincipal-azurerm.kubeform.com
    name: AzureadServicePrincipal
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## AzureadServicePrincipal
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AzureadServicePrincipal` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AzureadServicePrincipalSpec](#AzureadServicePrincipalSpec)***||
| `status` | ***[AzureadServicePrincipalStatus](#AzureadServicePrincipalStatus)***||
## AzureadServicePrincipalSpec
##### (Appears on:[AzureadServicePrincipal](#AzureadServicePrincipal), [AzureadServicePrincipalStatus](#AzureadServicePrincipalStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `applicationID` | ***string***||
| `displayName` | ***string***| ***(Optional)*** |
## AzureadServicePrincipalStatus
##### (Appears on:[AzureadServicePrincipal](#AzureadServicePrincipal))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AzureadServicePrincipalSpec](#AzureadServicePrincipalSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
