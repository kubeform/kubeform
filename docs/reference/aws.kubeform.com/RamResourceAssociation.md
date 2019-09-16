---
title: RamResourceAssociation
menu:
  docs_{{ .version }}:
    identifier: ramresourceassociation-aws.kubeform.com
    name: RamResourceAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RamResourceAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `RamResourceAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RamResourceAssociationSpec](#RamResourceAssociationSpec)***||
| `status` | ***[RamResourceAssociationStatus](#RamResourceAssociationStatus)***||
## RamResourceAssociationSpec

Appears on:[RamResourceAssociation](#RamResourceAssociation), [RamResourceAssociationStatus](#RamResourceAssociationStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `resourceArn` | ***string***||
| `resourceShareArn` | ***string***||
## RamResourceAssociationStatus

Appears on:[RamResourceAssociation](#RamResourceAssociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RamResourceAssociationSpec](#RamResourceAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
