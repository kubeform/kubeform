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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SecurityCenterContactSpec](#securitycentercontactspec)***||
| `status` | ***[SecurityCenterContactStatus](#securitycentercontactstatus)***||
## Phase(`string` alias)

Appears on:[SecurityCenterContactStatus](#securitycentercontactstatus)

## SecurityCenterContactSpec

Appears on:[SecurityCenterContact](#securitycentercontact), [SecurityCenterContactStatus](#securitycentercontactstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `alertNotifications` | ***bool***||
| `alertsToAdmins` | ***bool***||
| `email` | ***string***||
| `phone` | ***string***| ***(Optional)*** |
## SecurityCenterContactStatus

Appears on:[SecurityCenterContact](#securitycentercontact)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SecurityCenterContactSpec](#securitycentercontactspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
