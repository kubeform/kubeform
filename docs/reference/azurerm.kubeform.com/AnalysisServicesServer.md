---
title: AnalysisServicesServer
menu:
  docs_{{ .version }}:
    identifier: analysisservicesserver-azurerm.kubeform.com
    name: AnalysisServicesServer
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AnalysisServicesServer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AnalysisServicesServer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AnalysisServicesServerSpec](#AnalysisServicesServerSpec)***||
| `status` | ***[AnalysisServicesServerStatus](#AnalysisServicesServerStatus)***||
## AnalysisServicesServerSpec
##### (Appears on:[AnalysisServicesServer](#AnalysisServicesServer), [AnalysisServicesServerStatus](#AnalysisServicesServerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `adminUsers` | ***[]string***| ***(Optional)*** |
| `enablePowerBiService` | ***bool***| ***(Optional)*** |
| `ipv4FirewallRule` | ***[[]AnalysisServicesServerSpecIpv4FirewallRule](#AnalysisServicesServerSpecIpv4FirewallRule)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `querypoolConnectionMode` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `sku` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## AnalysisServicesServerSpecIpv4FirewallRule
##### (Appears on:[AnalysisServicesServerSpec](#AnalysisServicesServerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `rangeEnd` | ***string***||
| `rangeStart` | ***string***||
## AnalysisServicesServerStatus
##### (Appears on:[AnalysisServicesServer](#AnalysisServicesServer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AnalysisServicesServerSpec](#AnalysisServicesServerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
