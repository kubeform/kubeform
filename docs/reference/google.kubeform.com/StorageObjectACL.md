---
title: StorageObjectACL
menu:
  docs_{{ .version }}:
    identifier: storageobjectacl-google.kubeform.com
    name: StorageObjectACL
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageObjectACL
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageObjectACL` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageObjectACLSpec](#StorageObjectACLSpec)***||
| `status` | ***[StorageObjectACLStatus](#StorageObjectACLStatus)***||
## StorageObjectACLSpec
##### (Appears on:[StorageObjectACL](#StorageObjectACL), [StorageObjectACLStatus](#StorageObjectACLStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `object` | ***string***||
| `predefinedACL` | ***string***| ***(Optional)*** |
| `roleEntity` | ***[]string***| ***(Optional)*** |
## StorageObjectACLStatus
##### (Appears on:[StorageObjectACL](#StorageObjectACL))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageObjectACLSpec](#StorageObjectACLSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
