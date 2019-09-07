---
title: ConfigDeliveryChannel
menu:
  docs_v0.0.1:
    identifier: configdeliverychannel-aws.kubeform.com
    name: ConfigDeliveryChannel
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ConfigDeliveryChannel
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ConfigDeliveryChannel` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ConfigDeliveryChannelSpec](#ConfigDeliveryChannelSpec)***||
| `status` | ***[ConfigDeliveryChannelStatus](#ConfigDeliveryChannelStatus)***||
## ConfigDeliveryChannelSpec
##### (Appears on:[ConfigDeliveryChannel](#ConfigDeliveryChannel), [ConfigDeliveryChannelStatus](#ConfigDeliveryChannelStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `s3BucketName` | ***string***||
| `s3KeyPrefix` | ***string***| ***(Optional)*** |
| `snapshotDeliveryProperties` | ***[[]ConfigDeliveryChannelSpecSnapshotDeliveryProperties](#ConfigDeliveryChannelSpecSnapshotDeliveryProperties)***| ***(Optional)*** |
| `snsTopicArn` | ***string***| ***(Optional)*** |
## ConfigDeliveryChannelSpecSnapshotDeliveryProperties
##### (Appears on:[ConfigDeliveryChannelSpec](#ConfigDeliveryChannelSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deliveryFrequency` | ***string***| ***(Optional)*** |
## ConfigDeliveryChannelStatus
##### (Appears on:[ConfigDeliveryChannel](#ConfigDeliveryChannel))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ConfigDeliveryChannelSpec](#ConfigDeliveryChannelSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
