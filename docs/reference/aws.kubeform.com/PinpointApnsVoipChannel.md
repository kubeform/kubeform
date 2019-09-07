---
title: PinpointApnsVoipChannel
menu:
  docs_v0.0.1:
    identifier: pinpointapnsvoipchannel-aws.kubeform.com
    name: PinpointApnsVoipChannel
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## PinpointApnsVoipChannel
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `PinpointApnsVoipChannel` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PinpointApnsVoipChannelSpec](#PinpointApnsVoipChannelSpec)***||
| `status` | ***[PinpointApnsVoipChannelStatus](#PinpointApnsVoipChannelStatus)***||
## PinpointApnsVoipChannelSpec
##### (Appears on:[PinpointApnsVoipChannel](#PinpointApnsVoipChannel), [PinpointApnsVoipChannelStatus](#PinpointApnsVoipChannelStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `applicationID` | ***string***||
| `defaultAuthenticationMethod` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
## PinpointApnsVoipChannelStatus
##### (Appears on:[PinpointApnsVoipChannel](#PinpointApnsVoipChannel))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PinpointApnsVoipChannelSpec](#PinpointApnsVoipChannelSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `bundle_id` | ***string*** ||
| `certificate` | ***string*** ||
| `private_key` | ***string*** ||
| `team_id` | ***string*** ||
| `token_key` | ***string*** ||
| `token_key_id` | ***string*** ||
