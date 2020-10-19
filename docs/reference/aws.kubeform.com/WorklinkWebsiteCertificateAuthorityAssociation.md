---
title: WorklinkWebsiteCertificateAuthorityAssociation
menu:
  docs_{{ .version }}:
    identifier: worklinkwebsitecertificateauthorityassociation-aws.kubeform.com
    name: WorklinkWebsiteCertificateAuthorityAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WorklinkWebsiteCertificateAuthorityAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WorklinkWebsiteCertificateAuthorityAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WorklinkWebsiteCertificateAuthorityAssociationSpec](#worklinkwebsitecertificateauthorityassociationspec)***||
| `status` | ***[WorklinkWebsiteCertificateAuthorityAssociationStatus](#worklinkwebsitecertificateauthorityassociationstatus)***||
## Phase(`string` alias)

Appears on:[WorklinkWebsiteCertificateAuthorityAssociationStatus](#worklinkwebsitecertificateauthorityassociationstatus)

## WorklinkWebsiteCertificateAuthorityAssociationSpec

Appears on:[WorklinkWebsiteCertificateAuthorityAssociation](#worklinkwebsitecertificateauthorityassociation), [WorklinkWebsiteCertificateAuthorityAssociationStatus](#worklinkwebsitecertificateauthorityassociationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `certificate` | ***string***||
| `displayName` | ***string***| ***(Optional)*** |
| `fleetArn` | ***string***||
| `websiteCaID` | ***string***| ***(Optional)*** |
## WorklinkWebsiteCertificateAuthorityAssociationStatus

Appears on:[WorklinkWebsiteCertificateAuthorityAssociation](#worklinkwebsitecertificateauthorityassociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WorklinkWebsiteCertificateAuthorityAssociationSpec](#worklinkwebsitecertificateauthorityassociationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
