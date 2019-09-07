---
title: ConfigConfigRule
menu:
  docs_v0.0.1:
    identifier: configconfigrule-aws.kubeform.com
    name: ConfigConfigRule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ConfigConfigRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ConfigConfigRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ConfigConfigRuleSpec](#ConfigConfigRuleSpec)***||
| `status` | ***[ConfigConfigRuleStatus](#ConfigConfigRuleStatus)***||
## ConfigConfigRuleSpec
##### (Appears on:[ConfigConfigRule](#ConfigConfigRule), [ConfigConfigRuleStatus](#ConfigConfigRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `inputParameters` | ***string***| ***(Optional)*** |
| `maximumExecutionFrequency` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `ruleID` | ***string***| ***(Optional)*** |
| `scope` | ***[[]ConfigConfigRuleSpecScope](#ConfigConfigRuleSpecScope)***| ***(Optional)*** |
| `source` | ***[[]ConfigConfigRuleSpecSource](#ConfigConfigRuleSpecSource)***||
## ConfigConfigRuleSpecScope
##### (Appears on:[ConfigConfigRuleSpec](#ConfigConfigRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `complianceResourceID` | ***string***| ***(Optional)*** |
| `complianceResourceTypes` | ***[]string***| ***(Optional)*** |
| `tagKey` | ***string***| ***(Optional)*** |
| `tagValue` | ***string***| ***(Optional)*** |
## ConfigConfigRuleSpecSource
##### (Appears on:[ConfigConfigRuleSpec](#ConfigConfigRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `owner` | ***string***||
| `sourceDetail` | ***[[]ConfigConfigRuleSpecSourceSourceDetail](#ConfigConfigRuleSpecSourceSourceDetail)***| ***(Optional)*** |
| `sourceIdentifier` | ***string***||
## ConfigConfigRuleSpecSourceSourceDetail
##### (Appears on:[ConfigConfigRuleSpecSource](#ConfigConfigRuleSpecSource))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `eventSource` | ***string***| ***(Optional)*** |
| `maximumExecutionFrequency` | ***string***| ***(Optional)*** |
| `messageType` | ***string***| ***(Optional)*** |
## ConfigConfigRuleStatus
##### (Appears on:[ConfigConfigRule](#ConfigConfigRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ConfigConfigRuleSpec](#ConfigConfigRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
