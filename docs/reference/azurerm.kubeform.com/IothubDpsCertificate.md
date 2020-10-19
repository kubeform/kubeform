---
title: IothubDpsCertificate
menu:
  docs_{{ .version }}:
    identifier: iothubdpscertificate-azurerm.kubeform.com
    name: IothubDpsCertificate
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IothubDpsCertificate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `IothubDpsCertificate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IothubDpsCertificateSpec](#iothubdpscertificatespec)***||
| `status` | ***[IothubDpsCertificateStatus](#iothubdpscertificatestatus)***||
## IothubDpsCertificateSpec

Appears on:[IothubDpsCertificate](#iothubdpscertificate), [IothubDpsCertificateStatus](#iothubdpscertificatestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `iotDpsName` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
## IothubDpsCertificateStatus

Appears on:[IothubDpsCertificate](#iothubdpscertificate)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IothubDpsCertificateSpec](#iothubdpscertificatespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IothubDpsCertificateStatus](#iothubdpscertificatestatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `certificate_content` | ***string*** ||
