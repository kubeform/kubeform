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
| `spec` | ***[ComputeFirewallSpec](#ComputeFirewallSpec)***||
| `status` | ***[ComputeFirewallStatus](#ComputeFirewallStatus)***||
## ComputeFirewallSpec

Appears on:[ComputeFirewall](#ComputeFirewall), [ComputeFirewallStatus](#ComputeFirewallStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allow` | ***[[]ComputeFirewallSpecAllow](#ComputeFirewallSpecAllow)***| ***(Optional)*** |
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `deny` | ***[[]ComputeFirewallSpecDeny](#ComputeFirewallSpecDeny)***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `destinationRanges` | ***[]string***| ***(Optional)*** |
| `direction` | ***string***| ***(Optional)*** |
| `disabled` | ***bool***| ***(Optional)*** |
| `enableLogging` | ***bool***| ***(Optional)*** Deprecated|
| `name` | ***string***||
| `network` | ***string***||
| `priority` | ***int***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `sourceRanges` | ***[]string***| ***(Optional)*** |
| `sourceServiceAccounts` | ***[]string***| ***(Optional)*** |
| `sourceTags` | ***[]string***| ***(Optional)*** |
| `targetServiceAccounts` | ***[]string***| ***(Optional)*** |
| `targetTags` | ***[]string***| ***(Optional)*** |
## ComputeFirewallSpecAllow

Appears on:[ComputeFirewallSpec](#ComputeFirewallSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ports` | ***[]string***| ***(Optional)*** |
| `protocol` | ***string***||
## ComputeFirewallSpecDeny

Appears on:[ComputeFirewallSpec](#ComputeFirewallSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ports` | ***[]string***| ***(Optional)*** |
| `protocol` | ***string***||
## ComputeFirewallStatus

Appears on:[ComputeFirewall](#ComputeFirewall)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeFirewallSpec](#ComputeFirewallSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
