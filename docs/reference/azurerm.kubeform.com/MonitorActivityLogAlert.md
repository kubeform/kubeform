---
title: MonitorActivityLogAlert
menu:
  docs_{{ .version }}:
    identifier: monitoractivitylogalert-azurerm.kubeform.com
    name: MonitorActivityLogAlert
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MonitorActivityLogAlert
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MonitorActivityLogAlert` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MonitorActivityLogAlertSpec](#MonitorActivityLogAlertSpec)***||
| `status` | ***[MonitorActivityLogAlertStatus](#MonitorActivityLogAlertStatus)***||
## MonitorActivityLogAlertSpec
##### (Appears on:[MonitorActivityLogAlert](#MonitorActivityLogAlert), [MonitorActivityLogAlertStatus](#MonitorActivityLogAlertStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `action` | ***[[]MonitorActivityLogAlertSpecAction](#MonitorActivityLogAlertSpecAction)***| ***(Optional)*** |
| `criteria` | ***[[]MonitorActivityLogAlertSpecCriteria](#MonitorActivityLogAlertSpecCriteria)***||
| `description` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `scopes` | ***[]string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## MonitorActivityLogAlertSpecAction
##### (Appears on:[MonitorActivityLogAlertSpec](#MonitorActivityLogAlertSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `actionGroupID` | ***string***||
| `webhookProperties` | ***map[string]string***| ***(Optional)*** |
## MonitorActivityLogAlertSpecCriteria
##### (Appears on:[MonitorActivityLogAlertSpec](#MonitorActivityLogAlertSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `caller` | ***string***| ***(Optional)*** |
| `category` | ***string***||
| `level` | ***string***| ***(Optional)*** |
| `operationName` | ***string***| ***(Optional)*** |
| `resourceGroup` | ***string***| ***(Optional)*** |
| `resourceID` | ***string***| ***(Optional)*** |
| `resourceProvider` | ***string***| ***(Optional)*** |
| `resourceType` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `subStatus` | ***string***| ***(Optional)*** |
## MonitorActivityLogAlertStatus
##### (Appears on:[MonitorActivityLogAlert](#MonitorActivityLogAlert))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MonitorActivityLogAlertSpec](#MonitorActivityLogAlertSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
