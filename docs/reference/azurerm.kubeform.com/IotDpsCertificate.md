---
title: IotDpsCertificate
menu:
  docs_{{ .version }}:
    identifier: iotdpscertificate-azurerm.kubeform.com
    name: IotDpsCertificate
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IotDpsCertificate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `IotDpsCertificate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IotDpsCertificateSpec](#iotdpscertificatespec)***||
| `status` | ***[IotDpsCertificateStatus](#iotdpscertificatestatus)***||
## IotDpsCertificateSpec

Appears on:[IotDpsCertificate](#iotdpscertificate), [IotDpsCertificateStatus](#iotdpscertificatestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `iotDpsName` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
## IotDpsCertificateStatus

Appears on:[IotDpsCertificate](#iotdpscertificate)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IotDpsCertificateSpec](#iotdpscertificatespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IotDpsCertificateStatus](#iotdpscertificatestatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `certificate_content` | ***string*** ||
