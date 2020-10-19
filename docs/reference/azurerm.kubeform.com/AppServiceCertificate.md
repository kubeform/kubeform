---
title: AppServiceCertificate
menu:
  docs_{{ .version }}:
    identifier: appservicecertificate-azurerm.kubeform.com
    name: AppServiceCertificate
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AppServiceCertificate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AppServiceCertificate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppServiceCertificateSpec](#appservicecertificatespec)***||
| `status` | ***[AppServiceCertificateStatus](#appservicecertificatestatus)***||
## AppServiceCertificateSpec

Appears on:[AppServiceCertificate](#appservicecertificate), [AppServiceCertificateStatus](#appservicecertificatestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `expirationDate` | ***string***| ***(Optional)*** |
| `friendlyName` | ***string***| ***(Optional)*** |
| `hostNames` | ***[]string***| ***(Optional)*** |
| `issueDate` | ***string***| ***(Optional)*** |
| `issuer` | ***string***| ***(Optional)*** |
| `keyVaultSecretID` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `subjectName` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `thumbprint` | ***string***| ***(Optional)*** |
## AppServiceCertificateStatus

Appears on:[AppServiceCertificate](#appservicecertificate)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppServiceCertificateSpec](#appservicecertificatespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AppServiceCertificateStatus](#appservicecertificatestatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `password` | ***string*** ||
| `pfx_blob` | ***string*** ||
