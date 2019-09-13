---
title: ApiManagementCertificate
menu:
  docs_{{ .version }}:
    identifier: apimanagementcertificate-azurerm.kubeform.com
    name: ApiManagementCertificate
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiManagementCertificate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementCertificate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementCertificateSpec](#ApiManagementCertificateSpec)***||
| `status` | ***[ApiManagementCertificateStatus](#ApiManagementCertificateStatus)***||
## ApiManagementCertificateSpec
##### (Appears on:[ApiManagementCertificate](#ApiManagementCertificate), [ApiManagementCertificateStatus](#ApiManagementCertificateStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `apiManagementName` | ***string***||
| `expiration` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `subject` | ***string***| ***(Optional)*** |
| `thumbprint` | ***string***| ***(Optional)*** |
## ApiManagementCertificateStatus
##### (Appears on:[ApiManagementCertificate](#ApiManagementCertificate))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementCertificateSpec](#ApiManagementCertificateSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `data` | ***string*** ||
| `password` | ***string*** ||
