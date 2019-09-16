---
title: WafregionalRuleGroup
menu:
  docs_{{ .version }}:
    identifier: wafregionalrulegroup-aws.kubeform.com
    name: WafregionalRuleGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafregionalRuleGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafregionalRuleGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalRuleGroupSpec](#wafregionalrulegroupspec)***||
| `status` | ***[WafregionalRuleGroupStatus](#wafregionalrulegroupstatus)***||
## WafregionalRuleGroupSpec

Appears on:[WafregionalRuleGroup](#wafregionalrulegroup), [WafregionalRuleGroupStatus](#wafregionalrulegroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `activatedRule` | ***[[]WafregionalRuleGroupSpecActivatedRule](#wafregionalrulegroupspecactivatedrule)***| ***(Optional)*** |
| `metricName` | ***string***||
| `name` | ***string***||
## WafregionalRuleGroupSpecActivatedRule

Appears on:[WafregionalRuleGroupSpec](#wafregionalrulegroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***[[]WafregionalRuleGroupSpecActivatedRuleAction](#wafregionalrulegroupspecactivatedruleaction)***||
| `priority` | ***int***||
| `ruleID` | ***string***||
| `type` | ***string***| ***(Optional)*** |
## WafregionalRuleGroupSpecActivatedRuleAction

Appears on:[WafregionalRuleGroupSpecActivatedRule](#wafregionalrulegroupspecactivatedrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## WafregionalRuleGroupStatus

Appears on:[WafregionalRuleGroup](#wafregionalrulegroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalRuleGroupSpec](#wafregionalrulegroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
