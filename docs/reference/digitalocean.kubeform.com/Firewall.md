---
title: Firewall
menu:
  docs_{{ .version }}:
    identifier: firewall-digitalocean.kubeform.com
    name: Firewall
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Firewall
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `Firewall` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FirewallSpec](#FirewallSpec)***||
| `status` | ***[FirewallStatus](#FirewallStatus)***||
## FirewallSpec

Appears on:[Firewall](#Firewall), [FirewallStatus](#FirewallStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `createdAt` | ***string***| ***(Optional)*** |
| `dropletIDS` | ***[]int64***| ***(Optional)*** |
| `inboundRule` | ***[[]FirewallSpecInboundRule](#FirewallSpecInboundRule)***| ***(Optional)*** |
| `name` | ***string***||
| `outboundRule` | ***[[]FirewallSpecOutboundRule](#FirewallSpecOutboundRule)***| ***(Optional)*** |
| `pendingChanges` | ***[[]FirewallSpecPendingChanges](#FirewallSpecPendingChanges)***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
## FirewallSpecInboundRule

Appears on:[FirewallSpec](#FirewallSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `portRange` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
| `sourceAddresses` | ***[]string***| ***(Optional)*** |
| `sourceDropletIDS` | ***[]int64***| ***(Optional)*** |
| `sourceLoadBalancerUids` | ***[]string***| ***(Optional)*** |
| `sourceTags` | ***[]string***| ***(Optional)*** |
## FirewallSpecOutboundRule

Appears on:[FirewallSpec](#FirewallSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `destinationAddresses` | ***[]string***| ***(Optional)*** |
| `destinationDropletIDS` | ***[]int64***| ***(Optional)*** |
| `destinationLoadBalancerUids` | ***[]string***| ***(Optional)*** |
| `destinationTags` | ***[]string***| ***(Optional)*** |
| `portRange` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
## FirewallSpecPendingChanges

Appears on:[FirewallSpec](#FirewallSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dropletID` | ***int***| ***(Optional)*** |
| `removing` | ***bool***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
## FirewallStatus

Appears on:[Firewall](#Firewall)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FirewallSpec](#FirewallSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
