## SharedImageVersion
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SharedImageVersion` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SharedImageVersionSpec](#SharedImageVersionSpec)***||
| `status` | ***[SharedImageVersionStatus](#SharedImageVersionStatus)***||
## SharedImageVersionSpec
##### (Appears on:[SharedImageVersion](#SharedImageVersion), [SharedImageVersionStatus](#SharedImageVersionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `excludeFromLatest` | ***bool***| ***(Optional)*** |
| `galleryName` | ***string***||
| `imageName` | ***string***||
| `location` | ***string***||
| `managedImageID` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `targetRegion` | ***[[]SharedImageVersionSpecTargetRegion](#SharedImageVersionSpecTargetRegion)***||
## SharedImageVersionSpecTargetRegion
##### (Appears on:[SharedImageVersionSpec](#SharedImageVersionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `regionalReplicaCount` | ***int***||
## SharedImageVersionStatus
##### (Appears on:[SharedImageVersion](#SharedImageVersion))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SharedImageVersionSpec](#SharedImageVersionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
