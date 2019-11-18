---
title: NetworkConnectionMonitor
menu:
  docs_{{ .version }}:
    identifier: networkconnectionmonitor-azurerm.kubeform.com
    name: NetworkConnectionMonitor
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NetworkConnectionMonitor
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkConnectionMonitor` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkConnectionMonitorSpec](#networkconnectionmonitorspec)***||
| `status` | ***[NetworkConnectionMonitorStatus](#networkconnectionmonitorstatus)***||
## NetworkConnectionMonitorSpec

Appears on:[NetworkConnectionMonitor](#networkconnectionmonitor), [NetworkConnectionMonitorStatus](#networkconnectionmonitorstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoStart` | ***bool***| ***(Optional)*** |
| `destination` | ***[[]NetworkConnectionMonitorSpecDestination](#networkconnectionmonitorspecdestination)***||
| `intervalInSeconds` | ***int64***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `networkWatcherName` | ***string***||
| `resourceGroupName` | ***string***||
| `source` | ***[[]NetworkConnectionMonitorSpecSource](#networkconnectionmonitorspecsource)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## NetworkConnectionMonitorSpecDestination

Appears on:[NetworkConnectionMonitorSpec](#networkconnectionmonitorspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `address` | ***string***| ***(Optional)*** |
| `port` | ***int64***||
| `virtualMachineID` | ***string***| ***(Optional)*** |
## NetworkConnectionMonitorSpecSource

Appears on:[NetworkConnectionMonitorSpec](#networkconnectionmonitorspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `port` | ***int64***| ***(Optional)*** |
| `virtualMachineID` | ***string***||
## NetworkConnectionMonitorStatus

Appears on:[NetworkConnectionMonitor](#networkconnectionmonitor)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkConnectionMonitorSpec](#networkconnectionmonitorspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[NetworkConnectionMonitorStatus](#networkconnectionmonitorstatus)

---
