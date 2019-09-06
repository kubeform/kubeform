## ComputeInstanceFromTemplate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeInstanceFromTemplate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeInstanceFromTemplateSpec](#ComputeInstanceFromTemplateSpec)***||
| `status` | ***[ComputeInstanceFromTemplateStatus](#ComputeInstanceFromTemplateStatus)***||
## ComputeInstanceFromTemplateSpec
##### (Appears on:[ComputeInstanceFromTemplate](#ComputeInstanceFromTemplate), [ComputeInstanceFromTemplateStatus](#ComputeInstanceFromTemplateStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `allowStoppingForUpdate` | ***bool***| ***(Optional)*** |
| `attachedDisk` | ***[[]ComputeInstanceFromTemplateSpecAttachedDisk](#ComputeInstanceFromTemplateSpecAttachedDisk)***| ***(Optional)*** |
| `bootDisk` | ***[[]ComputeInstanceFromTemplateSpecBootDisk](#ComputeInstanceFromTemplateSpecBootDisk)***| ***(Optional)*** |
| `canIPForward` | ***bool***| ***(Optional)*** |
| `cpuPlatform` | ***string***| ***(Optional)*** |
| `deletionProtection` | ***bool***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `guestAccelerator` | ***[[]ComputeInstanceFromTemplateSpecGuestAccelerator](#ComputeInstanceFromTemplateSpecGuestAccelerator)***| ***(Optional)*** |
| `instanceID` | ***string***| ***(Optional)*** |
| `labelFingerprint` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `machineType` | ***string***| ***(Optional)*** |
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `metadataFingerprint` | ***string***| ***(Optional)*** |
| `metadataStartupScript` | ***string***| ***(Optional)*** |
| `minCPUPlatform` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `networkInterface` | ***[[]ComputeInstanceFromTemplateSpecNetworkInterface](#ComputeInstanceFromTemplateSpecNetworkInterface)***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `scheduling` | ***[[]ComputeInstanceFromTemplateSpecScheduling](#ComputeInstanceFromTemplateSpecScheduling)***| ***(Optional)*** |
| `scratchDisk` | ***[[]ComputeInstanceFromTemplateSpecScratchDisk](#ComputeInstanceFromTemplateSpecScratchDisk)***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `serviceAccount` | ***[[]ComputeInstanceFromTemplateSpecServiceAccount](#ComputeInstanceFromTemplateSpecServiceAccount)***| ***(Optional)*** |
| `sourceInstanceTemplate` | ***string***||
| `tags` | ***[]string***| ***(Optional)*** |
| `tagsFingerprint` | ***string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecAttachedDisk
##### (Appears on:[ComputeInstanceFromTemplateSpec](#ComputeInstanceFromTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deviceName` | ***string***| ***(Optional)*** |
| `diskEncryptionKeySha256` | ***string***| ***(Optional)*** |
| `mode` | ***string***| ***(Optional)*** |
| `source` | ***string***||
## ComputeInstanceFromTemplateSpecBootDisk
##### (Appears on:[ComputeInstanceFromTemplateSpec](#ComputeInstanceFromTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoDelete` | ***bool***| ***(Optional)*** |
| `deviceName` | ***string***| ***(Optional)*** |
| `diskEncryptionKeySha256` | ***string***| ***(Optional)*** |
| `initializeParams` | ***[[]ComputeInstanceFromTemplateSpecBootDiskInitializeParams](#ComputeInstanceFromTemplateSpecBootDiskInitializeParams)***| ***(Optional)*** |
| `source` | ***string***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecBootDiskInitializeParams
##### (Appears on:[ComputeInstanceFromTemplateSpecBootDisk](#ComputeInstanceFromTemplateSpecBootDisk))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `image` | ***string***| ***(Optional)*** |
| `size` | ***int***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecGuestAccelerator
##### (Appears on:[ComputeInstanceFromTemplateSpec](#ComputeInstanceFromTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***||
| `type` | ***string***||
## ComputeInstanceFromTemplateSpecNetworkInterface
##### (Appears on:[ComputeInstanceFromTemplateSpec](#ComputeInstanceFromTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `accessConfig` | ***[[]ComputeInstanceFromTemplateSpecNetworkInterfaceAccessConfig](#ComputeInstanceFromTemplateSpecNetworkInterfaceAccessConfig)***| ***(Optional)*** |
| `address` | ***string***| ***(Optional)*** Deprecated|
| `aliasIPRange` | ***[[]ComputeInstanceFromTemplateSpecNetworkInterfaceAliasIPRange](#ComputeInstanceFromTemplateSpecNetworkInterfaceAliasIPRange)***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `network` | ***string***| ***(Optional)*** |
| `networkIP` | ***string***| ***(Optional)*** |
| `subnetwork` | ***string***| ***(Optional)*** |
| `subnetworkProject` | ***string***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecNetworkInterfaceAccessConfig
##### (Appears on:[ComputeInstanceFromTemplateSpecNetworkInterface](#ComputeInstanceFromTemplateSpecNetworkInterface))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `assignedNATIP` | ***string***| ***(Optional)*** Deprecated|
| `natIP` | ***string***| ***(Optional)*** |
| `networkTier` | ***string***| ***(Optional)*** |
| `publicPtrDomainName` | ***string***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecNetworkInterfaceAliasIPRange
##### (Appears on:[ComputeInstanceFromTemplateSpecNetworkInterface](#ComputeInstanceFromTemplateSpecNetworkInterface))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ipCIDRRange` | ***string***||
| `subnetworkRangeName` | ***string***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecScheduling
##### (Appears on:[ComputeInstanceFromTemplateSpec](#ComputeInstanceFromTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `automaticRestart` | ***bool***| ***(Optional)*** |
| `onHostMaintenance` | ***string***| ***(Optional)*** |
| `preemptible` | ***bool***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecScratchDisk
##### (Appears on:[ComputeInstanceFromTemplateSpec](#ComputeInstanceFromTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `interface` | ***string***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecServiceAccount
##### (Appears on:[ComputeInstanceFromTemplateSpec](#ComputeInstanceFromTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `email` | ***string***| ***(Optional)*** |
| `scopes` | ***[]string***||
## ComputeInstanceFromTemplateStatus
##### (Appears on:[ComputeInstanceFromTemplate](#ComputeInstanceFromTemplate))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeInstanceFromTemplateSpec](#ComputeInstanceFromTemplateSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `attached_disk.<index>.disk_encryption_key_raw` | ***string*** ||
| `boot_disk.<index>.disk_encryption_key_raw` | ***string*** ||
