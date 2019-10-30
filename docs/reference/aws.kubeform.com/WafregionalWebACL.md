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
| `spec` | ***[WafregionalWebACLSpec](#wafregionalwebaclspec)***||
| `status` | ***[WafregionalWebACLStatus](#wafregionalwebaclstatus)***||
## Phase(`string` alias)

Appears on:[WafregionalWebACLStatus](#wafregionalwebaclstatus)

## WafregionalWebACLSpec

Appears on:[WafregionalWebACL](#wafregionalwebacl), [WafregionalWebACLStatus](#wafregionalwebaclstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `defaultAction` | ***[[]WafregionalWebACLSpecDefaultAction](#wafregionalwebaclspecdefaultaction)***||
| `loggingConfiguration` | ***[[]WafregionalWebACLSpecLoggingConfiguration](#wafregionalwebaclspecloggingconfiguration)***| ***(Optional)*** |
| `metricName` | ***string***||
| `name` | ***string***||
| `rule` | ***[[]WafregionalWebACLSpecRule](#wafregionalwebaclspecrule)***| ***(Optional)*** |
## WafregionalWebACLSpecDefaultAction

Appears on:[WafregionalWebACLSpec](#wafregionalwebaclspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## WafregionalWebACLSpecLoggingConfiguration

Appears on:[WafregionalWebACLSpec](#wafregionalwebaclspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `logDestination` | ***string***||
| `redactedFields` | ***[[]WafregionalWebACLSpecLoggingConfigurationRedactedFields](#wafregionalwebaclspecloggingconfigurationredactedfields)***| ***(Optional)*** |
## WafregionalWebACLSpecLoggingConfigurationRedactedFields

Appears on:[WafregionalWebACLSpecLoggingConfiguration](#wafregionalwebaclspecloggingconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `fieldToMatch` | ***[[]WafregionalWebACLSpecLoggingConfigurationRedactedFieldsFieldToMatch](#wafregionalwebaclspecloggingconfigurationredactedfieldsfieldtomatch)***||
## WafregionalWebACLSpecLoggingConfigurationRedactedFieldsFieldToMatch

Appears on:[WafregionalWebACLSpecLoggingConfigurationRedactedFields](#wafregionalwebaclspecloggingconfigurationredactedfields)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafregionalWebACLSpecRule

Appears on:[WafregionalWebACLSpec](#wafregionalwebaclspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***[[]WafregionalWebACLSpecRuleAction](#wafregionalwebaclspecruleaction)***| ***(Optional)*** |
| `overrideAction` | ***[[]WafregionalWebACLSpecRuleOverrideAction](#wafregionalwebaclspecruleoverrideaction)***| ***(Optional)*** |
| `priority` | ***int***||
| `ruleID` | ***string***||
| `type` | ***string***| ***(Optional)*** |
## WafregionalWebACLSpecRuleAction

Appears on:[WafregionalWebACLSpecRule](#wafregionalwebaclspecrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## WafregionalWebACLSpecRuleOverrideAction

Appears on:[WafregionalWebACLSpecRule](#wafregionalwebaclspecrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## WafregionalWebACLStatus

Appears on:[WafregionalWebACL](#wafregionalwebacl)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalWebACLSpec](#wafregionalwebaclspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
