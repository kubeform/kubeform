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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MonitorActivityLogAlertSpec](#monitoractivitylogalertspec)***||
| `status` | ***[MonitorActivityLogAlertStatus](#monitoractivitylogalertstatus)***||
## MonitorActivityLogAlertSpec

Appears on:[MonitorActivityLogAlert](#monitoractivitylogalert), [MonitorActivityLogAlertStatus](#monitoractivitylogalertstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `action` | ***[[]MonitorActivityLogAlertSpecAction](#monitoractivitylogalertspecaction)***| ***(Optional)*** |
| `criteria` | ***[[]MonitorActivityLogAlertSpecCriteria](#monitoractivitylogalertspeccriteria)***||
| `description` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `scopes` | ***[]string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## MonitorActivityLogAlertSpecAction

Appears on:[MonitorActivityLogAlertSpec](#monitoractivitylogalertspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `actionGroupID` | ***string***||
| `webhookProperties` | ***map[string]string***| ***(Optional)*** |
## MonitorActivityLogAlertSpecCriteria

Appears on:[MonitorActivityLogAlertSpec](#monitoractivitylogalertspec)

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

Appears on:[MonitorActivityLogAlert](#monitoractivitylogalert)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MonitorActivityLogAlertSpec](#monitoractivitylogalertspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[MonitorActivityLogAlertStatus](#monitoractivitylogalertstatus)

---
