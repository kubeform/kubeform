---
title: NatGateway
menu:
  docs_v0.0.1:
    identifier: natgateway-aws.kubeform.com
    name: NatGateway
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## NatGateway
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NatGateway` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NatGatewaySpec](#NatGatewaySpec)***||
| `status` | ***[NatGatewayStatus](#NatGatewayStatus)***||
## NatGatewaySpec
##### (Appears on:[NatGateway](#NatGateway), [NatGatewayStatus](#NatGatewayStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allocationID` | ***string***||
| `networkInterfaceID` | ***string***| ***(Optional)*** |
| `privateIP` | ***string***| ***(Optional)*** |
| `publicIP` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## NatGatewayStatus
##### (Appears on:[NatGateway](#NatGateway))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NatGatewaySpec](#NatGatewaySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
