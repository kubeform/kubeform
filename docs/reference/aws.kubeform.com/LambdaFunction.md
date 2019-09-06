## LambdaFunction
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LambdaFunction` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LambdaFunctionSpec](#LambdaFunctionSpec)***||
| `status` | ***[LambdaFunctionStatus](#LambdaFunctionStatus)***||
## LambdaFunctionSpec
##### (Appears on:[LambdaFunction](#LambdaFunction), [LambdaFunctionStatus](#LambdaFunctionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `deadLetterConfig` | ***[[]LambdaFunctionSpecDeadLetterConfig](#LambdaFunctionSpecDeadLetterConfig)***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `environment` | ***[[]LambdaFunctionSpecEnvironment](#LambdaFunctionSpecEnvironment)***| ***(Optional)*** |
| `filename` | ***string***| ***(Optional)*** |
| `functionName` | ***string***||
| `handler` | ***string***||
| `invokeArn` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `lastModified` | ***string***| ***(Optional)*** |
| `layers` | ***[]string***| ***(Optional)*** |
| `memorySize` | ***int***| ***(Optional)*** |
| `publish` | ***bool***| ***(Optional)*** |
| `qualifiedArn` | ***string***| ***(Optional)*** |
| `reservedConcurrentExecutions` | ***int***| ***(Optional)*** |
| `role` | ***string***||
| `runtime` | ***string***||
| `s3Bucket` | ***string***| ***(Optional)*** |
| `s3Key` | ***string***| ***(Optional)*** |
| `s3ObjectVersion` | ***string***| ***(Optional)*** |
| `sourceCodeHash` | ***string***| ***(Optional)*** |
| `sourceCodeSize` | ***int***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `timeout` | ***int***| ***(Optional)*** |
| `tracingConfig` | ***[[]LambdaFunctionSpecTracingConfig](#LambdaFunctionSpecTracingConfig)***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
| `vpcConfig` | ***[[]LambdaFunctionSpecVpcConfig](#LambdaFunctionSpecVpcConfig)***| ***(Optional)*** |
## LambdaFunctionSpecDeadLetterConfig
##### (Appears on:[LambdaFunctionSpec](#LambdaFunctionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `targetArn` | ***string***||
## LambdaFunctionSpecEnvironment
##### (Appears on:[LambdaFunctionSpec](#LambdaFunctionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `variables` | ***map[string]string***| ***(Optional)*** |
## LambdaFunctionSpecTracingConfig
##### (Appears on:[LambdaFunctionSpec](#LambdaFunctionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `mode` | ***string***||
## LambdaFunctionSpecVpcConfig
##### (Appears on:[LambdaFunctionSpec](#LambdaFunctionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `securityGroupIDS` | ***[]string***||
| `subnetIDS` | ***[]string***||
| `vpcID` | ***string***| ***(Optional)*** |
## LambdaFunctionStatus
##### (Appears on:[LambdaFunction](#LambdaFunction))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LambdaFunctionSpec](#LambdaFunctionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
