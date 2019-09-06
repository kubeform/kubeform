## CognitiveAccount
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `CognitiveAccount` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CognitiveAccountSpec](#CognitiveAccountSpec)***||
| `status` | ***[CognitiveAccountStatus](#CognitiveAccountStatus)***||
## CognitiveAccountSpec
##### (Appears on:[CognitiveAccount](#CognitiveAccount), [CognitiveAccountStatus](#CognitiveAccountStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `endpoint` | ***string***| ***(Optional)*** |
| `kind` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sku` | ***[[]CognitiveAccountSpecSku](#CognitiveAccountSpecSku)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## CognitiveAccountSpecSku
##### (Appears on:[CognitiveAccountSpec](#CognitiveAccountSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `tier` | ***string***||
## CognitiveAccountStatus
##### (Appears on:[CognitiveAccount](#CognitiveAccount))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CognitiveAccountSpec](#CognitiveAccountSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_access_key` | ***string*** ||
| `secondary_access_key` | ***string*** ||
