---
title: SpannerInstance
menu:
  docs_{{ .version }}:
    identifier: spannerinstance-google.kubeform.com
    name: SpannerInstance
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SpannerInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `SpannerInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SpannerInstanceSpec](#spannerinstancespec)***||
| `status` | ***[SpannerInstanceStatus](#spannerinstancestatus)***||
## Phase(`string` alias)

Appears on:[SpannerInstanceStatus](#spannerinstancestatus)

## SpannerInstanceSpec

Appears on:[SpannerInstance](#spannerinstance), [SpannerInstanceStatus](#spannerinstancestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `config` | ***string***||
| `displayName` | ***string***||
| `labels` | ***map[string]string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `numNodes` | ***int64***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `state` | ***string***| ***(Optional)*** |
## SpannerInstanceStatus

Appears on:[SpannerInstance](#spannerinstance)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SpannerInstanceSpec](#spannerinstancespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
