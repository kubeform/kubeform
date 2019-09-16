---
title: WafRuleGroup
menu:
  docs_{{ .version }}:
    identifier: wafrulegroup-aws.kubeform.com
    name: WafRuleGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafRuleGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafRuleGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafRuleGroupSpec](#WafRuleGroupSpec)***||
| `status` | ***[WafRuleGroupStatus](#WafRuleGroupStatus)***||
## WafRuleGroupSpec

Appears on:[WafRuleGroup](#WafRuleGroup), [WafRuleGroupStatus](#WafRuleGroupStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `activatedRule` | ***[[]WafRuleGroupSpecActivatedRule](#WafRuleGroupSpecActivatedRule)***| ***(Optional)*** |
| `metricName` | ***string***||
| `name` | ***string***||
## WafRuleGroupSpecActivatedRule

Appears on:[WafRuleGroupSpec](#WafRuleGroupSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***[[]WafRuleGroupSpecActivatedRuleAction](#WafRuleGroupSpecActivatedRuleAction)***||
| `priority` | ***int***||
| `ruleID` | ***string***||
| `type` | ***string***| ***(Optional)*** |
## WafRuleGroupSpecActivatedRuleAction

Appears on:[WafRuleGroupSpecActivatedRule](#WafRuleGroupSpecActivatedRule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## WafRuleGroupStatus

Appears on:[WafRuleGroup](#WafRuleGroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafRuleGroupSpec](#WafRuleGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
