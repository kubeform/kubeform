---
title: ConfigConfigRule
menu:
  docs_{{ .version }}:
    identifier: configconfigrule-aws.kubeform.com
    name: ConfigConfigRule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ConfigConfigRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ConfigConfigRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ConfigConfigRuleSpec](#configconfigrulespec)***||
| `status` | ***[ConfigConfigRuleStatus](#configconfigrulestatus)***||
## ConfigConfigRuleSpec

Appears on:[ConfigConfigRule](#configconfigrule), [ConfigConfigRuleStatus](#configconfigrulestatus)

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
| `scope` | ***[[]ConfigConfigRuleSpecScope](#configconfigrulespecscope)***| ***(Optional)*** |
| `source` | ***[[]ConfigConfigRuleSpecSource](#configconfigrulespecsource)***||
## ConfigConfigRuleSpecScope

Appears on:[ConfigConfigRuleSpec](#configconfigrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `complianceResourceID` | ***string***| ***(Optional)*** |
| `complianceResourceTypes` | ***[]string***| ***(Optional)*** |
| `tagKey` | ***string***| ***(Optional)*** |
| `tagValue` | ***string***| ***(Optional)*** |
## ConfigConfigRuleSpecSource

Appears on:[ConfigConfigRuleSpec](#configconfigrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `owner` | ***string***||
| `sourceDetail` | ***[[]ConfigConfigRuleSpecSourceSourceDetail](#configconfigrulespecsourcesourcedetail)***| ***(Optional)*** |
| `sourceIdentifier` | ***string***||
## ConfigConfigRuleSpecSourceSourceDetail

Appears on:[ConfigConfigRuleSpecSource](#configconfigrulespecsource)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `eventSource` | ***string***| ***(Optional)*** |
| `maximumExecutionFrequency` | ***string***| ***(Optional)*** |
| `messageType` | ***string***| ***(Optional)*** |
## ConfigConfigRuleStatus

Appears on:[ConfigConfigRule](#configconfigrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ConfigConfigRuleSpec](#configconfigrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
