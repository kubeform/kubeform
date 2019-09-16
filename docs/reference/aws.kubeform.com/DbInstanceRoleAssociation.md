---
title: DbInstanceRoleAssociation
menu:
  docs_{{ .version }}:
    identifier: dbinstanceroleassociation-aws.kubeform.com
    name: DbInstanceRoleAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DbInstanceRoleAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DbInstanceRoleAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DbInstanceRoleAssociationSpec](#dbinstanceroleassociationspec)***||
| `status` | ***[DbInstanceRoleAssociationStatus](#dbinstanceroleassociationstatus)***||
## DbInstanceRoleAssociationSpec

Appears on:[DbInstanceRoleAssociation](#dbinstanceroleassociation), [DbInstanceRoleAssociationStatus](#dbinstanceroleassociationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dbInstanceIdentifier` | ***string***||
| `featureName` | ***string***||
| `roleArn` | ***string***||
## DbInstanceRoleAssociationStatus

Appears on:[DbInstanceRoleAssociation](#dbinstanceroleassociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DbInstanceRoleAssociationSpec](#dbinstanceroleassociationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
