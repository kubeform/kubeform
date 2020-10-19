---
title: BackupPolicyFileShare
menu:
  docs_{{ .version }}:
    identifier: backuppolicyfileshare-azurerm.kubeform.com
    name: BackupPolicyFileShare
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BackupPolicyFileShare
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `BackupPolicyFileShare` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BackupPolicyFileShareSpec](#backuppolicyfilesharespec)***||
| `status` | ***[BackupPolicyFileShareStatus](#backuppolicyfilesharestatus)***||
## BackupPolicyFileShareSpec

Appears on:[BackupPolicyFileShare](#backuppolicyfileshare), [BackupPolicyFileShareStatus](#backuppolicyfilesharestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backup` | ***[[]BackupPolicyFileShareSpecBackup](#backuppolicyfilesharespecbackup)***||
| `name` | ***string***||
| `recoveryVaultName` | ***string***||
| `resourceGroupName` | ***string***||
| `retentionDaily` | ***[[]BackupPolicyFileShareSpecRetentionDaily](#backuppolicyfilesharespecretentiondaily)***||
| `timezone` | ***string***| ***(Optional)*** |
## BackupPolicyFileShareSpecBackup

Appears on:[BackupPolicyFileShareSpec](#backuppolicyfilesharespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `frequency` | ***string***||
| `time` | ***string***||
## BackupPolicyFileShareSpecRetentionDaily

Appears on:[BackupPolicyFileShareSpec](#backuppolicyfilesharespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int64***||
## BackupPolicyFileShareStatus

Appears on:[BackupPolicyFileShare](#backuppolicyfileshare)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BackupPolicyFileShareSpec](#backuppolicyfilesharespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BackupPolicyFileShareStatus](#backuppolicyfilesharestatus)

---
