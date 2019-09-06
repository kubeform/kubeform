## NotificationHubAuthorizationRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NotificationHubAuthorizationRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NotificationHubAuthorizationRuleSpec](#NotificationHubAuthorizationRuleSpec)***||
| `status` | ***[NotificationHubAuthorizationRuleStatus](#NotificationHubAuthorizationRuleStatus)***||
## NotificationHubAuthorizationRuleSpec
##### (Appears on:[NotificationHubAuthorizationRule](#NotificationHubAuthorizationRule), [NotificationHubAuthorizationRuleStatus](#NotificationHubAuthorizationRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `listen` | ***bool***| ***(Optional)*** |
| `manage` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `namespaceName` | ***string***||
| `notificationHubName` | ***string***||
| `primaryAccessKey` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `secondaryAccessKey` | ***string***| ***(Optional)*** |
| `send` | ***bool***| ***(Optional)*** |
## NotificationHubAuthorizationRuleStatus
##### (Appears on:[NotificationHubAuthorizationRule](#NotificationHubAuthorizationRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NotificationHubAuthorizationRuleSpec](#NotificationHubAuthorizationRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
