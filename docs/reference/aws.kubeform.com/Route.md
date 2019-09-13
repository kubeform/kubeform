---
title: Route
menu:
  docs_{{ .version }}:
    identifier: route-aws.kubeform.com
    name: Route
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Route
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
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
| `destinationCIDRBlock` | ***string***| ***(Optional)*** |
| `destinationIpv6CIDRBlock` | ***string***| ***(Optional)*** |
| `destinationPrefixListID` | ***string***| ***(Optional)*** |
| `egressOnlyGatewayID` | ***string***| ***(Optional)*** |
| `gatewayID` | ***string***| ***(Optional)*** |
| `instanceID` | ***string***| ***(Optional)*** |
| `instanceOwnerID` | ***string***| ***(Optional)*** |
| `natGatewayID` | ***string***| ***(Optional)*** |
| `networkInterfaceID` | ***string***| ***(Optional)*** |
| `origin` | ***string***| ***(Optional)*** |
| `routeTableID` | ***string***||
| `state` | ***string***| ***(Optional)*** |
| `transitGatewayID` | ***string***| ***(Optional)*** |
| `vpcPeeringConnectionID` | ***string***| ***(Optional)*** |
## RouteStatus
##### (Appears on:[Route](#Route))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RouteSpec](#RouteSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
