## BackupVault
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `BackupVault` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BackupVaultSpec](#BackupVaultSpec)***||
| `status` | ***[BackupVaultStatus](#BackupVaultStatus)***||
## BackupVaultSpec
##### (Appears on:[BackupVault](#BackupVault), [BackupVaultStatus](#BackupVaultStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `recoveryPoints` | ***int***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## BackupVaultStatus
##### (Appears on:[BackupVault](#BackupVault))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BackupVaultSpec](#BackupVaultSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
