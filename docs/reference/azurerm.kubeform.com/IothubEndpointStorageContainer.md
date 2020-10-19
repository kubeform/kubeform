---
title: IothubEndpointStorageContainer
menu:
  docs_{{ .version }}:
    identifier: iothubendpointstoragecontainer-azurerm.kubeform.com
    name: IothubEndpointStorageContainer
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IothubEndpointStorageContainer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `IothubEndpointStorageContainer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IothubEndpointStorageContainerSpec](#iothubendpointstoragecontainerspec)***||
| `status` | ***[IothubEndpointStorageContainerStatus](#iothubendpointstoragecontainerstatus)***||
## IothubEndpointStorageContainerSpec

Appears on:[IothubEndpointStorageContainer](#iothubendpointstoragecontainer), [IothubEndpointStorageContainerStatus](#iothubendpointstoragecontainerstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `batchFrequencyInSeconds` | ***int64***| ***(Optional)*** |
| `containerName` | ***string***||
| `encoding` | ***string***| ***(Optional)*** |
| `fileNameFormat` | ***string***| ***(Optional)*** |
| `iothubName` | ***string***||
| `maxChunkSizeInBytes` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
## IothubEndpointStorageContainerStatus

Appears on:[IothubEndpointStorageContainer](#iothubendpointstoragecontainer)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IothubEndpointStorageContainerSpec](#iothubendpointstoragecontainerspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IothubEndpointStorageContainerStatus](#iothubendpointstoragecontainerstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `connection_string` | ***string*** ||
