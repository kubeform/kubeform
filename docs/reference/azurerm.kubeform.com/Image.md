---
title: Image
menu:
  docs_v0.0.1:
    identifier: image-azurerm.kubeform.com
    name: Image
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## Image
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `Image` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ImageSpec](#ImageSpec)***||
| `status` | ***[ImageStatus](#ImageStatus)***||
## ImageSpec
##### (Appears on:[Image](#Image), [ImageStatus](#ImageStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dataDisk` | ***[[]ImageSpecDataDisk](#ImageSpecDataDisk)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `osDisk` | ***[[]ImageSpecOsDisk](#ImageSpecOsDisk)***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `sourceVirtualMachineID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `zoneResilient` | ***bool***| ***(Optional)*** |
## ImageSpecDataDisk
##### (Appears on:[ImageSpec](#ImageSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `blobURI` | ***string***| ***(Optional)*** |
| `caching` | ***string***| ***(Optional)*** |
| `lun` | ***int***| ***(Optional)*** |
| `managedDiskID` | ***string***| ***(Optional)*** |
| `sizeGb` | ***int***| ***(Optional)*** |
## ImageSpecOsDisk
##### (Appears on:[ImageSpec](#ImageSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `blobURI` | ***string***| ***(Optional)*** |
| `caching` | ***string***| ***(Optional)*** |
| `managedDiskID` | ***string***| ***(Optional)*** |
| `osState` | ***string***| ***(Optional)*** |
| `osType` | ***string***| ***(Optional)*** |
| `sizeGb` | ***int***| ***(Optional)*** |
## ImageStatus
##### (Appears on:[Image](#Image))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ImageSpec](#ImageSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
