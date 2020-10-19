---
title: PrivateDNSMxRecord
menu:
  docs_{{ .version }}:
    identifier: privatednsmxrecord-azurerm.kubeform.com
    name: PrivateDNSMxRecord
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PrivateDNSMxRecord
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PrivateDNSMxRecord` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PrivateDNSMxRecordSpec](#privatednsmxrecordspec)***||
| `status` | ***[PrivateDNSMxRecordStatus](#privatednsmxrecordstatus)***||
## Phase(`string` alias)

Appears on:[PrivateDNSMxRecordStatus](#privatednsmxrecordstatus)

## PrivateDNSMxRecordSpec

Appears on:[PrivateDNSMxRecord](#privatednsmxrecord), [PrivateDNSMxRecordStatus](#privatednsmxrecordstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `record` | ***[[]PrivateDNSMxRecordSpecRecord](#privatednsmxrecordspecrecord)***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `ttl` | ***int64***||
| `zoneName` | ***string***||
## PrivateDNSMxRecordSpecRecord

Appears on:[PrivateDNSMxRecordSpec](#privatednsmxrecordspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `exchange` | ***string***||
| `preference` | ***int64***||
## PrivateDNSMxRecordStatus

Appears on:[PrivateDNSMxRecord](#privatednsmxrecord)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PrivateDNSMxRecordSpec](#privatednsmxrecordspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
