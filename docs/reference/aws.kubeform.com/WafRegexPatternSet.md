---
title: WafRegexPatternSet
menu:
  docs_{{ .version }}:
    identifier: wafregexpatternset-aws.kubeform.com
    name: WafRegexPatternSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafRegexPatternSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafRegexPatternSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafRegexPatternSetSpec](#WafRegexPatternSetSpec)***||
| `status` | ***[WafRegexPatternSetStatus](#WafRegexPatternSetStatus)***||
## WafRegexPatternSetSpec

Appears on:[WafRegexPatternSet](#WafRegexPatternSet), [WafRegexPatternSetStatus](#WafRegexPatternSetStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `regexPatternStrings` | ***[]string***| ***(Optional)*** |
## WafRegexPatternSetStatus

Appears on:[WafRegexPatternSet](#WafRegexPatternSet)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafRegexPatternSetSpec](#WafRegexPatternSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
