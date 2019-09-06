## BatchPool
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `BatchPool` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BatchPoolSpec](#BatchPoolSpec)***||
| `status` | ***[BatchPoolStatus](#BatchPoolStatus)***||
## BatchPoolSpec
##### (Appears on:[BatchPool](#BatchPool), [BatchPoolStatus](#BatchPoolStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***||
| `autoScale` | ***[[]BatchPoolSpecAutoScale](#BatchPoolSpecAutoScale)***| ***(Optional)*** |
| `certificate` | ***[[]BatchPoolSpecCertificate](#BatchPoolSpecCertificate)***| ***(Optional)*** |
| `containerConfiguration` | ***[[]BatchPoolSpecContainerConfiguration](#BatchPoolSpecContainerConfiguration)***| ***(Optional)*** |
| `displayName` | ***string***| ***(Optional)*** |
| `fixedScale` | ***[[]BatchPoolSpecFixedScale](#BatchPoolSpecFixedScale)***| ***(Optional)*** |
| `maxTasksPerNode` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `nodeAgentSkuID` | ***string***||
| `resourceGroupName` | ***string***||
| `startTask` | ***[[]BatchPoolSpecStartTask](#BatchPoolSpecStartTask)***| ***(Optional)*** |
| `stopPendingResizeOperation` | ***bool***| ***(Optional)*** |
| `storageImageReference` | ***[[]BatchPoolSpecStorageImageReference](#BatchPoolSpecStorageImageReference)***||
| `vmSize` | ***string***||
## BatchPoolSpecAutoScale
##### (Appears on:[BatchPoolSpec](#BatchPoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `evaluationInterval` | ***string***| ***(Optional)*** |
| `formula` | ***string***||
## BatchPoolSpecCertificate
##### (Appears on:[BatchPoolSpec](#BatchPoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***||
| `storeLocation` | ***string***||
| `storeName` | ***string***| ***(Optional)*** |
| `visibility` | ***[]string***| ***(Optional)*** |
## BatchPoolSpecContainerConfiguration
##### (Appears on:[BatchPoolSpec](#BatchPoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***| ***(Optional)*** |
## BatchPoolSpecFixedScale
##### (Appears on:[BatchPoolSpec](#BatchPoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `resizeTimeout` | ***string***| ***(Optional)*** |
| `targetDedicatedNodes` | ***int***| ***(Optional)*** |
| `targetLowPriorityNodes` | ***int***| ***(Optional)*** |
## BatchPoolSpecStartTask
##### (Appears on:[BatchPoolSpec](#BatchPoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `commandLine` | ***string***||
| `environment` | ***map[string]string***| ***(Optional)*** |
| `maxTaskRetryCount` | ***int***| ***(Optional)*** |
| `resourceFile` | ***[[]BatchPoolSpecStartTaskResourceFile](#BatchPoolSpecStartTaskResourceFile)***| ***(Optional)*** |
| `userIdentity` | ***[[]BatchPoolSpecStartTaskUserIdentity](#BatchPoolSpecStartTaskUserIdentity)***||
| `waitForSuccess` | ***bool***| ***(Optional)*** |
## BatchPoolSpecStartTaskResourceFile
##### (Appears on:[BatchPoolSpecStartTask](#BatchPoolSpecStartTask))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoStorageContainerName` | ***string***| ***(Optional)*** |
| `blobPrefix` | ***string***| ***(Optional)*** |
| `fileMode` | ***string***| ***(Optional)*** |
| `filePath` | ***string***| ***(Optional)*** |
| `httpURL` | ***string***| ***(Optional)*** |
| `storageContainerURL` | ***string***| ***(Optional)*** |
## BatchPoolSpecStartTaskUserIdentity
##### (Appears on:[BatchPoolSpecStartTask](#BatchPoolSpecStartTask))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoUser` | ***[[]BatchPoolSpecStartTaskUserIdentityAutoUser](#BatchPoolSpecStartTaskUserIdentityAutoUser)***| ***(Optional)*** |
| `userName` | ***string***| ***(Optional)*** |
## BatchPoolSpecStartTaskUserIdentityAutoUser
##### (Appears on:[BatchPoolSpecStartTaskUserIdentity](#BatchPoolSpecStartTaskUserIdentity))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `elevationLevel` | ***string***| ***(Optional)*** |
| `scope` | ***string***| ***(Optional)*** |
## BatchPoolSpecStorageImageReference
##### (Appears on:[BatchPoolSpec](#BatchPoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `offer` | ***string***| ***(Optional)*** |
| `publisher` | ***string***| ***(Optional)*** |
| `sku` | ***string***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
## BatchPoolStatus
##### (Appears on:[BatchPool](#BatchPool))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BatchPoolSpec](#BatchPoolSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
