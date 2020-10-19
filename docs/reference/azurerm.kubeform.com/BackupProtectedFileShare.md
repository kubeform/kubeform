---
title: BackupProtectedFileShare
menu:
  docs_{{ .version }}:
    identifier: backupprotectedfileshare-azurerm.kubeform.com
    name: BackupProtectedFileShare
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BackupProtectedFileShare
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `BackupProtectedFileShare` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BackupProtectedFileShareSpec](#backupprotectedfilesharespec)***||
| `status` | ***[BackupProtectedFileShareStatus](#backupprotectedfilesharestatus)***||
## BackupProtectedFileShareSpec

Appears on:[BackupProtectedFileShare](#backupprotectedfileshare), [BackupProtectedFileShareStatus](#backupprotectedfilesharestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backupPolicyID` | ***string***||
| `recoveryVaultName` | ***string***||
| `resourceGroupName` | ***string***||
| `sourceFileShareName` | ***string***||
| `sourceStorageAccountID` | ***string***||
## BackupProtectedFileShareStatus

Appears on:[BackupProtectedFileShare](#backupprotectedfileshare)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BackupProtectedFileShareSpec](#backupprotectedfilesharespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BackupProtectedFileShareStatus](#backupprotectedfilesharestatus)

---
