## AppsyncGraphqlAPI
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AppsyncGraphqlAPI` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppsyncGraphqlAPISpec](#AppsyncGraphqlAPISpec)***||
| `status` | ***[AppsyncGraphqlAPIStatus](#AppsyncGraphqlAPIStatus)***||
## AppsyncGraphqlAPISpec
##### (Appears on:[AppsyncGraphqlAPI](#AppsyncGraphqlAPI), [AppsyncGraphqlAPIStatus](#AppsyncGraphqlAPIStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `authenticationType` | ***string***||
| `logConfig` | ***[[]AppsyncGraphqlAPISpecLogConfig](#AppsyncGraphqlAPISpecLogConfig)***| ***(Optional)*** |
| `name` | ***string***||
| `openidConnectConfig` | ***[[]AppsyncGraphqlAPISpecOpenidConnectConfig](#AppsyncGraphqlAPISpecOpenidConnectConfig)***| ***(Optional)*** |
| `schema` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `uris` | ***map[string]string***| ***(Optional)*** |
| `userPoolConfig` | ***[[]AppsyncGraphqlAPISpecUserPoolConfig](#AppsyncGraphqlAPISpecUserPoolConfig)***| ***(Optional)*** |
## AppsyncGraphqlAPISpecLogConfig
##### (Appears on:[AppsyncGraphqlAPISpec](#AppsyncGraphqlAPISpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cloudwatchLogsRoleArn` | ***string***||
| `fieldLogLevel` | ***string***||
## AppsyncGraphqlAPISpecOpenidConnectConfig
##### (Appears on:[AppsyncGraphqlAPISpec](#AppsyncGraphqlAPISpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `authTtl` | ***int***| ***(Optional)*** |
| `clientID` | ***string***| ***(Optional)*** |
| `iatTtl` | ***int***| ***(Optional)*** |
| `issuer` | ***string***||
## AppsyncGraphqlAPISpecUserPoolConfig
##### (Appears on:[AppsyncGraphqlAPISpec](#AppsyncGraphqlAPISpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `appIDClientRegex` | ***string***| ***(Optional)*** |
| `awsRegion` | ***string***| ***(Optional)*** |
| `defaultAction` | ***string***||
| `userPoolID` | ***string***||
## AppsyncGraphqlAPIStatus
##### (Appears on:[AppsyncGraphqlAPI](#AppsyncGraphqlAPI))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppsyncGraphqlAPISpec](#AppsyncGraphqlAPISpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
