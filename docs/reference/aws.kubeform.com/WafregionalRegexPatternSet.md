---
title: WafregionalRegexPatternSet
menu:
  docs_{{ .version }}:
    identifier: wafregionalregexpatternset-aws.kubeform.com
    name: WafregionalRegexPatternSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafregionalRegexPatternSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafregionalRegexPatternSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalRegexPatternSetSpec](#wafregionalregexpatternsetspec)***||
| `status` | ***[WafregionalRegexPatternSetStatus](#wafregionalregexpatternsetstatus)***||
## Phase(`string` alias)

Appears on:[WafregionalRegexPatternSetStatus](#wafregionalregexpatternsetstatus)

## WafregionalRegexPatternSetSpec

Appears on:[WafregionalRegexPatternSet](#wafregionalregexpatternset), [WafregionalRegexPatternSetStatus](#wafregionalregexpatternsetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `regexPatternStrings` | ***[]string***| ***(Optional)*** |
## WafregionalRegexPatternSetStatus

Appears on:[WafregionalRegexPatternSet](#wafregionalregexpatternset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalRegexPatternSetSpec](#wafregionalregexpatternsetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
