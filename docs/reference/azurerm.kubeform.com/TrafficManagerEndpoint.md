---
title: TrafficManagerEndpoint
menu:
  docs_{{ .version }}:
    identifier: trafficmanagerendpoint-azurerm.kubeform.com
    name: TrafficManagerEndpoint
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## TrafficManagerEndpoint
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `TrafficManagerEndpoint` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[TrafficManagerEndpointSpec](#trafficmanagerendpointspec)***||
| `status` | ***[TrafficManagerEndpointStatus](#trafficmanagerendpointstatus)***||
## Phase(`string` alias)

Appears on:[TrafficManagerEndpointStatus](#trafficmanagerendpointstatus)

## TrafficManagerEndpointSpec

Appears on:[TrafficManagerEndpoint](#trafficmanagerendpoint), [TrafficManagerEndpointStatus](#trafficmanagerendpointstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `customHeader` | ***[[]TrafficManagerEndpointSpecCustomHeader](#trafficmanagerendpointspeccustomheader)***| ***(Optional)*** |
| `endpointLocation` | ***string***| ***(Optional)*** |
| `endpointMonitorStatus` | ***string***| ***(Optional)*** |
| `endpointStatus` | ***string***| ***(Optional)*** |
| `geoMappings` | ***[]string***| ***(Optional)*** |
| `minChildEndpoints` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
| `priority` | ***int64***| ***(Optional)*** |
| `profileName` | ***string***||
| `resourceGroupName` | ***string***||
| `subnet` | ***[[]TrafficManagerEndpointSpecSubnet](#trafficmanagerendpointspecsubnet)***| ***(Optional)*** |
| `target` | ***string***| ***(Optional)*** |
| `targetResourceID` | ***string***| ***(Optional)*** |
| `type` | ***string***||
| `weight` | ***int64***| ***(Optional)*** |
## TrafficManagerEndpointSpecCustomHeader

Appears on:[TrafficManagerEndpointSpec](#trafficmanagerendpointspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `value` | ***string***||
## TrafficManagerEndpointSpecSubnet

Appears on:[TrafficManagerEndpointSpec](#trafficmanagerendpointspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `first` | ***string***||
| `last` | ***string***| ***(Optional)*** |
| `scope` | ***int64***| ***(Optional)*** |
## TrafficManagerEndpointStatus

Appears on:[TrafficManagerEndpoint](#trafficmanagerendpoint)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[TrafficManagerEndpointSpec](#trafficmanagerendpointspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
