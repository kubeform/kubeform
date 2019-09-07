---
title: WorklinkWebsiteCertificateAuthorityAssociation
menu:
  docs_v0.0.1:
    identifier: worklinkwebsitecertificateauthorityassociation-aws.kubeform.com
    name: WorklinkWebsiteCertificateAuthorityAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## WorklinkWebsiteCertificateAuthorityAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WorklinkWebsiteCertificateAuthorityAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WorklinkWebsiteCertificateAuthorityAssociationSpec](#WorklinkWebsiteCertificateAuthorityAssociationSpec)***||
| `status` | ***[WorklinkWebsiteCertificateAuthorityAssociationStatus](#WorklinkWebsiteCertificateAuthorityAssociationStatus)***||
## WorklinkWebsiteCertificateAuthorityAssociationSpec
##### (Appears on:[WorklinkWebsiteCertificateAuthorityAssociation](#WorklinkWebsiteCertificateAuthorityAssociation), [WorklinkWebsiteCertificateAuthorityAssociationStatus](#WorklinkWebsiteCertificateAuthorityAssociationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `certificate` | ***string***||
| `displayName` | ***string***| ***(Optional)*** |
| `fleetArn` | ***string***||
| `websiteCaID` | ***string***| ***(Optional)*** |
## WorklinkWebsiteCertificateAuthorityAssociationStatus
##### (Appears on:[WorklinkWebsiteCertificateAuthorityAssociation](#WorklinkWebsiteCertificateAuthorityAssociation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WorklinkWebsiteCertificateAuthorityAssociationSpec](#WorklinkWebsiteCertificateAuthorityAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
