## DnsSrvRecord
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DnsSrvRecord` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DnsSrvRecordSpec](#DnsSrvRecordSpec)***||
| `status` | ***[DnsSrvRecordStatus](#DnsSrvRecordStatus)***||
## DnsSrvRecordSpec
##### (Appears on:[DnsSrvRecord](#DnsSrvRecord), [DnsSrvRecordStatus](#DnsSrvRecordStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `record` | ***[[]DnsSrvRecordSpecRecord](#DnsSrvRecordSpecRecord)***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `ttl` | ***int***||
| `zoneName` | ***string***||
## DnsSrvRecordSpecRecord
##### (Appears on:[DnsSrvRecordSpec](#DnsSrvRecordSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `port` | ***int***||
| `priority` | ***int***||
| `target` | ***string***||
| `weight` | ***int***||
## DnsSrvRecordStatus
##### (Appears on:[DnsSrvRecord](#DnsSrvRecord))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DnsSrvRecordSpec](#DnsSrvRecordSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
