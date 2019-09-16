---
title: Ec2TransitGateway
menu:
  docs_{{ .version }}:
    identifier: ec2transitgateway-aws.kubeform.com
    name: Ec2TransitGateway
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Ec2TransitGateway
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Ec2TransitGateway` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Ec2TransitGatewaySpec](#Ec2TransitGatewaySpec)***||
| `status` | ***[Ec2TransitGatewayStatus](#Ec2TransitGatewayStatus)***||
## Ec2TransitGatewaySpec

Appears on:[Ec2TransitGateway](#Ec2TransitGateway), [Ec2TransitGatewayStatus](#Ec2TransitGatewayStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `amazonSideAsn` | ***int***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `associationDefaultRouteTableID` | ***string***| ***(Optional)*** |
| `autoAcceptSharedAttachments` | ***string***| ***(Optional)*** |
| `defaultRouteTableAssociation` | ***string***| ***(Optional)*** |
| `defaultRouteTablePropagation` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `dnsSupport` | ***string***| ***(Optional)*** |
| `ownerID` | ***string***| ***(Optional)*** |
| `propagationDefaultRouteTableID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpnEcmpSupport` | ***string***| ***(Optional)*** |
## Ec2TransitGatewayStatus

Appears on:[Ec2TransitGateway](#Ec2TransitGateway)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Ec2TransitGatewaySpec](#Ec2TransitGatewaySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
