---
title: AutomationCertificate
menu:
  docs_{{ .version }}:
    identifier: automationcertificate-azurerm.kubeform.com
    name: AutomationCertificate
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AutomationCertificate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AutomationCertificate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutomationCertificateSpec](#automationcertificatespec)***||
| `status` | ***[AutomationCertificateStatus](#automationcertificatestatus)***||
## AutomationCertificateSpec

Appears on:[AutomationCertificate](#automationcertificate), [AutomationCertificateStatus](#automationcertificatestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `automationAccountName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `exportable` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `thumbprint` | ***string***| ***(Optional)*** |
## AutomationCertificateStatus

Appears on:[AutomationCertificate](#automationcertificate)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutomationCertificateSpec](#automationcertificatespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AutomationCertificateStatus](#automationcertificatestatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `base64` | ***string*** ||
