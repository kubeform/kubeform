## DnsAaaaRecord
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DnsAaaaRecord` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DnsAaaaRecordSpec](#DnsAaaaRecordSpec)***||
| `status` | ***[DnsAaaaRecordStatus](#DnsAaaaRecordStatus)***||
## DnsAaaaRecordSpec
##### (Appears on:[DnsAaaaRecord](#DnsAaaaRecord), [DnsAaaaRecordStatus](#DnsAaaaRecordStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `records` | ***[]string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `ttl` | ***int***||
| `zoneName` | ***string***||
## DnsAaaaRecordStatus
##### (Appears on:[DnsAaaaRecord](#DnsAaaaRecord))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DnsAaaaRecordSpec](#DnsAaaaRecordSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
