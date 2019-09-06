## AutomationVariableString
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AutomationVariableString` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutomationVariableStringSpec](#AutomationVariableStringSpec)***||
| `status` | ***[AutomationVariableStringStatus](#AutomationVariableStringStatus)***||
## AutomationVariableStringSpec
##### (Appears on:[AutomationVariableString](#AutomationVariableString), [AutomationVariableStringStatus](#AutomationVariableStringStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `automationAccountName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `encrypted` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `value` | ***string***| ***(Optional)*** |
## AutomationVariableStringStatus
##### (Appears on:[AutomationVariableString](#AutomationVariableString))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutomationVariableStringSpec](#AutomationVariableStringSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
