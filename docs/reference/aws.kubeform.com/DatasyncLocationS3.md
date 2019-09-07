---
title: DatasyncLocationS3
menu:
  docs_v0.0.1:
    identifier: datasynclocations3-aws.kubeform.com
    name: DatasyncLocationS3
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DatasyncLocationS3
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DatasyncLocationS3` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DatasyncLocationS3Spec](#DatasyncLocationS3Spec)***||
| `status` | ***[DatasyncLocationS3Status](#DatasyncLocationS3Status)***||
## DatasyncLocationS3Spec
##### (Appears on:[DatasyncLocationS3](#DatasyncLocationS3), [DatasyncLocationS3Status](#DatasyncLocationS3Status))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `s3BucketArn` | ***string***||
| `s3Config` | ***[[]DatasyncLocationS3SpecS3Config](#DatasyncLocationS3SpecS3Config)***||
| `subdirectory` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `uri` | ***string***| ***(Optional)*** |
## DatasyncLocationS3SpecS3Config
##### (Appears on:[DatasyncLocationS3Spec](#DatasyncLocationS3Spec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketAccessRoleArn` | ***string***||
## DatasyncLocationS3Status
##### (Appears on:[DatasyncLocationS3](#DatasyncLocationS3))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DatasyncLocationS3Spec](#DatasyncLocationS3Spec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
