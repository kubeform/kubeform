## DataLakeAnalyticsFirewallRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DataLakeAnalyticsFirewallRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataLakeAnalyticsFirewallRuleSpec](#DataLakeAnalyticsFirewallRuleSpec)***||
| `status` | ***[DataLakeAnalyticsFirewallRuleStatus](#DataLakeAnalyticsFirewallRuleStatus)***||
## DataLakeAnalyticsFirewallRuleSpec
##### (Appears on:[DataLakeAnalyticsFirewallRule](#DataLakeAnalyticsFirewallRule), [DataLakeAnalyticsFirewallRuleStatus](#DataLakeAnalyticsFirewallRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***||
| `endIPAddress` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `startIPAddress` | ***string***||
## DataLakeAnalyticsFirewallRuleStatus
##### (Appears on:[DataLakeAnalyticsFirewallRule](#DataLakeAnalyticsFirewallRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataLakeAnalyticsFirewallRuleSpec](#DataLakeAnalyticsFirewallRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
