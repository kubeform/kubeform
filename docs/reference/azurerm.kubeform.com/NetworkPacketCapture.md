## NetworkPacketCapture
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkPacketCapture` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkPacketCaptureSpec](#NetworkPacketCaptureSpec)***||
| `status` | ***[NetworkPacketCaptureStatus](#NetworkPacketCaptureStatus)***||
## NetworkPacketCaptureSpec
##### (Appears on:[NetworkPacketCapture](#NetworkPacketCapture), [NetworkPacketCaptureStatus](#NetworkPacketCaptureStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `filter` | ***[[]NetworkPacketCaptureSpecFilter](#NetworkPacketCaptureSpecFilter)***| ***(Optional)*** |
| `maximumBytesPerPacket` | ***int***| ***(Optional)*** |
| `maximumBytesPerSession` | ***int***| ***(Optional)*** |
| `maximumCaptureDuration` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `networkWatcherName` | ***string***||
| `resourceGroupName` | ***string***||
| `storageLocation` | ***[[]NetworkPacketCaptureSpecStorageLocation](#NetworkPacketCaptureSpecStorageLocation)***||
| `targetResourceID` | ***string***||
## NetworkPacketCaptureSpecFilter
##### (Appears on:[NetworkPacketCaptureSpec](#NetworkPacketCaptureSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `localIPAddress` | ***string***| ***(Optional)*** |
| `localPort` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
| `remoteIPAddress` | ***string***| ***(Optional)*** |
| `remotePort` | ***string***| ***(Optional)*** |
## NetworkPacketCaptureSpecStorageLocation
##### (Appears on:[NetworkPacketCaptureSpec](#NetworkPacketCaptureSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `filePath` | ***string***| ***(Optional)*** |
| `storageAccountID` | ***string***| ***(Optional)*** |
| `storagePath` | ***string***| ***(Optional)*** |
## NetworkPacketCaptureStatus
##### (Appears on:[NetworkPacketCapture](#NetworkPacketCapture))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkPacketCaptureSpec](#NetworkPacketCaptureSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
