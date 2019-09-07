---
title: BatchCertificate
menu:
  docs_v0.0.1:
    identifier: batchcertificate-azurerm.kubeform.com
    name: BatchCertificate
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## BatchCertificate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `BatchCertificate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BatchCertificateSpec](#BatchCertificateSpec)***||
| `status` | ***[BatchCertificateStatus](#BatchCertificateStatus)***||
## BatchCertificateSpec
##### (Appears on:[BatchCertificate](#BatchCertificate), [BatchCertificateStatus](#BatchCertificateStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `accountName` | ***string***||
| `format` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `publicData` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `thumbprint` | ***string***||
| `thumbprintAlgorithm` | ***string***||
## BatchCertificateStatus
##### (Appears on:[BatchCertificate](#BatchCertificate))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BatchCertificateSpec](#BatchCertificateSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `certificate` | ***string*** ||
| `password` | ***string*** ||
