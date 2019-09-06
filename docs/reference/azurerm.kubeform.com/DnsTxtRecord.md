## DnsTxtRecord
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DnsTxtRecord` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DnsTxtRecordSpec](#DnsTxtRecordSpec)***||
| `status` | ***[DnsTxtRecordStatus](#DnsTxtRecordStatus)***||
## DnsTxtRecordSpec
##### (Appears on:[DnsTxtRecord](#DnsTxtRecord), [DnsTxtRecordStatus](#DnsTxtRecordStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `record` | ***[[]DnsTxtRecordSpecRecord](#DnsTxtRecordSpecRecord)***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `ttl` | ***int***||
| `zoneName` | ***string***||
## DnsTxtRecordSpecRecord
##### (Appears on:[DnsTxtRecordSpec](#DnsTxtRecordSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `value` | ***string***||
## DnsTxtRecordStatus
##### (Appears on:[DnsTxtRecord](#DnsTxtRecord))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DnsTxtRecordSpec](#DnsTxtRecordSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
