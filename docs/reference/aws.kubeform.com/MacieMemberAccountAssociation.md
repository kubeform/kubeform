---
title: MacieMemberAccountAssociation
menu:
  docs_{{ .version }}:
    identifier: maciememberaccountassociation-aws.kubeform.com
    name: MacieMemberAccountAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MacieMemberAccountAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `MacieMemberAccountAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MacieMemberAccountAssociationSpec](#MacieMemberAccountAssociationSpec)***||
| `status` | ***[MacieMemberAccountAssociationStatus](#MacieMemberAccountAssociationStatus)***||
## MacieMemberAccountAssociationSpec

Appears on:[MacieMemberAccountAssociation](#MacieMemberAccountAssociation), [MacieMemberAccountAssociationStatus](#MacieMemberAccountAssociationStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `memberAccountID` | ***string***||
## MacieMemberAccountAssociationStatus

Appears on:[MacieMemberAccountAssociation](#MacieMemberAccountAssociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MacieMemberAccountAssociationSpec](#MacieMemberAccountAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
