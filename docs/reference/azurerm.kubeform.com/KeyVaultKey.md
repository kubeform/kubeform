## KeyVaultKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `KeyVaultKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KeyVaultKeySpec](#KeyVaultKeySpec)***||
| `status` | ***[KeyVaultKeyStatus](#KeyVaultKeyStatus)***||
## KeyVaultKeySpec
##### (Appears on:[KeyVaultKey](#KeyVaultKey), [KeyVaultKeyStatus](#KeyVaultKeyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `curve` | ***string***| ***(Optional)*** |
| `e` | ***string***| ***(Optional)*** |
| `keyOpts` | ***[]string***||
| `keySize` | ***int***| ***(Optional)*** |
| `keyType` | ***string***||
| `keyVaultID` | ***string***| ***(Optional)*** |
| `n` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vaultURI` | ***string***| ***(Optional)*** Deprecated|
| `version` | ***string***| ***(Optional)*** |
| `x` | ***string***| ***(Optional)*** |
| `y` | ***string***| ***(Optional)*** |
## KeyVaultKeyStatus
##### (Appears on:[KeyVaultKey](#KeyVaultKey))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KeyVaultKeySpec](#KeyVaultKeySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
