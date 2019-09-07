---
title: MonitorLogProfile
menu:
  docs_v0.0.1:
    identifier: monitorlogprofile-azurerm.kubeform.com
    name: MonitorLogProfile
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## MonitorLogProfile
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MonitorLogProfile` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MonitorLogProfileSpec](#MonitorLogProfileSpec)***||
| `status` | ***[MonitorLogProfileStatus](#MonitorLogProfileStatus)***||
## MonitorLogProfileSpec
##### (Appears on:[MonitorLogProfile](#MonitorLogProfile), [MonitorLogProfileStatus](#MonitorLogProfileStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `categories` | ***[]string***||
| `locations` | ***[]string***||
| `name` | ***string***||
| `retentionPolicy` | ***[[]MonitorLogProfileSpecRetentionPolicy](#MonitorLogProfileSpecRetentionPolicy)***||
| `servicebusRuleID` | ***string***| ***(Optional)*** |
| `storageAccountID` | ***string***| ***(Optional)*** |
## MonitorLogProfileSpecRetentionPolicy
##### (Appears on:[MonitorLogProfileSpec](#MonitorLogProfileSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `days` | ***int***| ***(Optional)*** |
| `enabled` | ***bool***||
## MonitorLogProfileStatus
##### (Appears on:[MonitorLogProfile](#MonitorLogProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MonitorLogProfileSpec](#MonitorLogProfileSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
