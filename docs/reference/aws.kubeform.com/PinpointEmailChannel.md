---
title: PinpointEmailChannel
menu:
  docs_{{ .version }}:
    identifier: pinpointemailchannel-aws.kubeform.com
    name: PinpointEmailChannel
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PinpointEmailChannel
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `PinpointEmailChannel` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PinpointEmailChannelSpec](#pinpointemailchannelspec)***||
| `status` | ***[PinpointEmailChannelStatus](#pinpointemailchannelstatus)***||
## PinpointEmailChannelSpec

Appears on:[PinpointEmailChannel](#pinpointemailchannel), [PinpointEmailChannelStatus](#pinpointemailchannelstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `applicationID` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `fromAddress` | ***string***||
| `identity` | ***string***||
| `messagesPerSecond` | ***int***| ***(Optional)*** |
| `roleArn` | ***string***||
## PinpointEmailChannelStatus

Appears on:[PinpointEmailChannel](#pinpointemailchannel)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PinpointEmailChannelSpec](#pinpointemailchannelspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
