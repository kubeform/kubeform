---
title: PrivateLinkEndpoint
menu:
  docs_{{ .version }}:
    identifier: privatelinkendpoint-azurerm.kubeform.com
    name: PrivateLinkEndpoint
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PrivateLinkEndpoint
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PrivateLinkEndpoint` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PrivateLinkEndpointSpec](#privatelinkendpointspec)***||
| `status` | ***[PrivateLinkEndpointStatus](#privatelinkendpointstatus)***||
## Phase(`string` alias)

Appears on:[PrivateLinkEndpointStatus](#privatelinkendpointstatus)

## PrivateLinkEndpointSpec

Appears on:[PrivateLinkEndpoint](#privatelinkendpoint), [PrivateLinkEndpointStatus](#privatelinkendpointstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `privateServiceConnection` | ***[[]PrivateLinkEndpointSpecPrivateServiceConnection](#privatelinkendpointspecprivateserviceconnection)***||
| `resourceGroupName` | ***string***||
| `subnetID` | ***string***||
## PrivateLinkEndpointSpecPrivateServiceConnection

Appears on:[PrivateLinkEndpointSpec](#privatelinkendpointspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `isManualConnection` | ***bool***||
| `name` | ***string***||
| `privateConnectionResourceID` | ***string***||
| `requestMessage` | ***string***| ***(Optional)*** |
| `subresourceNames` | ***[]string***| ***(Optional)*** |
## PrivateLinkEndpointStatus

Appears on:[PrivateLinkEndpoint](#privatelinkendpoint)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PrivateLinkEndpointSpec](#privatelinkendpointspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
