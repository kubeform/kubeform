---
title: SccSource
menu:
  docs_{{ .version }}:
    identifier: sccsource-google.kubeform.com
    name: SccSource
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SccSource
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `SccSource` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SccSourceSpec](#sccsourcespec)***||
| `status` | ***[SccSourceStatus](#sccsourcestatus)***||
## Phase(`string` alias)

Appears on:[SccSourceStatus](#sccsourcestatus)

## SccSourceSpec

Appears on:[SccSource](#sccsource), [SccSourceStatus](#sccsourcestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `displayName` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `organization` | ***string***||
## SccSourceStatus

Appears on:[SccSource](#sccsource)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SccSourceSpec](#sccsourcespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
