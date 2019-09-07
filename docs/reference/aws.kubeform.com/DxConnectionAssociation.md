---
title: DxConnectionAssociation
menu:
  docs_v0.0.1:
    identifier: dxconnectionassociation-aws.kubeform.com
    name: DxConnectionAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DxConnectionAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DxConnectionAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DxConnectionAssociationSpec](#DxConnectionAssociationSpec)***||
| `status` | ***[DxConnectionAssociationStatus](#DxConnectionAssociationStatus)***||
## DxConnectionAssociationSpec
##### (Appears on:[DxConnectionAssociation](#DxConnectionAssociation), [DxConnectionAssociationStatus](#DxConnectionAssociationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `connectionID` | ***string***||
| `lagID` | ***string***||
## DxConnectionAssociationStatus
##### (Appears on:[DxConnectionAssociation](#DxConnectionAssociation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DxConnectionAssociationSpec](#DxConnectionAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
