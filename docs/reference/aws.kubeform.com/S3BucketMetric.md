---
title: S3BucketMetric
menu:
  docs_{{ .version }}:
    identifier: s3bucketmetric-aws.kubeform.com
    name: S3BucketMetric
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## S3BucketMetric
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `S3BucketMetric` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[S3BucketMetricSpec](#s3bucketmetricspec)***||
| `status` | ***[S3BucketMetricStatus](#s3bucketmetricstatus)***||
## S3BucketMetricSpec

Appears on:[S3BucketMetric](#s3bucketmetric), [S3BucketMetricStatus](#s3bucketmetricstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `filter` | ***[[]S3BucketMetricSpecFilter](#s3bucketmetricspecfilter)***| ***(Optional)*** |
| `name` | ***string***||
## S3BucketMetricSpecFilter

Appears on:[S3BucketMetricSpec](#s3bucketmetricspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `prefix` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## S3BucketMetricStatus

Appears on:[S3BucketMetric](#s3bucketmetric)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[S3BucketMetricSpec](#s3bucketmetricspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
