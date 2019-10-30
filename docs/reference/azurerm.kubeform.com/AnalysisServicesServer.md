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
| `spec` | ***[AnalysisServicesServerSpec](#analysisservicesserverspec)***||
| `status` | ***[AnalysisServicesServerStatus](#analysisservicesserverstatus)***||
## AnalysisServicesServerSpec

Appears on:[AnalysisServicesServer](#analysisservicesserver), [AnalysisServicesServerStatus](#analysisservicesserverstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `adminUsers` | ***[]string***| ***(Optional)*** |
| `enablePowerBiService` | ***bool***| ***(Optional)*** |
| `ipv4FirewallRule` | ***[[]AnalysisServicesServerSpecIpv4FirewallRule](#analysisservicesserverspecipv4firewallrule)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `querypoolConnectionMode` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `sku` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## AnalysisServicesServerSpecIpv4FirewallRule

Appears on:[AnalysisServicesServerSpec](#analysisservicesserverspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `rangeEnd` | ***string***||
| `rangeStart` | ***string***||
## AnalysisServicesServerStatus

Appears on:[AnalysisServicesServer](#analysisservicesserver)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AnalysisServicesServerSpec](#analysisservicesserverspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AnalysisServicesServerStatus](#analysisservicesserverstatus)

---
