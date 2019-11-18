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
| `spec` | ***[ConnectionMonitorSpec](#connectionmonitorspec)***||
| `status` | ***[ConnectionMonitorStatus](#connectionmonitorstatus)***||
## ConnectionMonitorSpec

Appears on:[ConnectionMonitor](#connectionmonitor), [ConnectionMonitorStatus](#connectionmonitorstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoStart` | ***bool***| ***(Optional)*** |
| `destination` | ***[[]ConnectionMonitorSpecDestination](#connectionmonitorspecdestination)***||
| `intervalInSeconds` | ***int64***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `networkWatcherName` | ***string***||
| `resourceGroupName` | ***string***||
| `source` | ***[[]ConnectionMonitorSpecSource](#connectionmonitorspecsource)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## ConnectionMonitorSpecDestination

Appears on:[ConnectionMonitorSpec](#connectionmonitorspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `address` | ***string***| ***(Optional)*** |
| `port` | ***int64***||
| `virtualMachineID` | ***string***| ***(Optional)*** |
## ConnectionMonitorSpecSource

Appears on:[ConnectionMonitorSpec](#connectionmonitorspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `port` | ***int64***| ***(Optional)*** |
| `virtualMachineID` | ***string***||
## ConnectionMonitorStatus

Appears on:[ConnectionMonitor](#connectionmonitor)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ConnectionMonitorSpec](#connectionmonitorspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ConnectionMonitorStatus](#connectionmonitorstatus)

---
