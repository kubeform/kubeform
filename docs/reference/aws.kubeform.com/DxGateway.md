---
title: DxGateway
menu:
  docs_{{ .version }}:
    identifier: dxgateway-aws.kubeform.com
    name: DxGateway
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DxGateway
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DxGateway` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DxGatewaySpec](#dxgatewayspec)***||
| `status` | ***[DxGatewayStatus](#dxgatewaystatus)***||
## DxGatewaySpec

Appears on:[DxGateway](#dxgateway), [DxGatewayStatus](#dxgatewaystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `amazonSideAsn` | ***string***||
| `name` | ***string***||
| `ownerAccountID` | ***string***| ***(Optional)*** |
## DxGatewayStatus

Appears on:[DxGateway](#dxgateway)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DxGatewaySpec](#dxgatewayspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
