---
title: WafIpset
menu:
  docs_{{ .version }}:
    identifier: wafipset-aws.kubeform.com
    name: WafIpset
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafIpset
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafIpset` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafIpsetSpec](#wafipsetspec)***||
| `status` | ***[WafIpsetStatus](#wafipsetstatus)***||
## Phase(`string` alias)

Appears on:[WafIpsetStatus](#wafipsetstatus)

## WafIpsetSpec

Appears on:[WafIpset](#wafipset), [WafIpsetStatus](#wafipsetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `ipSetDescriptors` | ***[[]WafIpsetSpecIpSetDescriptors](#wafipsetspecipsetdescriptors)***| ***(Optional)*** |
| `name` | ***string***||
## WafIpsetSpecIpSetDescriptors

Appears on:[WafIpsetSpec](#wafipsetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
| `value` | ***string***||
## WafIpsetStatus

Appears on:[WafIpset](#wafipset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafIpsetSpec](#wafipsetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
