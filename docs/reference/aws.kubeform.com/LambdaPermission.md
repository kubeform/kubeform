## LambdaPermission
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LambdaPermission` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LambdaPermissionSpec](#LambdaPermissionSpec)***||
| `status` | ***[LambdaPermissionStatus](#LambdaPermissionStatus)***||
## LambdaPermissionSpec
##### (Appears on:[LambdaPermission](#LambdaPermission), [LambdaPermissionStatus](#LambdaPermissionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `action` | ***string***||
| `eventSourceToken` | ***string***| ***(Optional)*** |
| `functionName` | ***string***||
| `principal` | ***string***||
| `qualifier` | ***string***| ***(Optional)*** |
| `sourceAccount` | ***string***| ***(Optional)*** |
| `sourceArn` | ***string***| ***(Optional)*** |
| `statementID` | ***string***| ***(Optional)*** |
| `statementIDPrefix` | ***string***| ***(Optional)*** |
## LambdaPermissionStatus
##### (Appears on:[LambdaPermission](#LambdaPermission))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LambdaPermissionSpec](#LambdaPermissionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
