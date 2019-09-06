## DevTestVirtualNetwork
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DevTestVirtualNetwork` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DevTestVirtualNetworkSpec](#DevTestVirtualNetworkSpec)***||
| `status` | ***[DevTestVirtualNetworkStatus](#DevTestVirtualNetworkStatus)***||
## DevTestVirtualNetworkSpec
##### (Appears on:[DevTestVirtualNetwork](#DevTestVirtualNetwork), [DevTestVirtualNetworkStatus](#DevTestVirtualNetworkStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `labName` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `subnet` | ***[[]DevTestVirtualNetworkSpecSubnet](#DevTestVirtualNetworkSpecSubnet)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `uniqueIdentifier` | ***string***| ***(Optional)*** |
## DevTestVirtualNetworkSpecSubnet
##### (Appears on:[DevTestVirtualNetworkSpec](#DevTestVirtualNetworkSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
| `useInVirtualMachineCreation` | ***string***| ***(Optional)*** |
| `usePublicIPAddress` | ***string***| ***(Optional)*** |
## DevTestVirtualNetworkStatus
##### (Appears on:[DevTestVirtualNetwork](#DevTestVirtualNetwork))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DevTestVirtualNetworkSpec](#DevTestVirtualNetworkSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
