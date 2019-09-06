## DevTestWindowsVirtualMachine
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DevTestWindowsVirtualMachine` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DevTestWindowsVirtualMachineSpec](#DevTestWindowsVirtualMachineSpec)***||
| `status` | ***[DevTestWindowsVirtualMachineStatus](#DevTestWindowsVirtualMachineStatus)***||
## DevTestWindowsVirtualMachineSpec
##### (Appears on:[DevTestWindowsVirtualMachine](#DevTestWindowsVirtualMachine), [DevTestWindowsVirtualMachineStatus](#DevTestWindowsVirtualMachineStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allowClaim` | ***bool***| ***(Optional)*** |
| `disallowPublicIPAddress` | ***bool***| ***(Optional)*** |
| `fqdn` | ***string***| ***(Optional)*** |
| `galleryImageReference` | ***[[]DevTestWindowsVirtualMachineSpecGalleryImageReference](#DevTestWindowsVirtualMachineSpecGalleryImageReference)***||
| `inboundNATRule` | ***[[]DevTestWindowsVirtualMachineSpecInboundNATRule](#DevTestWindowsVirtualMachineSpecInboundNATRule)***| ***(Optional)*** |
| `labName` | ***string***||
| `labSubnetName` | ***string***||
| `labVirtualNetworkID` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `notes` | ***string***| ***(Optional)*** |
| `password` | ***string***||
| `resourceGroupName` | ***string***||
| `size` | ***string***||
| `storageType` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `uniqueIdentifier` | ***string***| ***(Optional)*** |
| `username` | ***string***||
## DevTestWindowsVirtualMachineSpecGalleryImageReference
##### (Appears on:[DevTestWindowsVirtualMachineSpec](#DevTestWindowsVirtualMachineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `offer` | ***string***||
| `publisher` | ***string***||
| `sku` | ***string***||
| `version` | ***string***||
## DevTestWindowsVirtualMachineSpecInboundNATRule
##### (Appears on:[DevTestWindowsVirtualMachineSpec](#DevTestWindowsVirtualMachineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `backendPort` | ***int***||
| `frontendPort` | ***int***| ***(Optional)*** |
| `protocol` | ***string***||
## DevTestWindowsVirtualMachineStatus
##### (Appears on:[DevTestWindowsVirtualMachine](#DevTestWindowsVirtualMachine))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DevTestWindowsVirtualMachineSpec](#DevTestWindowsVirtualMachineSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
