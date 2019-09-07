---
title: Route
menu:
  docs_v0.0.1:
    identifier: route-azurerm.kubeform.com
    name: Route
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## Route
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `Route` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RouteSpec](#RouteSpec)***||
| `status` | ***[RouteStatus](#RouteStatus)***||
## RouteSpec
##### (Appears on:[Route](#Route), [RouteStatus](#RouteStatus))
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
##### (Appears on:[Route](#Route))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RouteSpec](#RouteSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
