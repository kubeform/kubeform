---
title: SubnetNATGatewayAssociation
menu:
  docs_{{ .version }}:
    identifier: subnetnatgatewayassociation-azurerm.kubeform.com
    name: SubnetNATGatewayAssociation
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SubnetNATGatewayAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SubnetNATGatewayAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SubnetNATGatewayAssociationSpec](#subnetnatgatewayassociationspec)***||
| `status` | ***[SubnetNATGatewayAssociationStatus](#subnetnatgatewayassociationstatus)***||
## Phase(`string` alias)

Appears on:[SubnetNATGatewayAssociationStatus](#subnetnatgatewayassociationstatus)

## SubnetNATGatewayAssociationSpec

Appears on:[SubnetNATGatewayAssociation](#subnetnatgatewayassociation), [SubnetNATGatewayAssociationStatus](#subnetnatgatewayassociationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `natGatewayID` | ***string***||
| `subnetID` | ***string***||
## SubnetNATGatewayAssociationStatus

Appears on:[SubnetNATGatewayAssociation](#subnetnatgatewayassociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SubnetNATGatewayAssociationSpec](#subnetnatgatewayassociationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
