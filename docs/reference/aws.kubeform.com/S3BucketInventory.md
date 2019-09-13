---
title: S3BucketInventory
menu:
  docs_{{ .version }}:
    identifier: s3bucketinventory-aws.kubeform.com
    name: S3BucketInventory
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## S3BucketInventory
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `S3BucketInventory` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[S3BucketInventorySpec](#S3BucketInventorySpec)***||
| `status` | ***[S3BucketInventoryStatus](#S3BucketInventoryStatus)***||
## S3BucketInventorySpec
##### (Appears on:[S3BucketInventory](#S3BucketInventory), [S3BucketInventoryStatus](#S3BucketInventoryStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `destination` | ***[[]S3BucketInventorySpecDestination](#S3BucketInventorySpecDestination)***||
| `enabled` | ***bool***| ***(Optional)*** |
| `filter` | ***[[]S3BucketInventorySpecFilter](#S3BucketInventorySpecFilter)***| ***(Optional)*** |
| `includedObjectVersions` | ***string***||
| `name` | ***string***||
| `optionalFields` | ***[]string***| ***(Optional)*** |
| `schedule` | ***[[]S3BucketInventorySpecSchedule](#S3BucketInventorySpecSchedule)***||
## S3BucketInventorySpecDestination
##### (Appears on:[S3BucketInventorySpec](#S3BucketInventorySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucket` | ***[[]S3BucketInventorySpecDestinationBucket](#S3BucketInventorySpecDestinationBucket)***||
## S3BucketInventorySpecDestinationBucket
##### (Appears on:[S3BucketInventorySpecDestination](#S3BucketInventorySpecDestination))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `accountID` | ***string***| ***(Optional)*** |
| `bucketArn` | ***string***||
| `encryption` | ***[[]S3BucketInventorySpecDestinationBucketEncryption](#S3BucketInventorySpecDestinationBucketEncryption)***| ***(Optional)*** |
| `format` | ***string***||
| `prefix` | ***string***| ***(Optional)*** |
## S3BucketInventorySpecDestinationBucketEncryption
##### (Appears on:[S3BucketInventorySpecDestinationBucket](#S3BucketInventorySpecDestinationBucket))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `sseKms` | ***[[]S3BucketInventorySpecDestinationBucketEncryptionSseKms](#S3BucketInventorySpecDestinationBucketEncryptionSseKms)***| ***(Optional)*** |
| `sseS3` | ***[[]S3BucketInventorySpecDestinationBucketEncryptionSseS3](#S3BucketInventorySpecDestinationBucketEncryptionSseS3)***| ***(Optional)*** |
## S3BucketInventorySpecDestinationBucketEncryptionSseKms
##### (Appears on:[S3BucketInventorySpecDestinationBucketEncryption](#S3BucketInventorySpecDestinationBucketEncryption))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `keyID` | ***string***||
## S3BucketInventorySpecDestinationBucketEncryptionSseS3
##### (Appears on:[S3BucketInventorySpecDestinationBucketEncryption](#S3BucketInventorySpecDestinationBucketEncryption))
## S3BucketInventorySpecFilter
##### (Appears on:[S3BucketInventorySpec](#S3BucketInventorySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `prefix` | ***string***| ***(Optional)*** |
## S3BucketInventorySpecSchedule
##### (Appears on:[S3BucketInventorySpec](#S3BucketInventorySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `frequency` | ***string***||
## S3BucketInventoryStatus
##### (Appears on:[S3BucketInventory](#S3BucketInventory))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[S3BucketInventorySpec](#S3BucketInventorySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
