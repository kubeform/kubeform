## MetricAlertrule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MetricAlertrule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MetricAlertruleSpec](#MetricAlertruleSpec)***||
| `status` | ***[MetricAlertruleStatus](#MetricAlertruleStatus)***||
## MetricAlertruleSpec
##### (Appears on:[MetricAlertrule](#MetricAlertrule), [MetricAlertruleStatus](#MetricAlertruleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `aggregation` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `emailAction` | ***[[]MetricAlertruleSpecEmailAction](#MetricAlertruleSpecEmailAction)***| ***(Optional)*** |
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
| `webhookAction` | ***[[]MetricAlertruleSpecWebhookAction](#MetricAlertruleSpecWebhookAction)***| ***(Optional)*** |
## MetricAlertruleSpecEmailAction
##### (Appears on:[MetricAlertruleSpec](#MetricAlertruleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `customEmails` | ***[]string***| ***(Optional)*** |
| `sendToServiceOwners` | ***bool***| ***(Optional)*** |
## MetricAlertruleSpecWebhookAction
##### (Appears on:[MetricAlertruleSpec](#MetricAlertruleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `properties` | ***map[string]string***| ***(Optional)*** |
| `serviceURI` | ***string***||
## MetricAlertruleStatus
##### (Appears on:[MetricAlertrule](#MetricAlertrule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MetricAlertruleSpec](#MetricAlertruleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
