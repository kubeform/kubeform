---
title: WafregionalRule
menu:
  docs_{{ .version }}:
    identifier: wafregionalrule-aws.kubeform.com
    name: WafregionalRule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafregionalRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafregionalRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalRuleSpec](#WafregionalRuleSpec)***||
| `status` | ***[WafregionalRuleStatus](#WafregionalRuleStatus)***||
## WafregionalRuleSpec
##### (Appears on:[WafregionalRule](#WafregionalRule), [WafregionalRuleStatus](#WafregionalRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `metricName` | ***string***||
| `name` | ***string***||
| `predicate` | ***[[]WafregionalRuleSpecPredicate](#WafregionalRuleSpecPredicate)***| ***(Optional)*** |
## WafregionalRuleSpecPredicate
##### (Appears on:[WafregionalRuleSpec](#WafregionalRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `dataID` | ***string***||
| `negated` | ***bool***||
| `type` | ***string***||
## WafregionalRuleStatus
##### (Appears on:[WafregionalRule](#WafregionalRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalRuleSpec](#WafregionalRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
