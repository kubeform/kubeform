## GlacierVaultLock
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GlacierVaultLock` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GlacierVaultLockSpec](#GlacierVaultLockSpec)***||
| `status` | ***[GlacierVaultLockStatus](#GlacierVaultLockStatus)***||
## GlacierVaultLockSpec
##### (Appears on:[GlacierVaultLock](#GlacierVaultLock), [GlacierVaultLockStatus](#GlacierVaultLockStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `completeLock` | ***bool***||
| `ignoreDeletionError` | ***bool***| ***(Optional)*** |
| `policy` | ***string***||
| `vaultName` | ***string***||
## GlacierVaultLockStatus
##### (Appears on:[GlacierVaultLock](#GlacierVaultLock))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlacierVaultLockSpec](#GlacierVaultLockSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
