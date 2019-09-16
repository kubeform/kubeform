---
title: ComputeInstanceFromTemplate
menu:
  docs_{{ .version }}:
    identifier: computeinstancefromtemplate-google.kubeform.com
    name: ComputeInstanceFromTemplate
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeInstanceFromTemplate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeInstanceFromTemplate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeInstanceFromTemplateSpec](#computeinstancefromtemplatespec)***||
| `status` | ***[ComputeInstanceFromTemplateStatus](#computeinstancefromtemplatestatus)***||
## ComputeInstanceFromTemplateSpec

Appears on:[ComputeInstanceFromTemplate](#computeinstancefromtemplate), [ComputeInstanceFromTemplateStatus](#computeinstancefromtemplatestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `allowStoppingForUpdate` | ***bool***| ***(Optional)*** |
| `attachedDisk` | ***[[]ComputeInstanceFromTemplateSpecAttachedDisk](#computeinstancefromtemplatespecattacheddisk)***| ***(Optional)*** |
| `bootDisk` | ***[[]ComputeInstanceFromTemplateSpecBootDisk](#computeinstancefromtemplatespecbootdisk)***| ***(Optional)*** |
| `canIPForward` | ***bool***| ***(Optional)*** |
| `cpuPlatform` | ***string***| ***(Optional)*** |
| `deletionProtection` | ***bool***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `guestAccelerator` | ***[[]ComputeInstanceFromTemplateSpecGuestAccelerator](#computeinstancefromtemplatespecguestaccelerator)***| ***(Optional)*** |
| `instanceID` | ***string***| ***(Optional)*** |
| `labelFingerprint` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `machineType` | ***string***| ***(Optional)*** |
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `metadataFingerprint` | ***string***| ***(Optional)*** |
| `metadataStartupScript` | ***string***| ***(Optional)*** |
| `minCPUPlatform` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `networkInterface` | ***[[]ComputeInstanceFromTemplateSpecNetworkInterface](#computeinstancefromtemplatespecnetworkinterface)***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `scheduling` | ***[[]ComputeInstanceFromTemplateSpecScheduling](#computeinstancefromtemplatespecscheduling)***| ***(Optional)*** |
| `scratchDisk` | ***[[]ComputeInstanceFromTemplateSpecScratchDisk](#computeinstancefromtemplatespecscratchdisk)***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `serviceAccount` | ***[[]ComputeInstanceFromTemplateSpecServiceAccount](#computeinstancefromtemplatespecserviceaccount)***| ***(Optional)*** |
| `sourceInstanceTemplate` | ***string***||
| `tags` | ***[]string***| ***(Optional)*** |
| `tagsFingerprint` | ***string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecAttachedDisk

Appears on:[ComputeInstanceFromTemplateSpec](#computeinstancefromtemplatespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deviceName` | ***string***| ***(Optional)*** |
| `diskEncryptionKeySha256` | ***string***| ***(Optional)*** |
| `mode` | ***string***| ***(Optional)*** |
| `source` | ***string***||
## ComputeInstanceFromTemplateSpecBootDisk

Appears on:[ComputeInstanceFromTemplateSpec](#computeinstancefromtemplatespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoDelete` | ***bool***| ***(Optional)*** |
| `deviceName` | ***string***| ***(Optional)*** |
| `diskEncryptionKeySha256` | ***string***| ***(Optional)*** |
| `initializeParams` | ***[[]ComputeInstanceFromTemplateSpecBootDiskInitializeParams](#computeinstancefromtemplatespecbootdiskinitializeparams)***| ***(Optional)*** |
| `source` | ***string***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecBootDiskInitializeParams

Appears on:[ComputeInstanceFromTemplateSpecBootDisk](#computeinstancefromtemplatespecbootdisk)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `image` | ***string***| ***(Optional)*** |
| `size` | ***int***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecGuestAccelerator

Appears on:[ComputeInstanceFromTemplateSpec](#computeinstancefromtemplatespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***||
| `type` | ***string***||
## ComputeInstanceFromTemplateSpecNetworkInterface

Appears on:[ComputeInstanceFromTemplateSpec](#computeinstancefromtemplatespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `accessConfig` | ***[[]ComputeInstanceFromTemplateSpecNetworkInterfaceAccessConfig](#computeinstancefromtemplatespecnetworkinterfaceaccessconfig)***| ***(Optional)*** |
| `address` | ***string***| ***(Optional)*** Deprecated|
| `aliasIPRange` | ***[[]ComputeInstanceFromTemplateSpecNetworkInterfaceAliasIPRange](#computeinstancefromtemplatespecnetworkinterfacealiasiprange)***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `network` | ***string***| ***(Optional)*** |
| `networkIP` | ***string***| ***(Optional)*** |
| `subnetwork` | ***string***| ***(Optional)*** |
| `subnetworkProject` | ***string***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecNetworkInterfaceAccessConfig

Appears on:[ComputeInstanceFromTemplateSpecNetworkInterface](#computeinstancefromtemplatespecnetworkinterface)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `assignedNATIP` | ***string***| ***(Optional)*** Deprecated|
| `natIP` | ***string***| ***(Optional)*** |
| `networkTier` | ***string***| ***(Optional)*** |
| `publicPtrDomainName` | ***string***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecNetworkInterfaceAliasIPRange

Appears on:[ComputeInstanceFromTemplateSpecNetworkInterface](#computeinstancefromtemplatespecnetworkinterface)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ipCIDRRange` | ***string***||
| `subnetworkRangeName` | ***string***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecScheduling

Appears on:[ComputeInstanceFromTemplateSpec](#computeinstancefromtemplatespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `automaticRestart` | ***bool***| ***(Optional)*** |
| `onHostMaintenance` | ***string***| ***(Optional)*** |
| `preemptible` | ***bool***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecScratchDisk

Appears on:[ComputeInstanceFromTemplateSpec](#computeinstancefromtemplatespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `interface` | ***string***| ***(Optional)*** |
## ComputeInstanceFromTemplateSpecServiceAccount

Appears on:[ComputeInstanceFromTemplateSpec](#computeinstancefromtemplatespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `email` | ***string***| ***(Optional)*** |
| `scopes` | ***[]string***||
## ComputeInstanceFromTemplateStatus

Appears on:[ComputeInstanceFromTemplate](#computeinstancefromtemplate)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeInstanceFromTemplateSpec](#computeinstancefromtemplatespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `attached_disk.<index>.disk_encryption_key_raw` | ***string*** ||
| `boot_disk.<index>.disk_encryption_key_raw` | ***string*** ||
