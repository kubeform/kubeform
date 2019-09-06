## DataLakeStore
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DataLakeStore` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataLakeStoreSpec](#DataLakeStoreSpec)***||
| `status` | ***[DataLakeStoreStatus](#DataLakeStoreStatus)***||
## DataLakeStoreSpec
##### (Appears on:[DataLakeStore](#DataLakeStore), [DataLakeStoreStatus](#DataLakeStoreStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `encryptionState` | ***string***| ***(Optional)*** |
| `encryptionType` | ***string***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `firewallAllowAzureIPS` | ***string***| ***(Optional)*** |
| `firewallState` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `tier` | ***string***| ***(Optional)*** |
## DataLakeStoreStatus
##### (Appears on:[DataLakeStore](#DataLakeStore))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataLakeStoreSpec](#DataLakeStoreSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
