---
title: ComputeBackendBucket
menu:
  docs_{{ .version }}:
    identifier: computebackendbucket-google.kubeform.com
    name: ComputeBackendBucket
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeBackendBucket
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeBackendBucket` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeBackendBucketSpec](#computebackendbucketspec)***||
| `status` | ***[ComputeBackendBucketStatus](#computebackendbucketstatus)***||
## ComputeBackendBucketSpec

Appears on:[ComputeBackendBucket](#computebackendbucket), [ComputeBackendBucketStatus](#computebackendbucketstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucketName` | ***string***||
| `cdnPolicy` | ***[[]ComputeBackendBucketSpecCdnPolicy](#computebackendbucketspeccdnpolicy)***| ***(Optional)*** |
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `enableCdn` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
## ComputeBackendBucketSpecCdnPolicy

Appears on:[ComputeBackendBucketSpec](#computebackendbucketspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `signedURLCacheMaxAgeSec` | ***int64***| ***(Optional)*** |
## ComputeBackendBucketStatus

Appears on:[ComputeBackendBucket](#computebackendbucket)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeBackendBucketSpec](#computebackendbucketspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeBackendBucketStatus](#computebackendbucketstatus)

---
