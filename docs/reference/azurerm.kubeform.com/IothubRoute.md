---
title: IothubRoute
menu:
  docs_{{ .version }}:
    identifier: iothubroute-azurerm.kubeform.com
    name: IothubRoute
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IothubRoute
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `IothubRoute` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IothubRouteSpec](#iothubroutespec)***||
| `status` | ***[IothubRouteStatus](#iothubroutestatus)***||
## IothubRouteSpec

Appears on:[IothubRoute](#iothubroute), [IothubRouteStatus](#iothubroutestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `condition` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***||
| `endpointNames` | ***[]string***||
| `iothubName` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `source` | ***string***||
## IothubRouteStatus

Appears on:[IothubRoute](#iothubroute)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IothubRouteSpec](#iothubroutespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IothubRouteStatus](#iothubroutestatus)

---
