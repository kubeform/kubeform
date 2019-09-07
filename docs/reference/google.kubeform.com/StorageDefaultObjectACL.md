---
title: StorageDefaultObjectACL
menu:
  docs_v0.0.1:
    identifier: storagedefaultobjectacl-google.kubeform.com
    name: StorageDefaultObjectACL
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## StorageDefaultObjectACL
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageDefaultObjectACL` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageDefaultObjectACLSpec](#StorageDefaultObjectACLSpec)***||
| `status` | ***[StorageDefaultObjectACLStatus](#StorageDefaultObjectACLStatus)***||
## StorageDefaultObjectACLSpec
##### (Appears on:[StorageDefaultObjectACL](#StorageDefaultObjectACL), [StorageDefaultObjectACLStatus](#StorageDefaultObjectACLStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `roleEntity` | ***[]string***| ***(Optional)*** |
## StorageDefaultObjectACLStatus
##### (Appears on:[StorageDefaultObjectACL](#StorageDefaultObjectACL))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageDefaultObjectACLSpec](#StorageDefaultObjectACLSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
