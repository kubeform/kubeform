---
title: AzureadServicePrincipal
menu:
  docs_{{ .version }}:
    identifier: azureadserviceprincipal-azurerm.kubeform.com
    name: AzureadServicePrincipal
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AzureadServicePrincipal
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AzureadServicePrincipal` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AzureadServicePrincipalSpec](#azureadserviceprincipalspec)***||
| `status` | ***[AzureadServicePrincipalStatus](#azureadserviceprincipalstatus)***||
## AzureadServicePrincipalSpec

Appears on:[AzureadServicePrincipal](#azureadserviceprincipal), [AzureadServicePrincipalStatus](#azureadserviceprincipalstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `applicationID` | ***string***||
| `displayName` | ***string***| ***(Optional)*** |
## AzureadServicePrincipalStatus

Appears on:[AzureadServicePrincipal](#azureadserviceprincipal)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AzureadServicePrincipalSpec](#azureadserviceprincipalspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AzureadServicePrincipalStatus](#azureadserviceprincipalstatus)

---
