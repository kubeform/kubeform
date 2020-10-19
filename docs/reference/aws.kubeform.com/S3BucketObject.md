---
title: S3BucketObject
menu:
  docs_{{ .version }}:
    identifier: s3bucketobject-aws.kubeform.com
    name: S3BucketObject
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## S3BucketObject
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `S3BucketObject` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[S3BucketObjectSpec](#s3bucketobjectspec)***||
| `status` | ***[S3BucketObjectStatus](#s3bucketobjectstatus)***||
## Phase(`string` alias)

Appears on:[S3BucketObjectStatus](#s3bucketobjectstatus)

## S3BucketObjectSpec

Appears on:[S3BucketObject](#s3bucketobject), [S3BucketObjectStatus](#s3bucketobjectstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `acl` | ***string***| ***(Optional)*** |
| `bucket` | ***string***||
| `cacheControl` | ***string***| ***(Optional)*** |
| `content` | ***string***| ***(Optional)*** |
| `contentBase64` | ***string***| ***(Optional)*** |
| `contentDisposition` | ***string***| ***(Optional)*** |
| `contentEncoding` | ***string***| ***(Optional)*** |
| `contentLanguage` | ***string***| ***(Optional)*** |
| `contentType` | ***string***| ***(Optional)*** |
| `etag` | ***string***| ***(Optional)*** |
| `forceDestroy` | ***bool***| ***(Optional)*** |
| `key` | ***string***||
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `objectLockLegalHoldStatus` | ***string***| ***(Optional)*** |
| `objectLockMode` | ***string***| ***(Optional)*** |
| `objectLockRetainUntilDate` | ***string***| ***(Optional)*** |
| `serverSideEncryption` | ***string***| ***(Optional)*** |
| `source` | ***string***| ***(Optional)*** |
| `storageClass` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `versionID` | ***string***| ***(Optional)*** |
| `websiteRedirect` | ***string***| ***(Optional)*** |
## S3BucketObjectStatus

Appears on:[S3BucketObject](#s3bucketobject)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[S3BucketObjectSpec](#s3bucketobjectspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
