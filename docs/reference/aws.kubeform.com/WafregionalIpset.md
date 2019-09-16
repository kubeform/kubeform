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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalIpsetSpec](#WafregionalIpsetSpec)***||
| `status` | ***[WafregionalIpsetStatus](#WafregionalIpsetStatus)***||
## WafregionalIpsetSpec

Appears on:[WafregionalIpset](#WafregionalIpset), [WafregionalIpsetStatus](#WafregionalIpsetStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `ipSetDescriptor` | ***[[]WafregionalIpsetSpecIpSetDescriptor](#WafregionalIpsetSpecIpSetDescriptor)***| ***(Optional)*** |
| `name` | ***string***||
## WafregionalIpsetSpecIpSetDescriptor

Appears on:[WafregionalIpsetSpec](#WafregionalIpsetSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
| `value` | ***string***||
## WafregionalIpsetStatus

Appears on:[WafregionalIpset](#WafregionalIpset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalIpsetSpec](#WafregionalIpsetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
