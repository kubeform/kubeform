---
title: DxGatewayAssociation
menu:
  docs_{{ .version }}:
    identifier: dxgatewayassociation-aws.kubeform.com
    name: DxGatewayAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DxGatewayAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DxGatewayAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DxGatewayAssociationSpec](#dxgatewayassociationspec)***||
| `status` | ***[DxGatewayAssociationStatus](#dxgatewayassociationstatus)***||
## DxGatewayAssociationSpec

Appears on:[DxGatewayAssociation](#dxgatewayassociation), [DxGatewayAssociationStatus](#dxgatewayassociationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allowedPrefixes` | ***[]string***| ***(Optional)*** |
| `dxGatewayAssociationID` | ***string***| ***(Optional)*** |
| `dxGatewayID` | ***string***||
| `vpnGatewayID` | ***string***||
## DxGatewayAssociationStatus

Appears on:[DxGatewayAssociation](#dxgatewayassociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DxGatewayAssociationSpec](#dxgatewayassociationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DxGatewayAssociationStatus](#dxgatewayassociationstatus)

---
