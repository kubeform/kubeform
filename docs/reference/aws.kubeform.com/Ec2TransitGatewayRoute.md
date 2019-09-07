---
title: Ec2TransitGatewayRoute
menu:
  docs_v0.0.1:
    identifier: ec2transitgatewayroute-aws.kubeform.com
    name: Ec2TransitGatewayRoute
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## Ec2TransitGatewayRoute
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Ec2TransitGatewayRoute` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Ec2TransitGatewayRouteSpec](#Ec2TransitGatewayRouteSpec)***||
| `status` | ***[Ec2TransitGatewayRouteStatus](#Ec2TransitGatewayRouteStatus)***||
## Ec2TransitGatewayRouteSpec
##### (Appears on:[Ec2TransitGatewayRoute](#Ec2TransitGatewayRoute), [Ec2TransitGatewayRouteStatus](#Ec2TransitGatewayRouteStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `destinationCIDRBlock` | ***string***||
| `transitGatewayAttachmentID` | ***string***||
| `transitGatewayRouteTableID` | ***string***||
## Ec2TransitGatewayRouteStatus
##### (Appears on:[Ec2TransitGatewayRoute](#Ec2TransitGatewayRoute))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Ec2TransitGatewayRouteSpec](#Ec2TransitGatewayRouteSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
