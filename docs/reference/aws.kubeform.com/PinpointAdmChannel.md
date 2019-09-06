## PinpointAdmChannel
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `PinpointAdmChannel` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PinpointAdmChannelSpec](#PinpointAdmChannelSpec)***||
| `status` | ***[PinpointAdmChannelStatus](#PinpointAdmChannelStatus)***||
## PinpointAdmChannelSpec
##### (Appears on:[PinpointAdmChannel](#PinpointAdmChannel), [PinpointAdmChannelStatus](#PinpointAdmChannelStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `applicationID` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
## PinpointAdmChannelStatus
##### (Appears on:[PinpointAdmChannel](#PinpointAdmChannel))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PinpointAdmChannelSpec](#PinpointAdmChannelSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `client_id` | ***string*** ||
| `client_secret` | ***string*** ||
