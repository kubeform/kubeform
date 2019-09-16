---
title: WafRateBasedRule
menu:
  docs_{{ .version }}:
    identifier: wafratebasedrule-aws.kubeform.com
    name: WafRateBasedRule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafRateBasedRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafRateBasedRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafRateBasedRuleSpec](#WafRateBasedRuleSpec)***||
| `status` | ***[WafRateBasedRuleStatus](#WafRateBasedRuleStatus)***||
## WafRateBasedRuleSpec

Appears on:[WafRateBasedRule](#WafRateBasedRule), [WafRateBasedRuleStatus](#WafRateBasedRuleStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `metricName` | ***string***||
| `name` | ***string***||
| `predicates` | ***[[]WafRateBasedRuleSpecPredicates](#WafRateBasedRuleSpecPredicates)***| ***(Optional)*** |
| `rateKey` | ***string***||
| `rateLimit` | ***int***||
## WafRateBasedRuleSpecPredicates

Appears on:[WafRateBasedRuleSpec](#WafRateBasedRuleSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dataID` | ***string***||
| `negated` | ***bool***||
| `type` | ***string***||
## WafRateBasedRuleStatus

Appears on:[WafRateBasedRule](#WafRateBasedRule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafRateBasedRuleSpec](#WafRateBasedRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
