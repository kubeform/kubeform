---
title: EgressOnlyInternetGateway
menu:
  docs_{{ .version }}:
    identifier: egressonlyinternetgateway-aws.kubeform.com
    name: EgressOnlyInternetGateway
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## EgressOnlyInternetGateway
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `EgressOnlyInternetGateway` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EgressOnlyInternetGatewaySpec](#EgressOnlyInternetGatewaySpec)***||
| `status` | ***[EgressOnlyInternetGatewayStatus](#EgressOnlyInternetGatewayStatus)***||
## EgressOnlyInternetGatewaySpec
##### (Appears on:[EgressOnlyInternetGateway](#EgressOnlyInternetGateway), [EgressOnlyInternetGatewayStatus](#EgressOnlyInternetGatewayStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `vpcID` | ***string***||
## EgressOnlyInternetGatewayStatus
##### (Appears on:[EgressOnlyInternetGateway](#EgressOnlyInternetGateway))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EgressOnlyInternetGatewaySpec](#EgressOnlyInternetGatewaySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
