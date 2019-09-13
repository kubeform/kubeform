---
title: WafRule
menu:
  docs_{{ .version }}:
    identifier: wafrule-aws.kubeform.com
    name: WafRule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafRuleSpec](#WafRuleSpec)***||
| `status` | ***[WafRuleStatus](#WafRuleStatus)***||
## WafRuleSpec
##### (Appears on:[WafRule](#WafRule), [WafRuleStatus](#WafRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `metricName` | ***string***||
| `name` | ***string***||
| `predicates` | ***[[]WafRuleSpecPredicates](#WafRuleSpecPredicates)***| ***(Optional)*** |
## WafRuleSpecPredicates
##### (Appears on:[WafRuleSpec](#WafRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `dataID` | ***string***||
| `negated` | ***bool***||
| `type` | ***string***||
## WafRuleStatus
##### (Appears on:[WafRule](#WafRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafRuleSpec](#WafRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
