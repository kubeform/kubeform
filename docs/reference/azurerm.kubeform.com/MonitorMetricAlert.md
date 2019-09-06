## MonitorMetricAlert
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MonitorMetricAlert` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MonitorMetricAlertSpec](#MonitorMetricAlertSpec)***||
| `status` | ***[MonitorMetricAlertStatus](#MonitorMetricAlertStatus)***||
## MonitorMetricAlertSpec
##### (Appears on:[MonitorMetricAlert](#MonitorMetricAlert), [MonitorMetricAlertStatus](#MonitorMetricAlertStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `action` | ***[[]MonitorMetricAlertSpecAction](#MonitorMetricAlertSpecAction)***| ***(Optional)*** |
| `autoMitigate` | ***bool***| ***(Optional)*** |
| `criteria` | ***[[]MonitorMetricAlertSpecCriteria](#MonitorMetricAlertSpecCriteria)***||
| `description` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `frequency` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `scopes` | ***[]string***||
| `severity` | ***int***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `windowSize` | ***string***| ***(Optional)*** |
## MonitorMetricAlertSpecAction
##### (Appears on:[MonitorMetricAlertSpec](#MonitorMetricAlertSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `actionGroupID` | ***string***||
| `webhookProperties` | ***map[string]string***| ***(Optional)*** |
## MonitorMetricAlertSpecCriteria
##### (Appears on:[MonitorMetricAlertSpec](#MonitorMetricAlertSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `aggregation` | ***string***||
| `dimension` | ***[[]MonitorMetricAlertSpecCriteriaDimension](#MonitorMetricAlertSpecCriteriaDimension)***| ***(Optional)*** |
| `metricName` | ***string***||
| `metricNamespace` | ***string***||
| `operator` | ***string***||
| `threshold` | ***encoding/json.Number***||
## MonitorMetricAlertSpecCriteriaDimension
##### (Appears on:[MonitorMetricAlertSpecCriteria](#MonitorMetricAlertSpecCriteria))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `operator` | ***string***||
| `values` | ***[]string***||
## MonitorMetricAlertStatus
##### (Appears on:[MonitorMetricAlert](#MonitorMetricAlert))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MonitorMetricAlertSpec](#MonitorMetricAlertSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
