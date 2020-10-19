---
title: ObjectStorageBucket
menu:
  docs_{{ .version }}:
    identifier: objectstoragebucket-linode.kubeform.com
    name: ObjectStorageBucket
    parent: linode.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ObjectStorageBucket
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `linode.kubeform.com/v1alpha1` |
|    `kind` | string | `ObjectStorageBucket` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ObjectStorageBucketSpec](#objectstoragebucketspec)***||
| `status` | ***[ObjectStorageBucketStatus](#objectstoragebucketstatus)***||
## ObjectStorageBucketSpec

Appears on:[ObjectStorageBucket](#objectstoragebucket), [ObjectStorageBucketStatus](#objectstoragebucketstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cluster` | ***string***|The cluster of the Linode Object Storage Bucket.|
| `label` | ***string***|The label of the Linode Object Storage Bucket.|
## ObjectStorageBucketStatus

Appears on:[ObjectStorageBucket](#objectstoragebucket)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ObjectStorageBucketSpec](#objectstoragebucketspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ObjectStorageBucketStatus](#objectstoragebucketstatus)

---
