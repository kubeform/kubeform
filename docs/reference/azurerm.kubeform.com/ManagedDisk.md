---
title: ManagedDisk
menu:
  docs_{{ .version }}:
    identifier: manageddisk-azurerm.kubeform.com
    name: ManagedDisk
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ManagedDisk
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ManagedDisk` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ManagedDiskSpec](#manageddiskspec)***||
| `status` | ***[ManagedDiskStatus](#manageddiskstatus)***||
## ManagedDiskSpec

Appears on:[ManagedDisk](#manageddisk), [ManagedDiskStatus](#manageddiskstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `createOption` | ***string***||
| `diskSizeGb` | ***int***| ***(Optional)*** |
| `encryptionSettings` | ***[[]ManagedDiskSpecEncryptionSettings](#manageddiskspecencryptionsettings)***| ***(Optional)*** |
| `imageReferenceID` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `osType` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `sourceResourceID` | ***string***| ***(Optional)*** |
| `sourceURI` | ***string***| ***(Optional)*** |
| `storageAccountType` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `zones` | ***[]string***| ***(Optional)*** |
## ManagedDiskSpecEncryptionSettings

Appears on:[ManagedDiskSpec](#manageddiskspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskEncryptionKey` | ***[[]ManagedDiskSpecEncryptionSettingsDiskEncryptionKey](#manageddiskspecencryptionsettingsdiskencryptionkey)***| ***(Optional)*** |
| `enabled` | ***bool***||
| `keyEncryptionKey` | ***[[]ManagedDiskSpecEncryptionSettingsKeyEncryptionKey](#manageddiskspecencryptionsettingskeyencryptionkey)***| ***(Optional)*** |
## ManagedDiskSpecEncryptionSettingsDiskEncryptionKey

Appears on:[ManagedDiskSpecEncryptionSettings](#manageddiskspecencryptionsettings)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `secretURL` | ***string***||
| `sourceVaultID` | ***string***||
## ManagedDiskSpecEncryptionSettingsKeyEncryptionKey

Appears on:[ManagedDiskSpecEncryptionSettings](#manageddiskspecencryptionsettings)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `keyURL` | ***string***||
| `sourceVaultID` | ***string***||
## ManagedDiskStatus

Appears on:[ManagedDisk](#manageddisk)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ManagedDiskSpec](#manageddiskspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
