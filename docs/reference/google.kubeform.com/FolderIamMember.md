---
title: FolderIamMember
menu:
  docs_v0.0.1:
    identifier: folderiammember-google.kubeform.com
    name: FolderIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## FolderIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `FolderIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FolderIamMemberSpec](#FolderIamMemberSpec)***||
| `status` | ***[FolderIamMemberStatus](#FolderIamMemberStatus)***||
## FolderIamMemberSpec
##### (Appears on:[FolderIamMember](#FolderIamMember), [FolderIamMemberStatus](#FolderIamMemberStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `folder` | ***string***||
| `member` | ***string***||
| `role` | ***string***||
## FolderIamMemberStatus
##### (Appears on:[FolderIamMember](#FolderIamMember))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FolderIamMemberSpec](#FolderIamMemberSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
