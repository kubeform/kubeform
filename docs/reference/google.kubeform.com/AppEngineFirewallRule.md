---
title: AppEngineFirewallRule
menu:
  docs_{{ .version }}:
    identifier: appenginefirewallrule-google.kubeform.com
    name: AppEngineFirewallRule
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AppEngineFirewallRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `AppEngineFirewallRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppEngineFirewallRuleSpec](#appenginefirewallrulespec)***||
| `status` | ***[AppEngineFirewallRuleStatus](#appenginefirewallrulestatus)***||
## AppEngineFirewallRuleSpec

Appears on:[AppEngineFirewallRule](#appenginefirewallrule), [AppEngineFirewallRuleStatus](#appenginefirewallrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `action` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `priority` | ***int64***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `sourceRange` | ***string***||
## AppEngineFirewallRuleStatus

Appears on:[AppEngineFirewallRule](#appenginefirewallrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppEngineFirewallRuleSpec](#appenginefirewallrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AppEngineFirewallRuleStatus](#appenginefirewallrulestatus)

---
