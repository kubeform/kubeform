---
title: FolderIamBinding
menu:
  docs_{{ .version }}:
    identifier: folderiambinding-google.kubeform.com
    name: FolderIamBinding
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## FolderIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `FolderIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FolderIamBindingSpec](#FolderIamBindingSpec)***||
| `status` | ***[FolderIamBindingStatus](#FolderIamBindingStatus)***||
## FolderIamBindingSpec
##### (Appears on:[FolderIamBinding](#FolderIamBinding), [FolderIamBindingStatus](#FolderIamBindingStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `folder` | ***string***||
| `members` | ***[]string***||
| `role` | ***string***||
## FolderIamBindingStatus
##### (Appears on:[FolderIamBinding](#FolderIamBinding))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FolderIamBindingSpec](#FolderIamBindingSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
