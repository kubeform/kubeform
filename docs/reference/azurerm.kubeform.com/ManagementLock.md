---
title: ManagementLock
menu:
  docs_{{ .version }}:
    identifier: managementlock-azurerm.kubeform.com
    name: ManagementLock
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ManagementLock
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ManagementLock` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ManagementLockSpec](#managementlockspec)***||
| `status` | ***[ManagementLockStatus](#managementlockstatus)***||
## ManagementLockSpec

Appears on:[ManagementLock](#managementlock), [ManagementLockStatus](#managementlockstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `lockLevel` | ***string***||
| `name` | ***string***||
| `notes` | ***string***| ***(Optional)*** |
| `scope` | ***string***||
## ManagementLockStatus

Appears on:[ManagementLock](#managementlock)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ManagementLockSpec](#managementlockspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ManagementLockStatus](#managementlockstatus)

---
