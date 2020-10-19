---
title: PrivateDNSCnameRecord
menu:
  docs_{{ .version }}:
    identifier: privatednscnamerecord-azurerm.kubeform.com
    name: PrivateDNSCnameRecord
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PrivateDNSCnameRecord
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PrivateDNSCnameRecord` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PrivateDNSCnameRecordSpec](#privatednscnamerecordspec)***||
| `status` | ***[PrivateDNSCnameRecordStatus](#privatednscnamerecordstatus)***||
## Phase(`string` alias)

Appears on:[PrivateDNSCnameRecordStatus](#privatednscnamerecordstatus)

## PrivateDNSCnameRecordSpec

Appears on:[PrivateDNSCnameRecord](#privatednscnamerecord), [PrivateDNSCnameRecordStatus](#privatednscnamerecordstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `record` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `ttl` | ***int64***||
| `zoneName` | ***string***||
## PrivateDNSCnameRecordStatus

Appears on:[PrivateDNSCnameRecord](#privatednscnamerecord)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PrivateDNSCnameRecordSpec](#privatednscnamerecordspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
