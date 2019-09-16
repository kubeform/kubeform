---
title: WafWebACL
menu:
  docs_{{ .version }}:
    identifier: wafwebacl-aws.kubeform.com
    name: WafWebACL
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafWebACL
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafWebACL` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafWebACLSpec](#WafWebACLSpec)***||
| `status` | ***[WafWebACLStatus](#WafWebACLStatus)***||
## WafWebACLSpec

Appears on:[WafWebACL](#WafWebACL), [WafWebACLStatus](#WafWebACLStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `defaultAction` | ***[[]WafWebACLSpecDefaultAction](#WafWebACLSpecDefaultAction)***||
| `loggingConfiguration` | ***[[]WafWebACLSpecLoggingConfiguration](#WafWebACLSpecLoggingConfiguration)***| ***(Optional)*** |
| `metricName` | ***string***||
| `name` | ***string***||
| `rules` | ***[[]WafWebACLSpecRules](#WafWebACLSpecRules)***| ***(Optional)*** |
## WafWebACLSpecDefaultAction

Appears on:[WafWebACLSpec](#WafWebACLSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## WafWebACLSpecLoggingConfiguration

Appears on:[WafWebACLSpec](#WafWebACLSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `logDestination` | ***string***||
| `redactedFields` | ***[[]WafWebACLSpecLoggingConfigurationRedactedFields](#WafWebACLSpecLoggingConfigurationRedactedFields)***| ***(Optional)*** |
## WafWebACLSpecLoggingConfigurationRedactedFields

Appears on:[WafWebACLSpecLoggingConfiguration](#WafWebACLSpecLoggingConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `fieldToMatch` | ***[[]WafWebACLSpecLoggingConfigurationRedactedFieldsFieldToMatch](#WafWebACLSpecLoggingConfigurationRedactedFieldsFieldToMatch)***||
## WafWebACLSpecLoggingConfigurationRedactedFieldsFieldToMatch

Appears on:[WafWebACLSpecLoggingConfigurationRedactedFields](#WafWebACLSpecLoggingConfigurationRedactedFields)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafWebACLSpecRules

Appears on:[WafWebACLSpec](#WafWebACLSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***[[]WafWebACLSpecRulesAction](#WafWebACLSpecRulesAction)***| ***(Optional)*** |
| `overrideAction` | ***[[]WafWebACLSpecRulesOverrideAction](#WafWebACLSpecRulesOverrideAction)***| ***(Optional)*** |
| `priority` | ***int***||
| `ruleID` | ***string***||
| `type` | ***string***| ***(Optional)*** |
## WafWebACLSpecRulesAction

Appears on:[WafWebACLSpecRules](#WafWebACLSpecRules)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## WafWebACLSpecRulesOverrideAction

Appears on:[WafWebACLSpecRules](#WafWebACLSpecRules)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## WafWebACLStatus

Appears on:[WafWebACL](#WafWebACL)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafWebACLSpec](#WafWebACLSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
