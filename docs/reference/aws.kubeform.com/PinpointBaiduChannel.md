---
title: PinpointBaiduChannel
menu:
  docs_{{ .version }}:
    identifier: pinpointbaiduchannel-aws.kubeform.com
    name: PinpointBaiduChannel
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PinpointBaiduChannel
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `PinpointBaiduChannel` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PinpointBaiduChannelSpec](#PinpointBaiduChannelSpec)***||
| `status` | ***[PinpointBaiduChannelStatus](#PinpointBaiduChannelStatus)***||
## PinpointBaiduChannelSpec

Appears on:[PinpointBaiduChannel](#PinpointBaiduChannel), [PinpointBaiduChannelStatus](#PinpointBaiduChannelStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `applicationID` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
## PinpointBaiduChannelStatus

Appears on:[PinpointBaiduChannel](#PinpointBaiduChannel)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PinpointBaiduChannelSpec](#PinpointBaiduChannelSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `api_key` | ***string*** ||
| `secret_key` | ***string*** ||
