## PinpointEventStream
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `PinpointEventStream` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PinpointEventStreamSpec](#PinpointEventStreamSpec)***||
| `status` | ***[PinpointEventStreamStatus](#PinpointEventStreamStatus)***||
## PinpointEventStreamSpec
##### (Appears on:[PinpointEventStream](#PinpointEventStream), [PinpointEventStreamStatus](#PinpointEventStreamStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `applicationID` | ***string***||
| `destinationStreamArn` | ***string***||
| `roleArn` | ***string***||
## PinpointEventStreamStatus
##### (Appears on:[PinpointEventStream](#PinpointEventStream))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PinpointEventStreamSpec](#PinpointEventStreamSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
