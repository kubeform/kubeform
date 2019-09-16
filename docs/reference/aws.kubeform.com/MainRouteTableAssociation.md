---
title: MainRouteTableAssociation
menu:
  docs_{{ .version }}:
    identifier: mainroutetableassociation-aws.kubeform.com
    name: MainRouteTableAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MainRouteTableAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `MainRouteTableAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MainRouteTableAssociationSpec](#mainroutetableassociationspec)***||
| `status` | ***[MainRouteTableAssociationStatus](#mainroutetableassociationstatus)***||
## MainRouteTableAssociationSpec

Appears on:[MainRouteTableAssociation](#mainroutetableassociation), [MainRouteTableAssociationStatus](#mainroutetableassociationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `originalRouteTableID` | ***string***| ***(Optional)*** |
| `routeTableID` | ***string***||
| `vpcID` | ***string***||
## MainRouteTableAssociationStatus

Appears on:[MainRouteTableAssociation](#mainroutetableassociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MainRouteTableAssociationSpec](#mainroutetableassociationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
