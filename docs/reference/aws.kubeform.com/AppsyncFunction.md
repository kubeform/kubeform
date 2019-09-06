## AppsyncFunction
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AppsyncFunction` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppsyncFunctionSpec](#AppsyncFunctionSpec)***||
| `status` | ***[AppsyncFunctionStatus](#AppsyncFunctionStatus)***||
## AppsyncFunctionSpec
##### (Appears on:[AppsyncFunction](#AppsyncFunction), [AppsyncFunctionStatus](#AppsyncFunctionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiID` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `dataSource` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `functionID` | ***string***| ***(Optional)*** |
| `functionVersion` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `requestMappingTemplate` | ***string***||
| `responseMappingTemplate` | ***string***||
## AppsyncFunctionStatus
##### (Appears on:[AppsyncFunction](#AppsyncFunction))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppsyncFunctionSpec](#AppsyncFunctionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
