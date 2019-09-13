---
title: SecurityCenterContact
menu:
  docs_{{ .version }}:
    identifier: securitycentercontact-azurerm.kubeform.com
    name: SecurityCenterContact
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SecurityCenterContact
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SecurityCenterContact` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SecurityCenterContactSpec](#SecurityCenterContactSpec)***||
| `status` | ***[SecurityCenterContactStatus](#SecurityCenterContactStatus)***||
## SecurityCenterContactSpec
##### (Appears on:[SecurityCenterContact](#SecurityCenterContact), [SecurityCenterContactStatus](#SecurityCenterContactStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `alertNotifications` | ***bool***||
| `alertsToAdmins` | ***bool***||
| `email` | ***string***||
| `phone` | ***string***| ***(Optional)*** |
## SecurityCenterContactStatus
##### (Appears on:[SecurityCenterContact](#SecurityCenterContact))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SecurityCenterContactSpec](#SecurityCenterContactSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
