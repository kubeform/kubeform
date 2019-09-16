---
title: PacketCapture
menu:
  docs_{{ .version }}:
    identifier: packetcapture-azurerm.kubeform.com
    name: PacketCapture
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PacketCapture
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PacketCapture` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PacketCaptureSpec](#PacketCaptureSpec)***||
| `status` | ***[PacketCaptureStatus](#PacketCaptureStatus)***||
## PacketCaptureSpec

Appears on:[PacketCapture](#PacketCapture), [PacketCaptureStatus](#PacketCaptureStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `filter` | ***[[]PacketCaptureSpecFilter](#PacketCaptureSpecFilter)***| ***(Optional)*** |
| `maximumBytesPerPacket` | ***int***| ***(Optional)*** |
| `maximumBytesPerSession` | ***int***| ***(Optional)*** |
| `maximumCaptureDuration` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `networkWatcherName` | ***string***||
| `resourceGroupName` | ***string***||
| `storageLocation` | ***[[]PacketCaptureSpecStorageLocation](#PacketCaptureSpecStorageLocation)***||
| `targetResourceID` | ***string***||
## PacketCaptureSpecFilter

Appears on:[PacketCaptureSpec](#PacketCaptureSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `localIPAddress` | ***string***| ***(Optional)*** |
| `localPort` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
| `remoteIPAddress` | ***string***| ***(Optional)*** |
| `remotePort` | ***string***| ***(Optional)*** |
## PacketCaptureSpecStorageLocation

Appears on:[PacketCaptureSpec](#PacketCaptureSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `filePath` | ***string***| ***(Optional)*** |
| `storageAccountID` | ***string***| ***(Optional)*** |
| `storagePath` | ***string***| ***(Optional)*** |
## PacketCaptureStatus

Appears on:[PacketCapture](#PacketCapture)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PacketCaptureSpec](#PacketCaptureSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
