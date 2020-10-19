---
title: DnsTxtRecord
menu:
  docs_{{ .version }}:
    identifier: dnstxtrecord-azurerm.kubeform.com
    name: DnsTxtRecord
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DnsTxtRecord
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DnsTxtRecord` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DnsTxtRecordSpec](#dnstxtrecordspec)***||
| `status` | ***[DnsTxtRecordStatus](#dnstxtrecordstatus)***||
## DnsTxtRecordSpec

Appears on:[DnsTxtRecord](#dnstxtrecord), [DnsTxtRecordStatus](#dnstxtrecordstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `fqdn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `record` | ***[[]DnsTxtRecordSpecRecord](#dnstxtrecordspecrecord)***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `ttl` | ***int64***||
| `zoneName` | ***string***||
## DnsTxtRecordSpecRecord

Appears on:[DnsTxtRecordSpec](#dnstxtrecordspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `value` | ***string***||
## DnsTxtRecordStatus

Appears on:[DnsTxtRecord](#dnstxtrecord)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DnsTxtRecordSpec](#dnstxtrecordspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DnsTxtRecordStatus](#dnstxtrecordstatus)

---
