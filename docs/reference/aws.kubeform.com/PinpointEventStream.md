---
title: PinpointEventStream
menu:
  docs_{{ .version }}:
    identifier: pinpointeventstream-aws.kubeform.com
    name: PinpointEventStream
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PinpointEventStream
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `PinpointEventStream` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PinpointEventStreamSpec](#pinpointeventstreamspec)***||
| `status` | ***[PinpointEventStreamStatus](#pinpointeventstreamstatus)***||
## Phase(`string` alias)

Appears on:[PinpointEventStreamStatus](#pinpointeventstreamstatus)

## PinpointEventStreamSpec

Appears on:[PinpointEventStream](#pinpointeventstream), [PinpointEventStreamStatus](#pinpointeventstreamstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `applicationID` | ***string***||
| `destinationStreamArn` | ***string***||
| `roleArn` | ***string***||
## PinpointEventStreamStatus

Appears on:[PinpointEventStream](#pinpointeventstream)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PinpointEventStreamSpec](#pinpointeventstreamspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
