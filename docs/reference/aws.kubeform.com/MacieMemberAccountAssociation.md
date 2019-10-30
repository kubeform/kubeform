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
| `spec` | ***[MacieMemberAccountAssociationSpec](#maciememberaccountassociationspec)***||
| `status` | ***[MacieMemberAccountAssociationStatus](#maciememberaccountassociationstatus)***||
## MacieMemberAccountAssociationSpec

Appears on:[MacieMemberAccountAssociation](#maciememberaccountassociation), [MacieMemberAccountAssociationStatus](#maciememberaccountassociationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `memberAccountID` | ***string***||
## MacieMemberAccountAssociationStatus

Appears on:[MacieMemberAccountAssociation](#maciememberaccountassociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MacieMemberAccountAssociationSpec](#maciememberaccountassociationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[MacieMemberAccountAssociationStatus](#maciememberaccountassociationstatus)

---
