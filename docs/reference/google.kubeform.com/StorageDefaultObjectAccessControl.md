---
title: StorageDefaultObjectAccessControl
menu:
  docs_{{ .version }}:
    identifier: storagedefaultobjectaccesscontrol-google.kubeform.com
    name: StorageDefaultObjectAccessControl
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageDefaultObjectAccessControl
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageDefaultObjectAccessControl` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageDefaultObjectAccessControlSpec](#StorageDefaultObjectAccessControlSpec)***||
| `status` | ***[StorageDefaultObjectAccessControlStatus](#StorageDefaultObjectAccessControlStatus)***||
## StorageDefaultObjectAccessControlSpec
##### (Appears on:[StorageDefaultObjectAccessControl](#StorageDefaultObjectAccessControl), [StorageDefaultObjectAccessControlStatus](#StorageDefaultObjectAccessControlStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `domain` | ***string***| ***(Optional)*** |
| `email` | ***string***| ***(Optional)*** |
| `entity` | ***string***||
| `entityID` | ***string***| ***(Optional)*** |
| `generation` | ***int***| ***(Optional)*** |
| `object` | ***string***| ***(Optional)*** |
| `projectTeam` | ***[[]StorageDefaultObjectAccessControlSpecProjectTeam](#StorageDefaultObjectAccessControlSpecProjectTeam)***| ***(Optional)*** |
| `role` | ***string***||
## StorageDefaultObjectAccessControlSpecProjectTeam
##### (Appears on:[StorageDefaultObjectAccessControlSpec](#StorageDefaultObjectAccessControlSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `projectNumber` | ***string***| ***(Optional)*** |
| `team` | ***string***| ***(Optional)*** |
## StorageDefaultObjectAccessControlStatus
##### (Appears on:[StorageDefaultObjectAccessControl](#StorageDefaultObjectAccessControl))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageDefaultObjectAccessControlSpec](#StorageDefaultObjectAccessControlSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
