---
title: PrivateDNSPtrRecord
menu:
  docs_{{ .version }}:
    identifier: privatednsptrrecord-azurerm.kubeform.com
    name: PrivateDNSPtrRecord
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PrivateDNSPtrRecord
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PrivateDNSPtrRecord` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PrivateDNSPtrRecordSpec](#privatednsptrrecordspec)***||
| `status` | ***[PrivateDNSPtrRecordStatus](#privatednsptrrecordstatus)***||
## Phase(`string` alias)

Appears on:[PrivateDNSPtrRecordStatus](#privatednsptrrecordstatus)

## PrivateDNSPtrRecordSpec

Appears on:[PrivateDNSPtrRecord](#privatednsptrrecord), [PrivateDNSPtrRecordStatus](#privatednsptrrecordstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `records` | ***[]string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `ttl` | ***int64***||
| `zoneName` | ***string***||
## PrivateDNSPtrRecordStatus

Appears on:[PrivateDNSPtrRecord](#privatednsptrrecord)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PrivateDNSPtrRecordSpec](#privatednsptrrecordspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
