---
title: WebApplicationFirewallPolicy
menu:
  docs_{{ .version }}:
    identifier: webapplicationfirewallpolicy-azurerm.kubeform.com
    name: WebApplicationFirewallPolicy
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WebApplicationFirewallPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `WebApplicationFirewallPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WebApplicationFirewallPolicySpec](#webapplicationfirewallpolicyspec)***||
| `status` | ***[WebApplicationFirewallPolicyStatus](#webapplicationfirewallpolicystatus)***||
## Phase(`string` alias)

Appears on:[WebApplicationFirewallPolicyStatus](#webapplicationfirewallpolicystatus)

## WebApplicationFirewallPolicySpec

Appears on:[WebApplicationFirewallPolicy](#webapplicationfirewallpolicy), [WebApplicationFirewallPolicyStatus](#webapplicationfirewallpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `customRules` | ***[[]WebApplicationFirewallPolicySpecCustomRules](#webapplicationfirewallpolicyspeccustomrules)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `policySettings` | ***[[]WebApplicationFirewallPolicySpecPolicySettings](#webapplicationfirewallpolicyspecpolicysettings)***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## WebApplicationFirewallPolicySpecCustomRules

Appears on:[WebApplicationFirewallPolicySpec](#webapplicationfirewallpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***||
| `matchConditions` | ***[[]WebApplicationFirewallPolicySpecCustomRulesMatchConditions](#webapplicationfirewallpolicyspeccustomrulesmatchconditions)***||
| `name` | ***string***| ***(Optional)*** |
| `priority` | ***int64***||
| `ruleType` | ***string***||
## WebApplicationFirewallPolicySpecCustomRulesMatchConditions

Appears on:[WebApplicationFirewallPolicySpecCustomRules](#webapplicationfirewallpolicyspeccustomrules)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `matchValues` | ***[]string***||
| `matchVariables` | ***[[]WebApplicationFirewallPolicySpecCustomRulesMatchConditionsMatchVariables](#webapplicationfirewallpolicyspeccustomrulesmatchconditionsmatchvariables)***||
| `negationCondition` | ***bool***| ***(Optional)*** |
| `operator` | ***string***||
## WebApplicationFirewallPolicySpecCustomRulesMatchConditionsMatchVariables

Appears on:[WebApplicationFirewallPolicySpecCustomRulesMatchConditions](#webapplicationfirewallpolicyspeccustomrulesmatchconditions)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `selector` | ***string***| ***(Optional)*** |
| `variableName` | ***string***||
## WebApplicationFirewallPolicySpecPolicySettings

Appears on:[WebApplicationFirewallPolicySpec](#webapplicationfirewallpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `mode` | ***string***| ***(Optional)*** |
## WebApplicationFirewallPolicyStatus

Appears on:[WebApplicationFirewallPolicy](#webapplicationfirewallpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WebApplicationFirewallPolicySpec](#webapplicationfirewallpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
