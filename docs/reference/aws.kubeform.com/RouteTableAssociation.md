---
title: RouteTableAssociation
menu:
  docs_{{ .version }}:
    identifier: routetableassociation-aws.kubeform.com
    name: RouteTableAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RouteTableAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `RouteTableAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RouteTableAssociationSpec](#routetableassociationspec)***||
| `status` | ***[RouteTableAssociationStatus](#routetableassociationstatus)***||
## Phase(`string` alias)

Appears on:[RouteTableAssociationStatus](#routetableassociationstatus)

## RouteTableAssociationSpec

Appears on:[RouteTableAssociation](#routetableassociation), [RouteTableAssociationStatus](#routetableassociationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `routeTableID` | ***string***||
| `subnetID` | ***string***||
## RouteTableAssociationStatus

Appears on:[RouteTableAssociation](#routetableassociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RouteTableAssociationSpec](#routetableassociationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
