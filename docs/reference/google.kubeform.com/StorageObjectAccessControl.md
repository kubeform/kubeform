---
title: StorageObjectAccessControl
menu:
  docs_{{ .version }}:
    identifier: storageobjectaccesscontrol-google.kubeform.com
    name: StorageObjectAccessControl
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageObjectAccessControl
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageObjectAccessControl` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageObjectAccessControlSpec](#StorageObjectAccessControlSpec)***||
| `status` | ***[StorageObjectAccessControlStatus](#StorageObjectAccessControlStatus)***||
## StorageObjectAccessControlSpec

Appears on:[StorageObjectAccessControl](#StorageObjectAccessControl), [StorageObjectAccessControlStatus](#StorageObjectAccessControlStatus)

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
| `object` | ***string***||
| `projectTeam` | ***[[]StorageObjectAccessControlSpecProjectTeam](#StorageObjectAccessControlSpecProjectTeam)***| ***(Optional)*** |
| `role` | ***string***||
## StorageObjectAccessControlSpecProjectTeam

Appears on:[StorageObjectAccessControlSpec](#StorageObjectAccessControlSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `projectNumber` | ***string***| ***(Optional)*** |
| `team` | ***string***| ***(Optional)*** |
## StorageObjectAccessControlStatus

Appears on:[StorageObjectAccessControl](#StorageObjectAccessControl)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageObjectAccessControlSpec](#StorageObjectAccessControlSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
