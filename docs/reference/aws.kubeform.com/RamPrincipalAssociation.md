---
title: RamPrincipalAssociation
menu:
  docs_{{ .version }}:
    identifier: ramprincipalassociation-aws.kubeform.com
    name: RamPrincipalAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RamPrincipalAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `RamPrincipalAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RamPrincipalAssociationSpec](#ramprincipalassociationspec)***||
| `status` | ***[RamPrincipalAssociationStatus](#ramprincipalassociationstatus)***||
## Phase(`string` alias)

Appears on:[RamPrincipalAssociationStatus](#ramprincipalassociationstatus)

## RamPrincipalAssociationSpec

Appears on:[RamPrincipalAssociation](#ramprincipalassociation), [RamPrincipalAssociationStatus](#ramprincipalassociationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `principal` | ***string***||
| `resourceShareArn` | ***string***||
## RamPrincipalAssociationStatus

Appears on:[RamPrincipalAssociation](#ramprincipalassociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RamPrincipalAssociationSpec](#ramprincipalassociationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
