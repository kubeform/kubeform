---
title: ConfigOrganizationCustomRule
menu:
  docs_{{ .version }}:
    identifier: configorganizationcustomrule-aws.kubeform.com
    name: ConfigOrganizationCustomRule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ConfigOrganizationCustomRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ConfigOrganizationCustomRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ConfigOrganizationCustomRuleSpec](#configorganizationcustomrulespec)***||
| `status` | ***[ConfigOrganizationCustomRuleStatus](#configorganizationcustomrulestatus)***||
## ConfigOrganizationCustomRuleSpec

Appears on:[ConfigOrganizationCustomRule](#configorganizationcustomrule), [ConfigOrganizationCustomRuleStatus](#configorganizationcustomrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `excludedAccounts` | ***[]string***| ***(Optional)*** |
| `inputParameters` | ***string***| ***(Optional)*** |
| `lambdaFunctionArn` | ***string***||
| `maximumExecutionFrequency` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `resourceIDScope` | ***string***| ***(Optional)*** |
| `resourceTypesScope` | ***[]string***| ***(Optional)*** |
| `tagKeyScope` | ***string***| ***(Optional)*** |
| `tagValueScope` | ***string***| ***(Optional)*** |
| `triggerTypes` | ***[]string***||
## ConfigOrganizationCustomRuleStatus

Appears on:[ConfigOrganizationCustomRule](#configorganizationcustomrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ConfigOrganizationCustomRuleSpec](#configorganizationcustomrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ConfigOrganizationCustomRuleStatus](#configorganizationcustomrulestatus)

---
