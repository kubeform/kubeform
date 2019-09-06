## BatchAccount
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `BatchAccount` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BatchAccountSpec](#BatchAccountSpec)***||
| `status` | ***[BatchAccountStatus](#BatchAccountStatus)***||
## BatchAccountSpec
##### (Appears on:[BatchAccount](#BatchAccount), [BatchAccountStatus](#BatchAccountStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `accountEndpoint` | ***string***| ***(Optional)*** |
| `keyVaultReference` | ***[[]BatchAccountSpecKeyVaultReference](#BatchAccountSpecKeyVaultReference)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `poolAllocationMode` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `storageAccountID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## BatchAccountSpecKeyVaultReference
##### (Appears on:[BatchAccountSpec](#BatchAccountSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***||
| `url` | ***string***||
## BatchAccountStatus
##### (Appears on:[BatchAccount](#BatchAccount))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BatchAccountSpec](#BatchAccountSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_access_key` | ***string*** ||
| `secondary_access_key` | ***string*** ||
