## EventhubAuthorizationRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `EventhubAuthorizationRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EventhubAuthorizationRuleSpec](#EventhubAuthorizationRuleSpec)***||
| `status` | ***[EventhubAuthorizationRuleStatus](#EventhubAuthorizationRuleStatus)***||
## EventhubAuthorizationRuleSpec
##### (Appears on:[EventhubAuthorizationRule](#EventhubAuthorizationRule), [EventhubAuthorizationRuleStatus](#EventhubAuthorizationRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `eventhubName` | ***string***||
| `listen` | ***bool***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** Deprecated|
| `manage` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `namespaceName` | ***string***||
| `resourceGroupName` | ***string***||
| `send` | ***bool***| ***(Optional)*** |
## EventhubAuthorizationRuleStatus
##### (Appears on:[EventhubAuthorizationRule](#EventhubAuthorizationRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EventhubAuthorizationRuleSpec](#EventhubAuthorizationRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_connection_string` | ***string*** ||
| `primary_key` | ***string*** ||
| `secondary_connection_string` | ***string*** ||
| `secondary_key` | ***string*** ||
