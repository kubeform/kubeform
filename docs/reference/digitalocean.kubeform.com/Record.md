---
title: Record
menu:
  docs_{{ .version }}:
    identifier: record-digitalocean.kubeform.com
    name: Record
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Record
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `Record` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RecordSpec](#recordspec)***||
| `status` | ***[RecordStatus](#recordstatus)***||
## Phase(`string` alias)

Appears on:[RecordStatus](#recordstatus)

## RecordSpec

Appears on:[Record](#record), [RecordStatus](#recordstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `domain` | ***string***||
| `flags` | ***int64***| ***(Optional)*** |
| `fqdn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `port` | ***int64***| ***(Optional)*** |
| `priority` | ***int64***| ***(Optional)*** |
| `tag` | ***string***| ***(Optional)*** |
| `ttl` | ***int64***| ***(Optional)*** |
| `type` | ***string***||
| `value` | ***string***||
| `weight` | ***int64***| ***(Optional)*** |
## RecordStatus

Appears on:[Record](#record)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RecordSpec](#recordspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
