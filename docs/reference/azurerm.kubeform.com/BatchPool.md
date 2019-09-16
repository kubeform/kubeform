---
title: BatchPool
menu:
  docs_{{ .version }}:
    identifier: batchpool-azurerm.kubeform.com
    name: BatchPool
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BatchPool
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `BatchPool` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BatchPoolSpec](#batchpoolspec)***||
| `status` | ***[BatchPoolStatus](#batchpoolstatus)***||
## BatchPoolSpec

Appears on:[BatchPool](#batchpool), [BatchPoolStatus](#batchpoolstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***||
| `autoScale` | ***[[]BatchPoolSpecAutoScale](#batchpoolspecautoscale)***| ***(Optional)*** |
| `certificate` | ***[[]BatchPoolSpecCertificate](#batchpoolspeccertificate)***| ***(Optional)*** |
| `containerConfiguration` | ***[[]BatchPoolSpecContainerConfiguration](#batchpoolspeccontainerconfiguration)***| ***(Optional)*** |
| `displayName` | ***string***| ***(Optional)*** |
| `fixedScale` | ***[[]BatchPoolSpecFixedScale](#batchpoolspecfixedscale)***| ***(Optional)*** |
| `maxTasksPerNode` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `nodeAgentSkuID` | ***string***||
| `resourceGroupName` | ***string***||
| `startTask` | ***[[]BatchPoolSpecStartTask](#batchpoolspecstarttask)***| ***(Optional)*** |
| `stopPendingResizeOperation` | ***bool***| ***(Optional)*** |
| `storageImageReference` | ***[[]BatchPoolSpecStorageImageReference](#batchpoolspecstorageimagereference)***||
| `vmSize` | ***string***||
## BatchPoolSpecAutoScale

Appears on:[BatchPoolSpec](#batchpoolspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `evaluationInterval` | ***string***| ***(Optional)*** |
| `formula` | ***string***||
## BatchPoolSpecCertificate

Appears on:[BatchPoolSpec](#batchpoolspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***||
| `storeLocation` | ***string***||
| `storeName` | ***string***| ***(Optional)*** |
| `visibility` | ***[]string***| ***(Optional)*** |
## BatchPoolSpecContainerConfiguration

Appears on:[BatchPoolSpec](#batchpoolspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***| ***(Optional)*** |
## BatchPoolSpecFixedScale

Appears on:[BatchPoolSpec](#batchpoolspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `resizeTimeout` | ***string***| ***(Optional)*** |
| `targetDedicatedNodes` | ***int***| ***(Optional)*** |
| `targetLowPriorityNodes` | ***int***| ***(Optional)*** |
## BatchPoolSpecStartTask

Appears on:[BatchPoolSpec](#batchpoolspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `commandLine` | ***string***||
| `environment` | ***map[string]string***| ***(Optional)*** |
| `maxTaskRetryCount` | ***int***| ***(Optional)*** |
| `resourceFile` | ***[[]BatchPoolSpecStartTaskResourceFile](#batchpoolspecstarttaskresourcefile)***| ***(Optional)*** |
| `userIdentity` | ***[[]BatchPoolSpecStartTaskUserIdentity](#batchpoolspecstarttaskuseridentity)***||
| `waitForSuccess` | ***bool***| ***(Optional)*** |
## BatchPoolSpecStartTaskResourceFile

Appears on:[BatchPoolSpecStartTask](#batchpoolspecstarttask)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoStorageContainerName` | ***string***| ***(Optional)*** |
| `blobPrefix` | ***string***| ***(Optional)*** |
| `fileMode` | ***string***| ***(Optional)*** |
| `filePath` | ***string***| ***(Optional)*** |
| `httpURL` | ***string***| ***(Optional)*** |
| `storageContainerURL` | ***string***| ***(Optional)*** |
## BatchPoolSpecStartTaskUserIdentity

Appears on:[BatchPoolSpecStartTask](#batchpoolspecstarttask)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoUser` | ***[[]BatchPoolSpecStartTaskUserIdentityAutoUser](#batchpoolspecstarttaskuseridentityautouser)***| ***(Optional)*** |
| `userName` | ***string***| ***(Optional)*** |
## BatchPoolSpecStartTaskUserIdentityAutoUser

Appears on:[BatchPoolSpecStartTaskUserIdentity](#batchpoolspecstarttaskuseridentity)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `elevationLevel` | ***string***| ***(Optional)*** |
| `scope` | ***string***| ***(Optional)*** |
## BatchPoolSpecStorageImageReference

Appears on:[BatchPoolSpec](#batchpoolspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `offer` | ***string***| ***(Optional)*** |
| `publisher` | ***string***| ***(Optional)*** |
| `sku` | ***string***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
## BatchPoolStatus

Appears on:[BatchPool](#batchpool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BatchPoolSpec](#batchpoolspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
