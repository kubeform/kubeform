---
title: SiteRecoveryReplicatedVm
menu:
  docs_{{ .version }}:
    identifier: siterecoveryreplicatedvm-azurerm.kubeform.com
    name: SiteRecoveryReplicatedVm
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SiteRecoveryReplicatedVm
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SiteRecoveryReplicatedVm` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SiteRecoveryReplicatedVmSpec](#siterecoveryreplicatedvmspec)***||
| `status` | ***[SiteRecoveryReplicatedVmStatus](#siterecoveryreplicatedvmstatus)***||
## Phase(`string` alias)

Appears on:[SiteRecoveryReplicatedVmStatus](#siterecoveryreplicatedvmstatus)

## SiteRecoveryReplicatedVmSpec

Appears on:[SiteRecoveryReplicatedVm](#siterecoveryreplicatedvm), [SiteRecoveryReplicatedVmStatus](#siterecoveryreplicatedvmstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `managedDisk` | ***[[]SiteRecoveryReplicatedVmSpecManagedDisk](#siterecoveryreplicatedvmspecmanageddisk)***| ***(Optional)*** |
| `name` | ***string***||
| `recoveryReplicationPolicyID` | ***string***||
| `recoveryVaultName` | ***string***||
| `resourceGroupName` | ***string***||
| `sourceRecoveryFabricName` | ***string***||
| `sourceRecoveryProtectionContainerName` | ***string***||
| `sourceVmID` | ***string***||
| `targetAvailabilitySetID` | ***string***| ***(Optional)*** |
| `targetRecoveryFabricID` | ***string***||
| `targetRecoveryProtectionContainerID` | ***string***||
| `targetResourceGroupID` | ***string***||
## SiteRecoveryReplicatedVmSpecManagedDisk

Appears on:[SiteRecoveryReplicatedVmSpec](#siterecoveryreplicatedvmspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskID` | ***string***||
| `stagingStorageAccountID` | ***string***||
| `targetDiskType` | ***string***||
| `targetReplicaDiskType` | ***string***||
| `targetResourceGroupID` | ***string***||
## SiteRecoveryReplicatedVmStatus

Appears on:[SiteRecoveryReplicatedVm](#siterecoveryreplicatedvm)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SiteRecoveryReplicatedVmSpec](#siterecoveryreplicatedvmspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
