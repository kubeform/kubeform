---
title: StorageBucketIamMember
menu:
  docs_{{ .version }}:
    identifier: storagebucketiammember-google.kubeform.com
    name: StorageBucketIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageBucketIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageBucketIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageBucketIamMemberSpec](#StorageBucketIamMemberSpec)***||
| `status` | ***[StorageBucketIamMemberStatus](#StorageBucketIamMemberStatus)***||
## StorageBucketIamMemberSpec

Appears on:[StorageBucketIamMember](#StorageBucketIamMember), [StorageBucketIamMemberStatus](#StorageBucketIamMemberStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `member` | ***string***||
| `role` | ***string***||
## StorageBucketIamMemberStatus

Appears on:[StorageBucketIamMember](#StorageBucketIamMember)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageBucketIamMemberSpec](#StorageBucketIamMemberSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
