---
title: MonitorMetricAlertrule
menu:
  docs_{{ .version }}:
    identifier: monitormetricalertrule-azurerm.kubeform.com
    name: MonitorMetricAlertrule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MonitorMetricAlertrule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MonitorMetricAlertrule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MonitorMetricAlertruleSpec](#MonitorMetricAlertruleSpec)***||
| `status` | ***[MonitorMetricAlertruleStatus](#MonitorMetricAlertruleStatus)***||
## MonitorMetricAlertruleSpec
##### (Appears on:[MonitorMetricAlertrule](#MonitorMetricAlertrule), [MonitorMetricAlertruleStatus](#MonitorMetricAlertruleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `aggregation` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `emailAction` | ***[[]MonitorMetricAlertruleSpecEmailAction](#MonitorMetricAlertruleSpecEmailAction)***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `location` | ***string***||
| `metricName` | ***string***||
| `name` | ***string***||
| `operator` | ***string***||
| `period` | ***string***||
| `resourceGroupName` | ***string***||
| `resourceID` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `threshold` | ***encoding/json.Number***||
| `webhookAction` | ***[[]MonitorMetricAlertruleSpecWebhookAction](#MonitorMetricAlertruleSpecWebhookAction)***| ***(Optional)*** |
## MonitorMetricAlertruleSpecEmailAction
##### (Appears on:[MonitorMetricAlertruleSpec](#MonitorMetricAlertruleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `customEmails` | ***[]string***| ***(Optional)*** |
| `sendToServiceOwners` | ***bool***| ***(Optional)*** |
## MonitorMetricAlertruleSpecWebhookAction
##### (Appears on:[MonitorMetricAlertruleSpec](#MonitorMetricAlertruleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `properties` | ***map[string]string***| ***(Optional)*** |
| `serviceURI` | ***string***||
## MonitorMetricAlertruleStatus
##### (Appears on:[MonitorMetricAlertrule](#MonitorMetricAlertrule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MonitorMetricAlertruleSpec](#MonitorMetricAlertruleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
