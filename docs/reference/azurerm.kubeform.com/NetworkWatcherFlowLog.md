---
title: NetworkWatcherFlowLog
menu:
  docs_{{ .version }}:
    identifier: networkwatcherflowlog-azurerm.kubeform.com
    name: NetworkWatcherFlowLog
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NetworkWatcherFlowLog
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkWatcherFlowLog` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkWatcherFlowLogSpec](#networkwatcherflowlogspec)***||
| `status` | ***[NetworkWatcherFlowLogStatus](#networkwatcherflowlogstatus)***||
## NetworkWatcherFlowLogSpec

Appears on:[NetworkWatcherFlowLog](#networkwatcherflowlog), [NetworkWatcherFlowLogStatus](#networkwatcherflowlogstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `enabled` | ***bool***||
| `networkSecurityGroupID` | ***string***||
| `networkWatcherName` | ***string***||
| `resourceGroupName` | ***string***||
| `retentionPolicy` | ***[[]NetworkWatcherFlowLogSpecRetentionPolicy](#networkwatcherflowlogspecretentionpolicy)***||
| `storageAccountID` | ***string***||
| `trafficAnalytics` | ***[[]NetworkWatcherFlowLogSpecTrafficAnalytics](#networkwatcherflowlogspectrafficanalytics)***| ***(Optional)*** |
| `version` | ***int64***| ***(Optional)*** |
## NetworkWatcherFlowLogSpecRetentionPolicy

Appears on:[NetworkWatcherFlowLogSpec](#networkwatcherflowlogspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `days` | ***int64***||
| `enabled` | ***bool***||
## NetworkWatcherFlowLogSpecTrafficAnalytics

Appears on:[NetworkWatcherFlowLogSpec](#networkwatcherflowlogspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `workspaceID` | ***string***||
| `workspaceRegion` | ***string***||
| `workspaceResourceID` | ***string***||
## NetworkWatcherFlowLogStatus

Appears on:[NetworkWatcherFlowLog](#networkwatcherflowlog)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkWatcherFlowLogSpec](#networkwatcherflowlogspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[NetworkWatcherFlowLogStatus](#networkwatcherflowlogstatus)

---
