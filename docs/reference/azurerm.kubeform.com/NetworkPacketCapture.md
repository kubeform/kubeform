---
title: NetworkPacketCapture
menu:
  docs_{{ .version }}:
    identifier: networkpacketcapture-azurerm.kubeform.com
    name: NetworkPacketCapture
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NetworkPacketCapture
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkPacketCapture` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkPacketCaptureSpec](#networkpacketcapturespec)***||
| `status` | ***[NetworkPacketCaptureStatus](#networkpacketcapturestatus)***||
## NetworkPacketCaptureSpec

Appears on:[NetworkPacketCapture](#networkpacketcapture), [NetworkPacketCaptureStatus](#networkpacketcapturestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `filter` | ***[[]NetworkPacketCaptureSpecFilter](#networkpacketcapturespecfilter)***| ***(Optional)*** |
| `maximumBytesPerPacket` | ***int***| ***(Optional)*** |
| `maximumBytesPerSession` | ***int***| ***(Optional)*** |
| `maximumCaptureDuration` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `networkWatcherName` | ***string***||
| `resourceGroupName` | ***string***||
| `storageLocation` | ***[[]NetworkPacketCaptureSpecStorageLocation](#networkpacketcapturespecstoragelocation)***||
| `targetResourceID` | ***string***||
## NetworkPacketCaptureSpecFilter

Appears on:[NetworkPacketCaptureSpec](#networkpacketcapturespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `localIPAddress` | ***string***| ***(Optional)*** |
| `localPort` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
| `remoteIPAddress` | ***string***| ***(Optional)*** |
| `remotePort` | ***string***| ***(Optional)*** |
## NetworkPacketCaptureSpecStorageLocation

Appears on:[NetworkPacketCaptureSpec](#networkpacketcapturespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `filePath` | ***string***| ***(Optional)*** |
| `storageAccountID` | ***string***| ***(Optional)*** |
| `storagePath` | ***string***| ***(Optional)*** |
## NetworkPacketCaptureStatus

Appears on:[NetworkPacketCapture](#networkpacketcapture)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkPacketCaptureSpec](#networkpacketcapturespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
