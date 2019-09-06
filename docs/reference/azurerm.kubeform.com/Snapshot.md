## Snapshot
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `Snapshot` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SnapshotSpec](#SnapshotSpec)***||
| `status` | ***[SnapshotStatus](#SnapshotStatus)***||
## SnapshotSpec
##### (Appears on:[Snapshot](#Snapshot), [SnapshotStatus](#SnapshotStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `createOption` | ***string***||
| `diskSizeGb` | ***int***| ***(Optional)*** |
| `encryptionSettings` | ***[[]SnapshotSpecEncryptionSettings](#SnapshotSpecEncryptionSettings)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sourceResourceID` | ***string***| ***(Optional)*** |
| `sourceURI` | ***string***| ***(Optional)*** |
| `storageAccountID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## SnapshotSpecEncryptionSettings
##### (Appears on:[SnapshotSpec](#SnapshotSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskEncryptionKey` | ***[[]SnapshotSpecEncryptionSettingsDiskEncryptionKey](#SnapshotSpecEncryptionSettingsDiskEncryptionKey)***| ***(Optional)*** |
| `enabled` | ***bool***||
| `keyEncryptionKey` | ***[[]SnapshotSpecEncryptionSettingsKeyEncryptionKey](#SnapshotSpecEncryptionSettingsKeyEncryptionKey)***| ***(Optional)*** |
## SnapshotSpecEncryptionSettingsDiskEncryptionKey
##### (Appears on:[SnapshotSpecEncryptionSettings](#SnapshotSpecEncryptionSettings))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `secretURL` | ***string***||
| `sourceVaultID` | ***string***||
## SnapshotSpecEncryptionSettingsKeyEncryptionKey
##### (Appears on:[SnapshotSpecEncryptionSettings](#SnapshotSpecEncryptionSettings))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `keyURL` | ***string***||
| `sourceVaultID` | ***string***||
## SnapshotStatus
##### (Appears on:[Snapshot](#Snapshot))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SnapshotSpec](#SnapshotSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
