---
title: Firewall
menu:
  docs_{{ .version }}:
    identifier: firewall-azurerm.kubeform.com
    name: Firewall
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Firewall
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `Firewall` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FirewallSpec](#firewallspec)***||
| `status` | ***[FirewallStatus](#firewallstatus)***||
## FirewallSpec

Appears on:[Firewall](#firewall), [FirewallStatus](#firewallstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `ipConfiguration` | ***[[]FirewallSpecIpConfiguration](#firewallspecipconfiguration)***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## FirewallSpecIpConfiguration

Appears on:[FirewallSpec](#firewallspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `internalPublicIPAddressID` | ***string***| ***(Optional)*** Deprecated|
| `name` | ***string***||
| `privateIPAddress` | ***string***| ***(Optional)*** |
| `publicIPAddressID` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***||
## FirewallStatus

Appears on:[Firewall](#firewall)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FirewallSpec](#firewallspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
