---
title: WafregionalRateBasedRule
menu:
  docs_v0.0.1:
    identifier: wafregionalratebasedrule-aws.kubeform.com
    name: WafregionalRateBasedRule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## WafregionalRateBasedRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafregionalRateBasedRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalRateBasedRuleSpec](#WafregionalRateBasedRuleSpec)***||
| `status` | ***[WafregionalRateBasedRuleStatus](#WafregionalRateBasedRuleStatus)***||
## WafregionalRateBasedRuleSpec
##### (Appears on:[WafregionalRateBasedRule](#WafregionalRateBasedRule), [WafregionalRateBasedRuleStatus](#WafregionalRateBasedRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `metricName` | ***string***||
| `name` | ***string***||
| `predicate` | ***[[]WafregionalRateBasedRuleSpecPredicate](#WafregionalRateBasedRuleSpecPredicate)***| ***(Optional)*** |
| `rateKey` | ***string***||
| `rateLimit` | ***int***||
## WafregionalRateBasedRuleSpecPredicate
##### (Appears on:[WafregionalRateBasedRuleSpec](#WafregionalRateBasedRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `dataID` | ***string***||
| `negated` | ***bool***||
| `type` | ***string***||
## WafregionalRateBasedRuleStatus
##### (Appears on:[WafregionalRateBasedRule](#WafregionalRateBasedRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalRateBasedRuleSpec](#WafregionalRateBasedRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
