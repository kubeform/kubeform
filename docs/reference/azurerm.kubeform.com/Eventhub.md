---
title: Eventhub
menu:
  docs_{{ .version }}:
    identifier: eventhub-azurerm.kubeform.com
    name: Eventhub
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Eventhub
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `Eventhub` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EventhubSpec](#eventhubspec)***||
| `status` | ***[EventhubStatus](#eventhubstatus)***||
## EventhubSpec

Appears on:[Eventhub](#eventhub), [EventhubStatus](#eventhubstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `captureDescription` | ***[[]EventhubSpecCaptureDescription](#eventhubspeccapturedescription)***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** Deprecated|
| `messageRetention` | ***int64***||
| `name` | ***string***||
| `namespaceName` | ***string***||
| `partitionCount` | ***int64***||
| `partitionIDS` | ***[]string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
## EventhubSpecCaptureDescription

Appears on:[EventhubSpec](#eventhubspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `destination` | ***[[]EventhubSpecCaptureDescriptionDestination](#eventhubspeccapturedescriptiondestination)***||
| `enabled` | ***bool***||
| `encoding` | ***string***||
| `intervalInSeconds` | ***int64***| ***(Optional)*** |
| `sizeLimitInBytes` | ***int64***| ***(Optional)*** |
| `skipEmptyArchives` | ***bool***| ***(Optional)*** |
## EventhubSpecCaptureDescriptionDestination

Appears on:[EventhubSpecCaptureDescription](#eventhubspeccapturedescription)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `archiveNameFormat` | ***string***||
| `blobContainerName` | ***string***||
| `name` | ***string***||
| `storageAccountID` | ***string***||
## EventhubStatus

Appears on:[Eventhub](#eventhub)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EventhubSpec](#eventhubspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[EventhubStatus](#eventhubstatus)

---
