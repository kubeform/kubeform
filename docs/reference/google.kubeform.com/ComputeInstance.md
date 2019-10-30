---
title: ComputeInstance
menu:
  docs_{{ .version }}:
    identifier: computeinstance-google.kubeform.com
    name: ComputeInstance
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeInstanceSpec](#computeinstancespec)***||
| `status` | ***[ComputeInstanceStatus](#computeinstancestatus)***||
## ComputeInstanceSpec

Appears on:[ComputeInstance](#computeinstance), [ComputeInstanceStatus](#computeinstancestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `allowStoppingForUpdate` | ***bool***| ***(Optional)*** |
| `attachedDisk` | ***[[]ComputeInstanceSpecAttachedDisk](#computeinstancespecattacheddisk)***| ***(Optional)*** |
| `bootDisk` | ***[[]ComputeInstanceSpecBootDisk](#computeinstancespecbootdisk)***||
| `canIPForward` | ***bool***| ***(Optional)*** |
| `cpuPlatform` | ***string***| ***(Optional)*** |
| `createTimeout` | ***int***| ***(Optional)*** Deprecated|
| `deletionProtection` | ***bool***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `guestAccelerator` | ***[[]ComputeInstanceSpecGuestAccelerator](#computeinstancespecguestaccelerator)***| ***(Optional)*** |
| `instanceID` | ***string***| ***(Optional)*** |
| `labelFingerprint` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `machineType` | ***string***||
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `metadataFingerprint` | ***string***| ***(Optional)*** |
| `metadataStartupScript` | ***string***| ***(Optional)*** |
| `minCPUPlatform` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `networkInterface` | ***[[]ComputeInstanceSpecNetworkInterface](#computeinstancespecnetworkinterface)***||
| `project` | ***string***| ***(Optional)*** |
| `scheduling` | ***[[]ComputeInstanceSpecScheduling](#computeinstancespecscheduling)***| ***(Optional)*** |
| `scratchDisk` | ***[[]ComputeInstanceSpecScratchDisk](#computeinstancespecscratchdisk)***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `serviceAccount` | ***[[]ComputeInstanceSpecServiceAccount](#computeinstancespecserviceaccount)***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
| `tagsFingerprint` | ***string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## ComputeInstanceSpecAttachedDisk

Appears on:[ComputeInstanceSpec](#computeinstancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deviceName` | ***string***| ***(Optional)*** |
| `diskEncryptionKeySha256` | ***string***| ***(Optional)*** |
| `mode` | ***string***| ***(Optional)*** |
| `source` | ***string***||
## ComputeInstanceSpecBootDisk

Appears on:[ComputeInstanceSpec](#computeinstancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoDelete` | ***bool***| ***(Optional)*** |
| `deviceName` | ***string***| ***(Optional)*** |
| `diskEncryptionKeySha256` | ***string***| ***(Optional)*** |
| `initializeParams` | ***[[]ComputeInstanceSpecBootDiskInitializeParams](#computeinstancespecbootdiskinitializeparams)***| ***(Optional)*** |
| `source` | ***string***| ***(Optional)*** |
## ComputeInstanceSpecBootDiskInitializeParams

Appears on:[ComputeInstanceSpecBootDisk](#computeinstancespecbootdisk)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `image` | ***string***| ***(Optional)*** |
| `size` | ***int***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## ComputeInstanceSpecGuestAccelerator

Appears on:[ComputeInstanceSpec](#computeinstancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***||
| `type` | ***string***||
## ComputeInstanceSpecNetworkInterface

Appears on:[ComputeInstanceSpec](#computeinstancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `accessConfig` | ***[[]ComputeInstanceSpecNetworkInterfaceAccessConfig](#computeinstancespecnetworkinterfaceaccessconfig)***| ***(Optional)*** |
| `address` | ***string***| ***(Optional)*** Deprecated|
| `aliasIPRange` | ***[[]ComputeInstanceSpecNetworkInterfaceAliasIPRange](#computeinstancespecnetworkinterfacealiasiprange)***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `network` | ***string***| ***(Optional)*** |
| `networkIP` | ***string***| ***(Optional)*** |
| `subnetwork` | ***string***| ***(Optional)*** |
| `subnetworkProject` | ***string***| ***(Optional)*** |
## ComputeInstanceSpecNetworkInterfaceAccessConfig

Appears on:[ComputeInstanceSpecNetworkInterface](#computeinstancespecnetworkinterface)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `assignedNATIP` | ***string***| ***(Optional)*** Deprecated|
| `natIP` | ***string***| ***(Optional)*** |
| `networkTier` | ***string***| ***(Optional)*** |
| `publicPtrDomainName` | ***string***| ***(Optional)*** |
## ComputeInstanceSpecNetworkInterfaceAliasIPRange

Appears on:[ComputeInstanceSpecNetworkInterface](#computeinstancespecnetworkinterface)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ipCIDRRange` | ***string***||
| `subnetworkRangeName` | ***string***| ***(Optional)*** |
## ComputeInstanceSpecScheduling

Appears on:[ComputeInstanceSpec](#computeinstancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `automaticRestart` | ***bool***| ***(Optional)*** |
| `onHostMaintenance` | ***string***| ***(Optional)*** |
| `preemptible` | ***bool***| ***(Optional)*** |
## ComputeInstanceSpecScratchDisk

Appears on:[ComputeInstanceSpec](#computeinstancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `interface` | ***string***| ***(Optional)*** |
## ComputeInstanceSpecServiceAccount

Appears on:[ComputeInstanceSpec](#computeinstancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `email` | ***string***| ***(Optional)*** |
| `scopes` | ***[]string***||
## ComputeInstanceStatus

Appears on:[ComputeInstance](#computeinstance)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeInstanceSpec](#computeinstancespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeInstanceStatus](#computeinstancestatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `attached_disk.<index>.disk_encryption_key_raw` | ***string*** ||
| `boot_disk.<index>.disk_encryption_key_raw` | ***string*** ||
