---
title: PrivateEndpoint
menu:
  docs_{{ .version }}:
    identifier: privateendpoint-azurerm.kubeform.com
    name: PrivateEndpoint
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PrivateEndpoint
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PrivateEndpoint` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PrivateEndpointSpec](#privateendpointspec)***||
| `status` | ***[PrivateEndpointStatus](#privateendpointstatus)***||
## Phase(`string` alias)

Appears on:[PrivateEndpointStatus](#privateendpointstatus)

## PrivateEndpointSpec

Appears on:[PrivateEndpoint](#privateendpoint), [PrivateEndpointStatus](#privateendpointstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `privateServiceConnection` | ***[[]PrivateEndpointSpecPrivateServiceConnection](#privateendpointspecprivateserviceconnection)***||
| `resourceGroupName` | ***string***||
| `subnetID` | ***string***||
## PrivateEndpointSpecPrivateServiceConnection

Appears on:[PrivateEndpointSpec](#privateendpointspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `isManualConnection` | ***bool***||
| `name` | ***string***||
| `privateConnectionResourceID` | ***string***||
| `requestMessage` | ***string***| ***(Optional)*** |
| `subresourceNames` | ***[]string***| ***(Optional)*** |
## PrivateEndpointStatus

Appears on:[PrivateEndpoint](#privateendpoint)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PrivateEndpointSpec](#privateendpointspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
