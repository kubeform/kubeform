## DnsNsRecord
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DnsNsRecord` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DnsNsRecordSpec](#DnsNsRecordSpec)***||
| `status` | ***[DnsNsRecordStatus](#DnsNsRecordStatus)***||
## DnsNsRecordSpec
##### (Appears on:[DnsNsRecord](#DnsNsRecord), [DnsNsRecordStatus](#DnsNsRecordStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `record` | ***[[]DnsNsRecordSpecRecord](#DnsNsRecordSpecRecord)***| ***(Optional)*** Deprecated|
| `records` | ***[]string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `ttl` | ***int***||
| `zoneName` | ***string***||
## DnsNsRecordSpecRecord
##### (Appears on:[DnsNsRecordSpec](#DnsNsRecordSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `nsdname` | ***string***||
## DnsNsRecordStatus
##### (Appears on:[DnsNsRecord](#DnsNsRecord))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DnsNsRecordSpec](#DnsNsRecordSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
