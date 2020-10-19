---
title: FrontdoorFirewallPolicy
menu:
  docs_{{ .version }}:
    identifier: frontdoorfirewallpolicy-azurerm.kubeform.com
    name: FrontdoorFirewallPolicy
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## FrontdoorFirewallPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `FrontdoorFirewallPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FrontdoorFirewallPolicySpec](#frontdoorfirewallpolicyspec)***||
| `status` | ***[FrontdoorFirewallPolicyStatus](#frontdoorfirewallpolicystatus)***||
## FrontdoorFirewallPolicySpec

Appears on:[FrontdoorFirewallPolicy](#frontdoorfirewallpolicy), [FrontdoorFirewallPolicyStatus](#frontdoorfirewallpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `customBlockResponseBody` | ***string***| ***(Optional)*** |
| `customBlockResponseStatusCode` | ***int64***| ***(Optional)*** |
| `customRule` | ***[[]FrontdoorFirewallPolicySpecCustomRule](#frontdoorfirewallpolicyspeccustomrule)***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `frontendEndpointIDS` | ***[]string***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** |
| `managedRule` | ***[[]FrontdoorFirewallPolicySpecManagedRule](#frontdoorfirewallpolicyspecmanagedrule)***| ***(Optional)*** |
| `mode` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `redirectURL` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## FrontdoorFirewallPolicySpecCustomRule

Appears on:[FrontdoorFirewallPolicySpec](#frontdoorfirewallpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `matchCondition` | ***[[]FrontdoorFirewallPolicySpecCustomRuleMatchCondition](#frontdoorfirewallpolicyspeccustomrulematchcondition)***| ***(Optional)*** |
| `name` | ***string***||
| `priority` | ***int64***| ***(Optional)*** |
| `rateLimitDurationInMinutes` | ***int64***| ***(Optional)*** |
| `rateLimitThreshold` | ***int64***| ***(Optional)*** |
| `type` | ***string***||
## FrontdoorFirewallPolicySpecCustomRuleMatchCondition

Appears on:[FrontdoorFirewallPolicySpecCustomRule](#frontdoorfirewallpolicyspeccustomrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `matchValues` | ***[]string***||
| `matchVariable` | ***string***||
| `negationCondition` | ***bool***| ***(Optional)*** |
| `operator` | ***string***||
| `selector` | ***string***| ***(Optional)*** |
| `transforms` | ***[]string***| ***(Optional)*** |
## FrontdoorFirewallPolicySpecManagedRule

Appears on:[FrontdoorFirewallPolicySpec](#frontdoorfirewallpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `exclusion` | ***[[]FrontdoorFirewallPolicySpecManagedRuleExclusion](#frontdoorfirewallpolicyspecmanagedruleexclusion)***| ***(Optional)*** |
| `override` | ***[[]FrontdoorFirewallPolicySpecManagedRuleOverride](#frontdoorfirewallpolicyspecmanagedruleoverride)***| ***(Optional)*** |
| `type` | ***string***||
| `version` | ***string***||
## FrontdoorFirewallPolicySpecManagedRuleExclusion

Appears on:[FrontdoorFirewallPolicySpecManagedRule](#frontdoorfirewallpolicyspecmanagedrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `matchVariable` | ***string***||
| `operator` | ***string***||
| `selector` | ***string***||
## FrontdoorFirewallPolicySpecManagedRuleOverride

Appears on:[FrontdoorFirewallPolicySpecManagedRule](#frontdoorfirewallpolicyspecmanagedrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `exclusion` | ***[[]FrontdoorFirewallPolicySpecManagedRuleOverrideExclusion](#frontdoorfirewallpolicyspecmanagedruleoverrideexclusion)***| ***(Optional)*** |
| `rule` | ***[[]FrontdoorFirewallPolicySpecManagedRuleOverrideRule](#frontdoorfirewallpolicyspecmanagedruleoverriderule)***| ***(Optional)*** |
| `ruleGroupName` | ***string***||
## FrontdoorFirewallPolicySpecManagedRuleOverrideExclusion

Appears on:[FrontdoorFirewallPolicySpecManagedRuleOverride](#frontdoorfirewallpolicyspecmanagedruleoverride)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `matchVariable` | ***string***||
| `operator` | ***string***||
| `selector` | ***string***||
## FrontdoorFirewallPolicySpecManagedRuleOverrideRule

Appears on:[FrontdoorFirewallPolicySpecManagedRuleOverride](#frontdoorfirewallpolicyspecmanagedruleoverride)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `exclusion` | ***[[]FrontdoorFirewallPolicySpecManagedRuleOverrideRuleExclusion](#frontdoorfirewallpolicyspecmanagedruleoverrideruleexclusion)***| ***(Optional)*** |
| `ruleID` | ***string***||
## FrontdoorFirewallPolicySpecManagedRuleOverrideRuleExclusion

Appears on:[FrontdoorFirewallPolicySpecManagedRuleOverrideRule](#frontdoorfirewallpolicyspecmanagedruleoverriderule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `matchVariable` | ***string***||
| `operator` | ***string***||
| `selector` | ***string***||
## FrontdoorFirewallPolicyStatus

Appears on:[FrontdoorFirewallPolicy](#frontdoorfirewallpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FrontdoorFirewallPolicySpec](#frontdoorfirewallpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[FrontdoorFirewallPolicyStatus](#frontdoorfirewallpolicystatus)

---
