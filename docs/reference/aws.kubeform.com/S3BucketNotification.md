---
title: S3BucketNotification
menu:
  docs_{{ .version }}:
    identifier: s3bucketnotification-aws.kubeform.com
    name: S3BucketNotification
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## S3BucketNotification
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `S3BucketNotification` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[S3BucketNotificationSpec](#s3bucketnotificationspec)***||
| `status` | ***[S3BucketNotificationStatus](#s3bucketnotificationstatus)***||
## S3BucketNotificationSpec

Appears on:[S3BucketNotification](#s3bucketnotification), [S3BucketNotificationStatus](#s3bucketnotificationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `lambdaFunction` | ***[[]S3BucketNotificationSpecLambdaFunction](#s3bucketnotificationspeclambdafunction)***| ***(Optional)*** |
| `queue` | ***[[]S3BucketNotificationSpecQueue](#s3bucketnotificationspecqueue)***| ***(Optional)*** |
| `topic` | ***[[]S3BucketNotificationSpecTopic](#s3bucketnotificationspectopic)***| ***(Optional)*** |
## S3BucketNotificationSpecLambdaFunction

Appears on:[S3BucketNotificationSpec](#s3bucketnotificationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `events` | ***[]string***||
| `filterPrefix` | ***string***| ***(Optional)*** |
| `filterSuffix` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `lambdaFunctionArn` | ***string***| ***(Optional)*** |
## S3BucketNotificationSpecQueue

Appears on:[S3BucketNotificationSpec](#s3bucketnotificationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `events` | ***[]string***||
| `filterPrefix` | ***string***| ***(Optional)*** |
| `filterSuffix` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `queueArn` | ***string***||
## S3BucketNotificationSpecTopic

Appears on:[S3BucketNotificationSpec](#s3bucketnotificationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `events` | ***[]string***||
| `filterPrefix` | ***string***| ***(Optional)*** |
| `filterSuffix` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `topicArn` | ***string***||
## S3BucketNotificationStatus

Appears on:[S3BucketNotification](#s3bucketnotification)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[S3BucketNotificationSpec](#s3bucketnotificationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
