---
title: DefaultRouteTable
menu:
  docs_{{ .version }}:
    identifier: defaultroutetable-aws.kubeform.com
    name: DefaultRouteTable
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DefaultRouteTable
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DefaultRouteTable` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DefaultRouteTableSpec](#defaultroutetablespec)***||
| `status` | ***[DefaultRouteTableStatus](#defaultroutetablestatus)***||
## DefaultRouteTableSpec

Appears on:[DefaultRouteTable](#defaultroutetable), [DefaultRouteTableStatus](#defaultroutetablestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `defaultRouteTableID` | ***string***||
| `ownerID` | ***string***| ***(Optional)*** |
| `propagatingVgws` | ***[]string***| ***(Optional)*** |
| `route` | ***[[]DefaultRouteTableSpecRoute](#defaultroutetablespecroute)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
## DefaultRouteTableSpecRoute

Appears on:[DefaultRouteTableSpec](#defaultroutetablespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cidrBlock` | ***string***| ***(Optional)*** |
| `egressOnlyGatewayID` | ***string***| ***(Optional)*** |
| `gatewayID` | ***string***| ***(Optional)*** |
| `instanceID` | ***string***| ***(Optional)*** |
| `ipv6CIDRBlock` | ***string***| ***(Optional)*** |
| `natGatewayID` | ***string***| ***(Optional)*** |
| `networkInterfaceID` | ***string***| ***(Optional)*** |
| `transitGatewayID` | ***string***| ***(Optional)*** |
| `vpcPeeringConnectionID` | ***string***| ***(Optional)*** |
## DefaultRouteTableStatus

Appears on:[DefaultRouteTable](#defaultroutetable)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DefaultRouteTableSpec](#defaultroutetablespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
