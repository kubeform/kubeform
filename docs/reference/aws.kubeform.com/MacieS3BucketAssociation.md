---
title: MacieS3BucketAssociation
menu:
  docs_{{ .version }}:
    identifier: macies3bucketassociation-aws.kubeform.com
    name: MacieS3BucketAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MacieS3BucketAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `MacieS3BucketAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MacieS3BucketAssociationSpec](#macies3bucketassociationspec)***||
| `status` | ***[MacieS3BucketAssociationStatus](#macies3bucketassociationstatus)***||
## MacieS3BucketAssociationSpec

Appears on:[MacieS3BucketAssociation](#macies3bucketassociation), [MacieS3BucketAssociationStatus](#macies3bucketassociationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucketName` | ***string***||
| `classificationType` | ***[[]MacieS3BucketAssociationSpecClassificationType](#macies3bucketassociationspecclassificationtype)***| ***(Optional)*** |
| `memberAccountID` | ***string***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
## MacieS3BucketAssociationSpecClassificationType

Appears on:[MacieS3BucketAssociationSpec](#macies3bucketassociationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `continuous` | ***string***| ***(Optional)*** |
| `oneTime` | ***string***| ***(Optional)*** |
## MacieS3BucketAssociationStatus

Appears on:[MacieS3BucketAssociation](#macies3bucketassociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MacieS3BucketAssociationSpec](#macies3bucketassociationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[MacieS3BucketAssociationStatus](#macies3bucketassociationstatus)

---
