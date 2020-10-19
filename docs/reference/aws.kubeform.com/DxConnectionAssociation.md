---
title: DxConnectionAssociation
menu:
  docs_{{ .version }}:
    identifier: dxconnectionassociation-aws.kubeform.com
    name: DxConnectionAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DxConnectionAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DxConnectionAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DxConnectionAssociationSpec](#dxconnectionassociationspec)***||
| `status` | ***[DxConnectionAssociationStatus](#dxconnectionassociationstatus)***||
## DxConnectionAssociationSpec

Appears on:[DxConnectionAssociation](#dxconnectionassociation), [DxConnectionAssociationStatus](#dxconnectionassociationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `connectionID` | ***string***||
| `lagID` | ***string***||
## DxConnectionAssociationStatus

Appears on:[DxConnectionAssociation](#dxconnectionassociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DxConnectionAssociationSpec](#dxconnectionassociationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DxConnectionAssociationStatus](#dxconnectionassociationstatus)

---
