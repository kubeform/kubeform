---
title: AzureadServicePrincipalPassword
menu:
  docs_{{ .version }}:
    identifier: azureadserviceprincipalpassword-azurerm.kubeform.com
    name: AzureadServicePrincipalPassword
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AzureadServicePrincipalPassword
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AzureadServicePrincipalPassword` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AzureadServicePrincipalPasswordSpec](#azureadserviceprincipalpasswordspec)***||
| `status` | ***[AzureadServicePrincipalPasswordStatus](#azureadserviceprincipalpasswordstatus)***||
## AzureadServicePrincipalPasswordSpec

Appears on:[AzureadServicePrincipalPassword](#azureadserviceprincipalpassword), [AzureadServicePrincipalPasswordStatus](#azureadserviceprincipalpasswordstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `endDate` | ***string***||
| `keyID` | ***string***| ***(Optional)*** |
| `servicePrincipalID` | ***string***||
| `startDate` | ***string***| ***(Optional)*** |
## AzureadServicePrincipalPasswordStatus

Appears on:[AzureadServicePrincipalPassword](#azureadserviceprincipalpassword)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AzureadServicePrincipalPasswordSpec](#azureadserviceprincipalpasswordspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AzureadServicePrincipalPasswordStatus](#azureadserviceprincipalpasswordstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `value` | ***string*** ||
