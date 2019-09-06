## CognitoResourceServer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CognitoResourceServer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CognitoResourceServerSpec](#CognitoResourceServerSpec)***||
| `status` | ***[CognitoResourceServerStatus](#CognitoResourceServerStatus)***||
## CognitoResourceServerSpec
##### (Appears on:[CognitoResourceServer](#CognitoResourceServer), [CognitoResourceServerStatus](#CognitoResourceServerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `identifier` | ***string***||
| `name` | ***string***||
| `scope` | ***[[]CognitoResourceServerSpecScope](#CognitoResourceServerSpecScope)***| ***(Optional)*** |
| `scopeIdentifiers` | ***[]string***| ***(Optional)*** |
| `userPoolID` | ***string***||
## CognitoResourceServerSpecScope
##### (Appears on:[CognitoResourceServerSpec](#CognitoResourceServerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `scopeDescription` | ***string***||
| `scopeName` | ***string***||
## CognitoResourceServerStatus
##### (Appears on:[CognitoResourceServer](#CognitoResourceServer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CognitoResourceServerSpec](#CognitoResourceServerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
