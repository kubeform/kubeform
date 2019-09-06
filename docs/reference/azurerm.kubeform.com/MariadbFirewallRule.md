## MariadbFirewallRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MariadbFirewallRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MariadbFirewallRuleSpec](#MariadbFirewallRuleSpec)***||
| `status` | ***[MariadbFirewallRuleStatus](#MariadbFirewallRuleStatus)***||
## MariadbFirewallRuleSpec
##### (Appears on:[MariadbFirewallRule](#MariadbFirewallRule), [MariadbFirewallRuleStatus](#MariadbFirewallRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `endIPAddress` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `serverName` | ***string***||
| `startIPAddress` | ***string***||
## MariadbFirewallRuleStatus
##### (Appears on:[MariadbFirewallRule](#MariadbFirewallRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MariadbFirewallRuleSpec](#MariadbFirewallRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
