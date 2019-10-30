---
title: StorageBucketObject
menu:
  docs_{{ .version }}:
    identifier: storagebucketobject-google.kubeform.com
    name: StorageBucketObject
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageBucketObject
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageBucketObject` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageBucketObjectSpec](#storagebucketobjectspec)***||
| `status` | ***[StorageBucketObjectStatus](#storagebucketobjectstatus)***||
## Phase(`string` alias)

Appears on:[StorageBucketObjectStatus](#storagebucketobjectstatus)

## StorageBucketObjectSpec

Appears on:[StorageBucketObject](#storagebucketobject), [StorageBucketObjectStatus](#storagebucketobjectstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `cacheControl` | ***string***| ***(Optional)*** |
| `content` | ***string***| ***(Optional)*** |
| `contentDisposition` | ***string***| ***(Optional)*** |
| `contentEncoding` | ***string***| ***(Optional)*** |
| `contentLanguage` | ***string***| ***(Optional)*** |
| `contentType` | ***string***| ***(Optional)*** |
| `crc32c` | ***string***| ***(Optional)*** |
| `detectMd5hash` | ***string***| ***(Optional)*** |
| `md5hash` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `source` | ***string***| ***(Optional)*** |
| `storageClass` | ***string***| ***(Optional)*** |
## StorageBucketObjectStatus

Appears on:[StorageBucketObject](#storagebucketobject)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageBucketObjectSpec](#storagebucketobjectspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
