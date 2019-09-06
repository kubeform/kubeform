## SsmMaintenanceWindowTask
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SsmMaintenanceWindowTask` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SsmMaintenanceWindowTaskSpec](#SsmMaintenanceWindowTaskSpec)***||
| `status` | ***[SsmMaintenanceWindowTaskStatus](#SsmMaintenanceWindowTaskStatus)***||
## SsmMaintenanceWindowTaskSpec
##### (Appears on:[SsmMaintenanceWindowTask](#SsmMaintenanceWindowTask), [SsmMaintenanceWindowTaskStatus](#SsmMaintenanceWindowTaskStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `loggingInfo` | ***[[]SsmMaintenanceWindowTaskSpecLoggingInfo](#SsmMaintenanceWindowTaskSpecLoggingInfo)***| ***(Optional)*** |
| `maxConcurrency` | ***string***||
| `maxErrors` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `priority` | ***int***| ***(Optional)*** |
| `serviceRoleArn` | ***string***||
| `targets` | ***[[]SsmMaintenanceWindowTaskSpecTargets](#SsmMaintenanceWindowTaskSpecTargets)***||
| `taskArn` | ***string***||
| `taskParameters` | ***[[]SsmMaintenanceWindowTaskSpecTaskParameters](#SsmMaintenanceWindowTaskSpecTaskParameters)***| ***(Optional)*** |
| `taskType` | ***string***||
| `windowID` | ***string***||
## SsmMaintenanceWindowTaskSpecLoggingInfo
##### (Appears on:[SsmMaintenanceWindowTaskSpec](#SsmMaintenanceWindowTaskSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `s3BucketName` | ***string***||
| `s3BucketPrefix` | ***string***| ***(Optional)*** |
| `s3Region` | ***string***||
## SsmMaintenanceWindowTaskSpecTargets
##### (Appears on:[SsmMaintenanceWindowTaskSpec](#SsmMaintenanceWindowTaskSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***||
| `values` | ***[]string***||
## SsmMaintenanceWindowTaskSpecTaskParameters
##### (Appears on:[SsmMaintenanceWindowTaskSpec](#SsmMaintenanceWindowTaskSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `values` | ***[]string***||
## SsmMaintenanceWindowTaskStatus
##### (Appears on:[SsmMaintenanceWindowTask](#SsmMaintenanceWindowTask))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SsmMaintenanceWindowTaskSpec](#SsmMaintenanceWindowTaskSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
