---
title: ComputeFirewall
menu:
  docs_{{ .version }}:
    identifier: computefirewall-google.kubeform.com
    name: ComputeFirewall
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeFirewall
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeFirewall` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeFirewallSpec](#computefirewallspec)***||
| `status` | ***[ComputeFirewallStatus](#computefirewallstatus)***||
## ComputeFirewallSpec

Appears on:[ComputeFirewall](#computefirewall), [ComputeFirewallStatus](#computefirewallstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allow` | ***[[]ComputeFirewallSpecAllow](#computefirewallspecallow)***| ***(Optional)*** |
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `deny` | ***[[]ComputeFirewallSpecDeny](#computefirewallspecdeny)***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `destinationRanges` | ***[]string***| ***(Optional)*** |
| `direction` | ***string***| ***(Optional)*** |
| `disabled` | ***bool***| ***(Optional)*** |
| `enableLogging` | ***bool***| ***(Optional)*** Deprecated|
| `name` | ***string***||
| `network` | ***string***||
| `priority` | ***int64***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `sourceRanges` | ***[]string***| ***(Optional)*** |
| `sourceServiceAccounts` | ***[]string***| ***(Optional)*** |
| `sourceTags` | ***[]string***| ***(Optional)*** |
| `targetServiceAccounts` | ***[]string***| ***(Optional)*** |
| `targetTags` | ***[]string***| ***(Optional)*** |
## ComputeFirewallSpecAllow

Appears on:[ComputeFirewallSpec](#computefirewallspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ports` | ***[]string***| ***(Optional)*** |
| `protocol` | ***string***||
## ComputeFirewallSpecDeny

Appears on:[ComputeFirewallSpec](#computefirewallspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ports` | ***[]string***| ***(Optional)*** |
| `protocol` | ***string***||
## ComputeFirewallStatus

Appears on:[ComputeFirewall](#computefirewall)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeFirewallSpec](#computefirewallspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeFirewallStatus](#computefirewallstatus)

---
