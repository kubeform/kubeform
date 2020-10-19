---
title: WafregionalIpset
menu:
  docs_{{ .version }}:
    identifier: wafregionalipset-aws.kubeform.com
    name: WafregionalIpset
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafregionalIpset
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafregionalIpset` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalIpsetSpec](#wafregionalipsetspec)***||
| `status` | ***[WafregionalIpsetStatus](#wafregionalipsetstatus)***||
## Phase(`string` alias)

Appears on:[WafregionalIpsetStatus](#wafregionalipsetstatus)

## WafregionalIpsetSpec

Appears on:[WafregionalIpset](#wafregionalipset), [WafregionalIpsetStatus](#wafregionalipsetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `ipSetDescriptor` | ***[[]WafregionalIpsetSpecIpSetDescriptor](#wafregionalipsetspecipsetdescriptor)***| ***(Optional)*** |
| `name` | ***string***||
## WafregionalIpsetSpecIpSetDescriptor

Appears on:[WafregionalIpsetSpec](#wafregionalipsetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
| `value` | ***string***||
## WafregionalIpsetStatus

Appears on:[WafregionalIpset](#wafregionalipset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalIpsetSpec](#wafregionalipsetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
