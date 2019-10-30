---
title: Route
menu:
  docs_{{ .version }}:
    identifier: route-azurerm.kubeform.com
    name: Route
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Route
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `Route` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RouteSpec](#routespec)***||
| `status` | ***[RouteStatus](#routestatus)***||
## Phase(`string` alias)

Appears on:[RouteStatus](#routestatus)

## RouteSpec

Appears on:[Route](#route), [RouteStatus](#routestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `addressPrefix` | ***string***||
| `name` | ***string***||
| `nextHopInIPAddress` | ***string***| ***(Optional)*** |
| `nextHopType` | ***string***||
| `resourceGroupName` | ***string***||
| `routeTableName` | ***string***||
## RouteStatus

Appears on:[Route](#route)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RouteSpec](#routespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
