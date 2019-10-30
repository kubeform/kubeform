---
title: MetricAlertrule
menu:
  docs_{{ .version }}:
    identifier: metricalertrule-azurerm.kubeform.com
    name: MetricAlertrule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MetricAlertrule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MetricAlertrule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MetricAlertruleSpec](#metricalertrulespec)***||
| `status` | ***[MetricAlertruleStatus](#metricalertrulestatus)***||
## MetricAlertruleSpec

Appears on:[MetricAlertrule](#metricalertrule), [MetricAlertruleStatus](#metricalertrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `aggregation` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `emailAction` | ***[[]MetricAlertruleSpecEmailAction](#metricalertrulespecemailaction)***| ***(Optional)*** |
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
| `webhookAction` | ***[[]MetricAlertruleSpecWebhookAction](#metricalertrulespecwebhookaction)***| ***(Optional)*** |
## MetricAlertruleSpecEmailAction

Appears on:[MetricAlertruleSpec](#metricalertrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `customEmails` | ***[]string***| ***(Optional)*** |
| `sendToServiceOwners` | ***bool***| ***(Optional)*** |
## MetricAlertruleSpecWebhookAction

Appears on:[MetricAlertruleSpec](#metricalertrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `properties` | ***map[string]string***| ***(Optional)*** |
| `serviceURI` | ***string***||
## MetricAlertruleStatus

Appears on:[MetricAlertrule](#metricalertrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MetricAlertruleSpec](#metricalertrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[MetricAlertruleStatus](#metricalertrulestatus)

---
