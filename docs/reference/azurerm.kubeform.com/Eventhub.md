---
title: Eventhub
menu:
  docs_v0.0.1:
    identifier: eventhub-azurerm.kubeform.com
    name: Eventhub
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## Eventhub
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `Eventhub` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EventhubSpec](#EventhubSpec)***||
| `status` | ***[EventhubStatus](#EventhubStatus)***||
## EventhubSpec
##### (Appears on:[Eventhub](#Eventhub), [EventhubStatus](#EventhubStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `captureDescription` | ***[[]EventhubSpecCaptureDescription](#EventhubSpecCaptureDescription)***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** Deprecated|
| `messageRetention` | ***int***||
| `name` | ***string***||
| `namespaceName` | ***string***||
| `partitionCount` | ***int***||
| `partitionIDS` | ***[]string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
## EventhubSpecCaptureDescription
##### (Appears on:[EventhubSpec](#EventhubSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `destination` | ***[[]EventhubSpecCaptureDescriptionDestination](#EventhubSpecCaptureDescriptionDestination)***||
| `enabled` | ***bool***||
| `encoding` | ***string***||
| `intervalInSeconds` | ***int***| ***(Optional)*** |
| `sizeLimitInBytes` | ***int***| ***(Optional)*** |
| `skipEmptyArchives` | ***bool***| ***(Optional)*** |
## EventhubSpecCaptureDescriptionDestination
##### (Appears on:[EventhubSpecCaptureDescription](#EventhubSpecCaptureDescription))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `archiveNameFormat` | ***string***||
| `blobContainerName` | ***string***||
| `name` | ***string***||
| `storageAccountID` | ***string***||
## EventhubStatus
##### (Appears on:[Eventhub](#Eventhub))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EventhubSpec](#EventhubSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
