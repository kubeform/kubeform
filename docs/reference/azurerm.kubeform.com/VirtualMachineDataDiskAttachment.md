## VirtualMachineDataDiskAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `VirtualMachineDataDiskAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VirtualMachineDataDiskAttachmentSpec](#VirtualMachineDataDiskAttachmentSpec)***||
| `status` | ***[VirtualMachineDataDiskAttachmentStatus](#VirtualMachineDataDiskAttachmentStatus)***||
## VirtualMachineDataDiskAttachmentSpec
##### (Appears on:[VirtualMachineDataDiskAttachment](#VirtualMachineDataDiskAttachment), [VirtualMachineDataDiskAttachmentStatus](#VirtualMachineDataDiskAttachmentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `caching` | ***string***||
| `createOption` | ***string***| ***(Optional)*** |
| `lun` | ***int***||
| `managedDiskID` | ***string***||
| `virtualMachineID` | ***string***||
| `writeAcceleratorEnabled` | ***bool***| ***(Optional)*** |
## VirtualMachineDataDiskAttachmentStatus
##### (Appears on:[VirtualMachineDataDiskAttachment](#VirtualMachineDataDiskAttachment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VirtualMachineDataDiskAttachmentSpec](#VirtualMachineDataDiskAttachmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
