---
title: ConfigDeliveryChannel
menu:
  docs_{{ .version }}:
    identifier: configdeliverychannel-aws.kubeform.com
    name: ConfigDeliveryChannel
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ConfigDeliveryChannel
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ConfigDeliveryChannel` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ConfigDeliveryChannelSpec](#configdeliverychannelspec)***||
| `status` | ***[ConfigDeliveryChannelStatus](#configdeliverychannelstatus)***||
## ConfigDeliveryChannelSpec

Appears on:[ConfigDeliveryChannel](#configdeliverychannel), [ConfigDeliveryChannelStatus](#configdeliverychannelstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `s3BucketName` | ***string***||
| `s3KeyPrefix` | ***string***| ***(Optional)*** |
| `snapshotDeliveryProperties` | ***[[]ConfigDeliveryChannelSpecSnapshotDeliveryProperties](#configdeliverychannelspecsnapshotdeliveryproperties)***| ***(Optional)*** |
| `snsTopicArn` | ***string***| ***(Optional)*** |
## ConfigDeliveryChannelSpecSnapshotDeliveryProperties

Appears on:[ConfigDeliveryChannelSpec](#configdeliverychannelspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deliveryFrequency` | ***string***| ***(Optional)*** |
## ConfigDeliveryChannelStatus

Appears on:[ConfigDeliveryChannel](#configdeliverychannel)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ConfigDeliveryChannelSpec](#configdeliverychannelspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ConfigDeliveryChannelStatus](#configdeliverychannelstatus)

---
