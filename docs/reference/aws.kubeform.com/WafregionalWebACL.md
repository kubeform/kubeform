---
title: WafregionalWebACL
menu:
  docs_{{ .version }}:
    identifier: wafregionalwebacl-aws.kubeform.com
    name: WafregionalWebACL
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafregionalWebACL
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafregionalWebACL` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalWebACLSpec](#WafregionalWebACLSpec)***||
| `status` | ***[WafregionalWebACLStatus](#WafregionalWebACLStatus)***||
## WafregionalWebACLSpec
##### (Appears on:[WafregionalWebACL](#WafregionalWebACL), [WafregionalWebACLStatus](#WafregionalWebACLStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `defaultAction` | ***[[]WafregionalWebACLSpecDefaultAction](#WafregionalWebACLSpecDefaultAction)***||
| `loggingConfiguration` | ***[[]WafregionalWebACLSpecLoggingConfiguration](#WafregionalWebACLSpecLoggingConfiguration)***| ***(Optional)*** |
| `metricName` | ***string***||
| `name` | ***string***||
| `rule` | ***[[]WafregionalWebACLSpecRule](#WafregionalWebACLSpecRule)***| ***(Optional)*** |
## WafregionalWebACLSpecDefaultAction
##### (Appears on:[WafregionalWebACLSpec](#WafregionalWebACLSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## WafregionalWebACLSpecLoggingConfiguration
##### (Appears on:[WafregionalWebACLSpec](#WafregionalWebACLSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `logDestination` | ***string***||
| `redactedFields` | ***[[]WafregionalWebACLSpecLoggingConfigurationRedactedFields](#WafregionalWebACLSpecLoggingConfigurationRedactedFields)***| ***(Optional)*** |
## WafregionalWebACLSpecLoggingConfigurationRedactedFields
##### (Appears on:[WafregionalWebACLSpecLoggingConfiguration](#WafregionalWebACLSpecLoggingConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `fieldToMatch` | ***[[]WafregionalWebACLSpecLoggingConfigurationRedactedFieldsFieldToMatch](#WafregionalWebACLSpecLoggingConfigurationRedactedFieldsFieldToMatch)***||
## WafregionalWebACLSpecLoggingConfigurationRedactedFieldsFieldToMatch
##### (Appears on:[WafregionalWebACLSpecLoggingConfigurationRedactedFields](#WafregionalWebACLSpecLoggingConfigurationRedactedFields))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafregionalWebACLSpecRule
##### (Appears on:[WafregionalWebACLSpec](#WafregionalWebACLSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***[[]WafregionalWebACLSpecRuleAction](#WafregionalWebACLSpecRuleAction)***| ***(Optional)*** |
| `overrideAction` | ***[[]WafregionalWebACLSpecRuleOverrideAction](#WafregionalWebACLSpecRuleOverrideAction)***| ***(Optional)*** |
| `priority` | ***int***||
| `ruleID` | ***string***||
| `type` | ***string***| ***(Optional)*** |
## WafregionalWebACLSpecRuleAction
##### (Appears on:[WafregionalWebACLSpecRule](#WafregionalWebACLSpecRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## WafregionalWebACLSpecRuleOverrideAction
##### (Appears on:[WafregionalWebACLSpecRule](#WafregionalWebACLSpecRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## WafregionalWebACLStatus
##### (Appears on:[WafregionalWebACL](#WafregionalWebACL))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalWebACLSpec](#WafregionalWebACLSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
