## AutoscalingLifecycleHook
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AutoscalingLifecycleHook` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutoscalingLifecycleHookSpec](#AutoscalingLifecycleHookSpec)***||
| `status` | ***[AutoscalingLifecycleHookStatus](#AutoscalingLifecycleHookStatus)***||
## AutoscalingLifecycleHookSpec
##### (Appears on:[AutoscalingLifecycleHook](#AutoscalingLifecycleHook), [AutoscalingLifecycleHookStatus](#AutoscalingLifecycleHookStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoscalingGroupName` | ***string***||
| `defaultResult` | ***string***| ***(Optional)*** |
| `heartbeatTimeout` | ***int***| ***(Optional)*** |
| `lifecycleTransition` | ***string***||
| `name` | ***string***||
| `notificationMetadata` | ***string***| ***(Optional)*** |
| `notificationTargetArn` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***| ***(Optional)*** |
## AutoscalingLifecycleHookStatus
##### (Appears on:[AutoscalingLifecycleHook](#AutoscalingLifecycleHook))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutoscalingLifecycleHookSpec](#AutoscalingLifecycleHookSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
