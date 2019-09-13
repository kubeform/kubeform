---
title: PinpointGcmChannel
menu:
  docs_{{ .version }}:
    identifier: pinpointgcmchannel-aws.kubeform.com
    name: PinpointGcmChannel
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PinpointGcmChannel
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `PinpointGcmChannel` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PinpointGcmChannelSpec](#PinpointGcmChannelSpec)***||
| `status` | ***[PinpointGcmChannelStatus](#PinpointGcmChannelStatus)***||
## PinpointGcmChannelSpec
##### (Appears on:[PinpointGcmChannel](#PinpointGcmChannel), [PinpointGcmChannelStatus](#PinpointGcmChannelStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `applicationID` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
## PinpointGcmChannelStatus
##### (Appears on:[PinpointGcmChannel](#PinpointGcmChannel))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PinpointGcmChannelSpec](#PinpointGcmChannelSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `api_key` | ***string*** ||
