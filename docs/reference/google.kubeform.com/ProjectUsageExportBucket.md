---
title: ProjectUsageExportBucket
menu:
  docs_v0.0.1:
    identifier: projectusageexportbucket-google.kubeform.com
    name: ProjectUsageExportBucket
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ProjectUsageExportBucket
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ProjectUsageExportBucket` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ProjectUsageExportBucketSpec](#ProjectUsageExportBucketSpec)***||
| `status` | ***[ProjectUsageExportBucketStatus](#ProjectUsageExportBucketStatus)***||
## ProjectUsageExportBucketSpec
##### (Appears on:[ProjectUsageExportBucket](#ProjectUsageExportBucket), [ProjectUsageExportBucketStatus](#ProjectUsageExportBucketStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucketName` | ***string***||
| `prefix` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
## ProjectUsageExportBucketStatus
##### (Appears on:[ProjectUsageExportBucket](#ProjectUsageExportBucket))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ProjectUsageExportBucketSpec](#ProjectUsageExportBucketSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
