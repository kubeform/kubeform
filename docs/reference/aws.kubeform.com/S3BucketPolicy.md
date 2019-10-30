---
title: S3BucketPolicy
menu:
  docs_{{ .version }}:
    identifier: s3bucketpolicy-aws.kubeform.com
    name: S3BucketPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## S3BucketPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `S3BucketPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[S3BucketPolicySpec](#s3bucketpolicyspec)***||
| `status` | ***[S3BucketPolicyStatus](#s3bucketpolicystatus)***||
## Phase(`string` alias)

Appears on:[S3BucketPolicyStatus](#s3bucketpolicystatus)

## S3BucketPolicySpec

Appears on:[S3BucketPolicy](#s3bucketpolicy), [S3BucketPolicyStatus](#s3bucketpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `policy` | ***string***||
## S3BucketPolicyStatus

Appears on:[S3BucketPolicy](#s3bucketpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[S3BucketPolicySpec](#s3bucketpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
