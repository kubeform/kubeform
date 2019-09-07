---
title: ComputeInstance
menu:
  docs_v0.0.1:
    identifier: computeinstance-google.kubeform.com
    name: ComputeInstance
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ComputeInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeInstanceSpec](#ComputeInstanceSpec)***||
| `status` | ***[ComputeInstanceStatus](#ComputeInstanceStatus)***||
## ComputeInstanceSpec
##### (Appears on:[ComputeInstance](#ComputeInstance), [ComputeInstanceStatus](#ComputeInstanceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `allowStoppingForUpdate` | ***bool***| ***(Optional)*** |
| `attachedDisk` | ***[[]ComputeInstanceSpecAttachedDisk](#ComputeInstanceSpecAttachedDisk)***| ***(Optional)*** |
| `bootDisk` | ***[[]ComputeInstanceSpecBootDisk](#ComputeInstanceSpecBootDisk)***||
| `canIPForward` | ***bool***| ***(Optional)*** |
| `cpuPlatform` | ***string***| ***(Optional)*** |
| `createTimeout` | ***int***| ***(Optional)*** Deprecated|
| `deletionProtection` | ***bool***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `guestAccelerator` | ***[[]ComputeInstanceSpecGuestAccelerator](#ComputeInstanceSpecGuestAccelerator)***| ***(Optional)*** |
| `instanceID` | ***string***| ***(Optional)*** |
| `labelFingerprint` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `machineType` | ***string***||
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `metadataFingerprint` | ***string***| ***(Optional)*** |
| `metadataStartupScript` | ***string***| ***(Optional)*** |
| `minCPUPlatform` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `networkInterface` | ***[[]ComputeInstanceSpecNetworkInterface](#ComputeInstanceSpecNetworkInterface)***||
| `project` | ***string***| ***(Optional)*** |
| `scheduling` | ***[[]ComputeInstanceSpecScheduling](#ComputeInstanceSpecScheduling)***| ***(Optional)*** |
| `scratchDisk` | ***[[]ComputeInstanceSpecScratchDisk](#ComputeInstanceSpecScratchDisk)***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `serviceAccount` | ***[[]ComputeInstanceSpecServiceAccount](#ComputeInstanceSpecServiceAccount)***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
| `tagsFingerprint` | ***string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## ComputeInstanceSpecAttachedDisk
##### (Appears on:[ComputeInstanceSpec](#ComputeInstanceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deviceName` | ***string***| ***(Optional)*** |
| `diskEncryptionKeySha256` | ***string***| ***(Optional)*** |
| `mode` | ***string***| ***(Optional)*** |
| `source` | ***string***||
## ComputeInstanceSpecBootDisk
##### (Appears on:[ComputeInstanceSpec](#ComputeInstanceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoDelete` | ***bool***| ***(Optional)*** |
| `deviceName` | ***string***| ***(Optional)*** |
| `diskEncryptionKeySha256` | ***string***| ***(Optional)*** |
| `initializeParams` | ***[[]ComputeInstanceSpecBootDiskInitializeParams](#ComputeInstanceSpecBootDiskInitializeParams)***| ***(Optional)*** |
| `source` | ***string***| ***(Optional)*** |
## ComputeInstanceSpecBootDiskInitializeParams
##### (Appears on:[ComputeInstanceSpecBootDisk](#ComputeInstanceSpecBootDisk))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `image` | ***string***| ***(Optional)*** |
| `size` | ***int***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## ComputeInstanceSpecGuestAccelerator
##### (Appears on:[ComputeInstanceSpec](#ComputeInstanceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***||
| `type` | ***string***||
## ComputeInstanceSpecNetworkInterface
##### (Appears on:[ComputeInstanceSpec](#ComputeInstanceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `accessConfig` | ***[[]ComputeInstanceSpecNetworkInterfaceAccessConfig](#ComputeInstanceSpecNetworkInterfaceAccessConfig)***| ***(Optional)*** |
| `address` | ***string***| ***(Optional)*** Deprecated|
| `aliasIPRange` | ***[[]ComputeInstanceSpecNetworkInterfaceAliasIPRange](#ComputeInstanceSpecNetworkInterfaceAliasIPRange)***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `network` | ***string***| ***(Optional)*** |
| `networkIP` | ***string***| ***(Optional)*** |
| `subnetwork` | ***string***| ***(Optional)*** |
| `subnetworkProject` | ***string***| ***(Optional)*** |
## ComputeInstanceSpecNetworkInterfaceAccessConfig
##### (Appears on:[ComputeInstanceSpecNetworkInterface](#ComputeInstanceSpecNetworkInterface))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `assignedNATIP` | ***string***| ***(Optional)*** Deprecated|
| `natIP` | ***string***| ***(Optional)*** |
| `networkTier` | ***string***| ***(Optional)*** |
| `publicPtrDomainName` | ***string***| ***(Optional)*** |
## ComputeInstanceSpecNetworkInterfaceAliasIPRange
##### (Appears on:[ComputeInstanceSpecNetworkInterface](#ComputeInstanceSpecNetworkInterface))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ipCIDRRange` | ***string***||
| `subnetworkRangeName` | ***string***| ***(Optional)*** |
## ComputeInstanceSpecScheduling
##### (Appears on:[ComputeInstanceSpec](#ComputeInstanceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `automaticRestart` | ***bool***| ***(Optional)*** |
| `onHostMaintenance` | ***string***| ***(Optional)*** |
| `preemptible` | ***bool***| ***(Optional)*** |
## ComputeInstanceSpecScratchDisk
##### (Appears on:[ComputeInstanceSpec](#ComputeInstanceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `interface` | ***string***| ***(Optional)*** |
## ComputeInstanceSpecServiceAccount
##### (Appears on:[ComputeInstanceSpec](#ComputeInstanceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `email` | ***string***| ***(Optional)*** |
| `scopes` | ***[]string***||
## ComputeInstanceStatus
##### (Appears on:[ComputeInstance](#ComputeInstance))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeInstanceSpec](#ComputeInstanceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `attached_disk.<index>.disk_encryption_key_raw` | ***string*** ||
| `boot_disk.<index>.disk_encryption_key_raw` | ***string*** ||
