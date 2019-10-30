---
title: PrivateDNSARecord
menu:
  docs_{{ .version }}:
    identifier: privatednsarecord-azurerm.kubeform.com
    name: PrivateDNSARecord
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PrivateDNSARecord
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PrivateDNSARecord` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PrivateDNSARecordSpec](#privatednsarecordspec)***||
| `status` | ***[PrivateDNSARecordStatus](#privatednsarecordstatus)***||
## Phase(`string` alias)

Appears on:[PrivateDNSARecordStatus](#privatednsarecordstatus)

## PrivateDNSARecordSpec

Appears on:[PrivateDNSARecord](#privatednsarecord), [PrivateDNSARecordStatus](#privatednsarecordstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `records` | ***[]string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `ttl` | ***int***||
| `zoneName` | ***string***||
## PrivateDNSARecordStatus

Appears on:[PrivateDNSARecord](#privatednsarecord)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PrivateDNSARecordSpec](#privatednsarecordspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
