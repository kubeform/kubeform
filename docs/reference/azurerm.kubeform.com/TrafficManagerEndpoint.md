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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[TrafficManagerEndpointSpec](#TrafficManagerEndpointSpec)***||
| `status` | ***[TrafficManagerEndpointStatus](#TrafficManagerEndpointStatus)***||
## TrafficManagerEndpointSpec
##### (Appears on:[TrafficManagerEndpoint](#TrafficManagerEndpoint), [TrafficManagerEndpointStatus](#TrafficManagerEndpointStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `customHeader` | ***[[]TrafficManagerEndpointSpecCustomHeader](#TrafficManagerEndpointSpecCustomHeader)***| ***(Optional)*** |
| `endpointLocation` | ***string***| ***(Optional)*** |
| `endpointMonitorStatus` | ***string***| ***(Optional)*** |
| `endpointStatus` | ***string***| ***(Optional)*** |
| `geoMappings` | ***[]string***| ***(Optional)*** |
| `minChildEndpoints` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `priority` | ***int***| ***(Optional)*** |
| `profileName` | ***string***||
| `resourceGroupName` | ***string***||
| `subnet` | ***[[]TrafficManagerEndpointSpecSubnet](#TrafficManagerEndpointSpecSubnet)***| ***(Optional)*** |
| `target` | ***string***| ***(Optional)*** |
| `targetResourceID` | ***string***| ***(Optional)*** |
| `type` | ***string***||
| `weight` | ***int***| ***(Optional)*** |
## TrafficManagerEndpointSpecCustomHeader
##### (Appears on:[TrafficManagerEndpointSpec](#TrafficManagerEndpointSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `value` | ***string***||
## TrafficManagerEndpointSpecSubnet
##### (Appears on:[TrafficManagerEndpointSpec](#TrafficManagerEndpointSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `first` | ***string***||
| `last` | ***string***| ***(Optional)*** |
| `scope` | ***int***| ***(Optional)*** |
## TrafficManagerEndpointStatus
##### (Appears on:[TrafficManagerEndpoint](#TrafficManagerEndpoint))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[TrafficManagerEndpointSpec](#TrafficManagerEndpointSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
