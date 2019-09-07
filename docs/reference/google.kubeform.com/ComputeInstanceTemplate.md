---
title: ComputeInstanceTemplate
menu:
  docs_v0.0.1:
    identifier: computeinstancetemplate-google.kubeform.com
    name: ComputeInstanceTemplate
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ComputeInstanceTemplate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeInstanceTemplate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeInstanceTemplateSpec](#ComputeInstanceTemplateSpec)***||
| `status` | ***[ComputeInstanceTemplateStatus](#ComputeInstanceTemplateStatus)***||
## ComputeInstanceTemplateSpec
##### (Appears on:[ComputeInstanceTemplate](#ComputeInstanceTemplate), [ComputeInstanceTemplateStatus](#ComputeInstanceTemplateStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `canIPForward` | ***bool***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `disk` | ***[[]ComputeInstanceTemplateSpecDisk](#ComputeInstanceTemplateSpecDisk)***||
| `guestAccelerator` | ***[[]ComputeInstanceTemplateSpecGuestAccelerator](#ComputeInstanceTemplateSpecGuestAccelerator)***| ***(Optional)*** |
| `instanceDescription` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `machineType` | ***string***||
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `metadataFingerprint` | ***string***| ***(Optional)*** |
| `metadataStartupScript` | ***string***| ***(Optional)*** |
| `minCPUPlatform` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `networkInterface` | ***[[]ComputeInstanceTemplateSpecNetworkInterface](#ComputeInstanceTemplateSpecNetworkInterface)***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `scheduling` | ***[[]ComputeInstanceTemplateSpecScheduling](#ComputeInstanceTemplateSpecScheduling)***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `serviceAccount` | ***[[]ComputeInstanceTemplateSpecServiceAccount](#ComputeInstanceTemplateSpecServiceAccount)***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
| `tagsFingerprint` | ***string***| ***(Optional)*** |
## ComputeInstanceTemplateSpecDisk
##### (Appears on:[ComputeInstanceTemplateSpec](#ComputeInstanceTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoDelete` | ***bool***| ***(Optional)*** |
| `boot` | ***bool***| ***(Optional)*** |
| `deviceName` | ***string***| ***(Optional)*** |
| `diskEncryptionKey` | ***[[]ComputeInstanceTemplateSpecDiskDiskEncryptionKey](#ComputeInstanceTemplateSpecDiskDiskEncryptionKey)***| ***(Optional)*** |
| `diskName` | ***string***| ***(Optional)*** |
| `diskSizeGb` | ***int***| ***(Optional)*** |
| `diskType` | ***string***| ***(Optional)*** |
| `interface` | ***string***| ***(Optional)*** |
| `mode` | ***string***| ***(Optional)*** |
| `source` | ***string***| ***(Optional)*** |
| `sourceImage` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## ComputeInstanceTemplateSpecDiskDiskEncryptionKey
##### (Appears on:[ComputeInstanceTemplateSpecDisk](#ComputeInstanceTemplateSpecDisk))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `kmsKeySelfLink` | ***string***| ***(Optional)*** |
## ComputeInstanceTemplateSpecGuestAccelerator
##### (Appears on:[ComputeInstanceTemplateSpec](#ComputeInstanceTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***||
| `type` | ***string***||
## ComputeInstanceTemplateSpecNetworkInterface
##### (Appears on:[ComputeInstanceTemplateSpec](#ComputeInstanceTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `accessConfig` | ***[[]ComputeInstanceTemplateSpecNetworkInterfaceAccessConfig](#ComputeInstanceTemplateSpecNetworkInterfaceAccessConfig)***| ***(Optional)*** |
| `address` | ***string***| ***(Optional)*** Deprecated|
| `aliasIPRange` | ***[[]ComputeInstanceTemplateSpecNetworkInterfaceAliasIPRange](#ComputeInstanceTemplateSpecNetworkInterfaceAliasIPRange)***| ***(Optional)*** |
| `network` | ***string***| ***(Optional)*** |
| `networkIP` | ***string***| ***(Optional)*** |
| `subnetwork` | ***string***| ***(Optional)*** |
| `subnetworkProject` | ***string***| ***(Optional)*** |
## ComputeInstanceTemplateSpecNetworkInterfaceAccessConfig
##### (Appears on:[ComputeInstanceTemplateSpecNetworkInterface](#ComputeInstanceTemplateSpecNetworkInterface))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `assignedNATIP` | ***string***| ***(Optional)*** Deprecated|
| `natIP` | ***string***| ***(Optional)*** |
| `networkTier` | ***string***| ***(Optional)*** |
## ComputeInstanceTemplateSpecNetworkInterfaceAliasIPRange
##### (Appears on:[ComputeInstanceTemplateSpecNetworkInterface](#ComputeInstanceTemplateSpecNetworkInterface))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ipCIDRRange` | ***string***||
| `subnetworkRangeName` | ***string***| ***(Optional)*** |
## ComputeInstanceTemplateSpecScheduling
##### (Appears on:[ComputeInstanceTemplateSpec](#ComputeInstanceTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `automaticRestart` | ***bool***| ***(Optional)*** |
| `onHostMaintenance` | ***string***| ***(Optional)*** |
| `preemptible` | ***bool***| ***(Optional)*** |
## ComputeInstanceTemplateSpecServiceAccount
##### (Appears on:[ComputeInstanceTemplateSpec](#ComputeInstanceTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `email` | ***string***| ***(Optional)*** |
| `scopes` | ***[]string***||
## ComputeInstanceTemplateStatus
##### (Appears on:[ComputeInstanceTemplate](#ComputeInstanceTemplate))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeInstanceTemplateSpec](#ComputeInstanceTemplateSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
