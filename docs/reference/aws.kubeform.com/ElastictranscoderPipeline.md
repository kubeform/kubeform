## ElastictranscoderPipeline
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ElastictranscoderPipeline` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElastictranscoderPipelineSpec](#ElastictranscoderPipelineSpec)***||
| `status` | ***[ElastictranscoderPipelineStatus](#ElastictranscoderPipelineStatus)***||
## ElastictranscoderPipelineSpec
##### (Appears on:[ElastictranscoderPipeline](#ElastictranscoderPipeline), [ElastictranscoderPipelineStatus](#ElastictranscoderPipelineStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `awsKmsKeyArn` | ***string***| ***(Optional)*** |
| `contentConfig` | ***[[]ElastictranscoderPipelineSpecContentConfig](#ElastictranscoderPipelineSpecContentConfig)***| ***(Optional)*** |
| `contentConfigPermissions` | ***[[]ElastictranscoderPipelineSpecContentConfigPermissions](#ElastictranscoderPipelineSpecContentConfigPermissions)***| ***(Optional)*** |
| `inputBucket` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `notifications` | ***[[]ElastictranscoderPipelineSpecNotifications](#ElastictranscoderPipelineSpecNotifications)***| ***(Optional)*** |
| `outputBucket` | ***string***| ***(Optional)*** |
| `role` | ***string***||
| `thumbnailConfig` | ***[[]ElastictranscoderPipelineSpecThumbnailConfig](#ElastictranscoderPipelineSpecThumbnailConfig)***| ***(Optional)*** |
| `thumbnailConfigPermissions` | ***[[]ElastictranscoderPipelineSpecThumbnailConfigPermissions](#ElastictranscoderPipelineSpecThumbnailConfigPermissions)***| ***(Optional)*** |
## ElastictranscoderPipelineSpecContentConfig
##### (Appears on:[ElastictranscoderPipelineSpec](#ElastictranscoderPipelineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucket` | ***string***| ***(Optional)*** |
| `storageClass` | ***string***| ***(Optional)*** |
## ElastictranscoderPipelineSpecContentConfigPermissions
##### (Appears on:[ElastictranscoderPipelineSpec](#ElastictranscoderPipelineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `access` | ***[]string***| ***(Optional)*** |
| `grantee` | ***string***| ***(Optional)*** |
| `granteeType` | ***string***| ***(Optional)*** |
## ElastictranscoderPipelineSpecNotifications
##### (Appears on:[ElastictranscoderPipelineSpec](#ElastictranscoderPipelineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `completed` | ***string***| ***(Optional)*** |
| `error` | ***string***| ***(Optional)*** |
| `progressing` | ***string***| ***(Optional)*** |
| `warning` | ***string***| ***(Optional)*** |
## ElastictranscoderPipelineSpecThumbnailConfig
##### (Appears on:[ElastictranscoderPipelineSpec](#ElastictranscoderPipelineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucket` | ***string***| ***(Optional)*** |
| `storageClass` | ***string***| ***(Optional)*** |
## ElastictranscoderPipelineSpecThumbnailConfigPermissions
##### (Appears on:[ElastictranscoderPipelineSpec](#ElastictranscoderPipelineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `access` | ***[]string***| ***(Optional)*** |
| `grantee` | ***string***| ***(Optional)*** |
| `granteeType` | ***string***| ***(Optional)*** |
## ElastictranscoderPipelineStatus
##### (Appears on:[ElastictranscoderPipeline](#ElastictranscoderPipeline))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElastictranscoderPipelineSpec](#ElastictranscoderPipelineSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
