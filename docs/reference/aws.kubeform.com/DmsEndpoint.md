## DmsEndpoint
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DmsEndpoint` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DmsEndpointSpec](#DmsEndpointSpec)***||
| `status` | ***[DmsEndpointStatus](#DmsEndpointStatus)***||
## DmsEndpointSpec
##### (Appears on:[DmsEndpoint](#DmsEndpoint), [DmsEndpointStatus](#DmsEndpointStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `certificateArn` | ***string***| ***(Optional)*** |
| `databaseName` | ***string***| ***(Optional)*** |
| `endpointArn` | ***string***| ***(Optional)*** |
| `endpointID` | ***string***||
| `endpointType` | ***string***||
| `engineName` | ***string***||
| `extraConnectionAttributes` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `mongodbSettings` | ***[[]DmsEndpointSpecMongodbSettings](#DmsEndpointSpecMongodbSettings)***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `s3Settings` | ***[[]DmsEndpointSpecS3Settings](#DmsEndpointSpecS3Settings)***| ***(Optional)*** |
| `serverName` | ***string***| ***(Optional)*** |
| `serviceAccessRole` | ***string***| ***(Optional)*** |
| `sslMode` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `username` | ***string***| ***(Optional)*** |
## DmsEndpointSpecMongodbSettings
##### (Appears on:[DmsEndpointSpec](#DmsEndpointSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `authMechanism` | ***string***| ***(Optional)*** |
| `authSource` | ***string***| ***(Optional)*** |
| `authType` | ***string***| ***(Optional)*** |
| `docsToInvestigate` | ***string***| ***(Optional)*** |
| `extractDocID` | ***string***| ***(Optional)*** |
| `nestingLevel` | ***string***| ***(Optional)*** |
## DmsEndpointSpecS3Settings
##### (Appears on:[DmsEndpointSpec](#DmsEndpointSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketFolder` | ***string***| ***(Optional)*** |
| `bucketName` | ***string***| ***(Optional)*** |
| `compressionType` | ***string***| ***(Optional)*** |
| `csvDelimiter` | ***string***| ***(Optional)*** |
| `csvRowDelimiter` | ***string***| ***(Optional)*** |
| `externalTableDefinition` | ***string***| ***(Optional)*** |
| `serviceAccessRoleArn` | ***string***| ***(Optional)*** |
## DmsEndpointStatus
##### (Appears on:[DmsEndpoint](#DmsEndpoint))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DmsEndpointSpec](#DmsEndpointSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `password` | ***string*** ||
