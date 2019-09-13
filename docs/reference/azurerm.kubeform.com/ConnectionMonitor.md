---
title: ConnectionMonitor
menu:
  docs_{{ .version }}:
    identifier: connectionmonitor-azurerm.kubeform.com
    name: ConnectionMonitor
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ConnectionMonitor
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ConnectionMonitor` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ConnectionMonitorSpec](#ConnectionMonitorSpec)***||
| `status` | ***[ConnectionMonitorStatus](#ConnectionMonitorStatus)***||
## ConnectionMonitorSpec
##### (Appears on:[ConnectionMonitor](#ConnectionMonitor), [ConnectionMonitorStatus](#ConnectionMonitorStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoStart` | ***bool***| ***(Optional)*** |
| `destination` | ***[[]ConnectionMonitorSpecDestination](#ConnectionMonitorSpecDestination)***||
| `intervalInSeconds` | ***int***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `networkWatcherName` | ***string***||
| `resourceGroupName` | ***string***||
| `source` | ***[[]ConnectionMonitorSpecSource](#ConnectionMonitorSpecSource)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## ConnectionMonitorSpecDestination
##### (Appears on:[ConnectionMonitorSpec](#ConnectionMonitorSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `address` | ***string***| ***(Optional)*** |
| `port` | ***int***||
| `virtualMachineID` | ***string***| ***(Optional)*** |
## ConnectionMonitorSpecSource
##### (Appears on:[ConnectionMonitorSpec](#ConnectionMonitorSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `port` | ***int***| ***(Optional)*** |
| `virtualMachineID` | ***string***||
## ConnectionMonitorStatus
##### (Appears on:[ConnectionMonitor](#ConnectionMonitor))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ConnectionMonitorSpec](#ConnectionMonitorSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
