---
title: RouteTableAssociation
menu:
  docs_v0.0.1:
    identifier: routetableassociation-aws.kubeform.com
    name: RouteTableAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## RouteTableAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `RouteTableAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RouteTableAssociationSpec](#RouteTableAssociationSpec)***||
| `status` | ***[RouteTableAssociationStatus](#RouteTableAssociationStatus)***||
## RouteTableAssociationSpec
##### (Appears on:[RouteTableAssociation](#RouteTableAssociation), [RouteTableAssociationStatus](#RouteTableAssociationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `routeTableID` | ***string***||
| `subnetID` | ***string***||
## RouteTableAssociationStatus
##### (Appears on:[RouteTableAssociation](#RouteTableAssociation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RouteTableAssociationSpec](#RouteTableAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
