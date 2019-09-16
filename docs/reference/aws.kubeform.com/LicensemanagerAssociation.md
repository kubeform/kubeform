---
title: LicensemanagerAssociation
menu:
  docs_{{ .version }}:
    identifier: licensemanagerassociation-aws.kubeform.com
    name: LicensemanagerAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LicensemanagerAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LicensemanagerAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LicensemanagerAssociationSpec](#LicensemanagerAssociationSpec)***||
| `status` | ***[LicensemanagerAssociationStatus](#LicensemanagerAssociationStatus)***||
## LicensemanagerAssociationSpec

Appears on:[LicensemanagerAssociation](#LicensemanagerAssociation), [LicensemanagerAssociationStatus](#LicensemanagerAssociationStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `licenseConfigurationArn` | ***string***||
| `resourceArn` | ***string***||
## LicensemanagerAssociationStatus

Appears on:[LicensemanagerAssociation](#LicensemanagerAssociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LicensemanagerAssociationSpec](#LicensemanagerAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
