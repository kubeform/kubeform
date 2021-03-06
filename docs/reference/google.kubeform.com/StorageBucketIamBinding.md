---
title: StorageBucketIamBinding
menu:
  docs_{{ .version }}:
    identifier: storagebucketiambinding-google.kubeform.com
    name: StorageBucketIamBinding
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageBucketIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageBucketIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageBucketIamBindingSpec](#storagebucketiambindingspec)***||
| `status` | ***[StorageBucketIamBindingStatus](#storagebucketiambindingstatus)***||
## Phase(`string` alias)

Appears on:[StorageBucketIamBindingStatus](#storagebucketiambindingstatus)

## StorageBucketIamBindingSpec

Appears on:[StorageBucketIamBinding](#storagebucketiambinding), [StorageBucketIamBindingStatus](#storagebucketiambindingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `members` | ***[]string***||
| `role` | ***string***||
## StorageBucketIamBindingStatus

Appears on:[StorageBucketIamBinding](#storagebucketiambinding)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageBucketIamBindingSpec](#storagebucketiambindingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
