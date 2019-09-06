## ManagementGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ManagementGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ManagementGroupSpec](#ManagementGroupSpec)***||
| `status` | ***[ManagementGroupStatus](#ManagementGroupStatus)***||
## ManagementGroupSpec
##### (Appears on:[ManagementGroup](#ManagementGroup), [ManagementGroupStatus](#ManagementGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `displayName` | ***string***| ***(Optional)*** |
| `groupID` | ***string***| ***(Optional)*** |
| `parentManagementGroupID` | ***string***| ***(Optional)*** |
| `subscriptionIDS` | ***[]string***| ***(Optional)*** |
## ManagementGroupStatus
##### (Appears on:[ManagementGroup](#ManagementGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ManagementGroupSpec](#ManagementGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
