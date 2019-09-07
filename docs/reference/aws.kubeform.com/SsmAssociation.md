---
title: SsmAssociation
menu:
  docs_v0.0.1:
    identifier: ssmassociation-aws.kubeform.com
    name: SsmAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## SsmAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SsmAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SsmAssociationSpec](#SsmAssociationSpec)***||
| `status` | ***[SsmAssociationStatus](#SsmAssociationStatus)***||
## SsmAssociationSpec
##### (Appears on:[SsmAssociation](#SsmAssociation), [SsmAssociationStatus](#SsmAssociationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `associationID` | ***string***| ***(Optional)*** |
| `associationName` | ***string***| ***(Optional)*** |
| `complianceSeverity` | ***string***| ***(Optional)*** |
| `documentVersion` | ***string***| ***(Optional)*** |
| `instanceID` | ***string***| ***(Optional)*** |
| `maxConcurrency` | ***string***| ***(Optional)*** |
| `maxErrors` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `outputLocation` | ***[[]SsmAssociationSpecOutputLocation](#SsmAssociationSpecOutputLocation)***| ***(Optional)*** |
| `parameters` | ***map[string]string***| ***(Optional)*** |
| `scheduleExpression` | ***string***| ***(Optional)*** |
| `targets` | ***[[]SsmAssociationSpecTargets](#SsmAssociationSpecTargets)***| ***(Optional)*** |
## SsmAssociationSpecOutputLocation
##### (Appears on:[SsmAssociationSpec](#SsmAssociationSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `s3BucketName` | ***string***||
| `s3KeyPrefix` | ***string***| ***(Optional)*** |
## SsmAssociationSpecTargets
##### (Appears on:[SsmAssociationSpec](#SsmAssociationSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***||
| `values` | ***[]string***||
## SsmAssociationStatus
##### (Appears on:[SsmAssociation](#SsmAssociation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SsmAssociationSpec](#SsmAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
