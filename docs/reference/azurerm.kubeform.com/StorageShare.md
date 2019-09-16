---
title: StorageShare
menu:
  docs_{{ .version }}:
    identifier: storageshare-azurerm.kubeform.com
    name: StorageShare
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageShare
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageShare` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageShareSpec](#storagesharespec)***||
| `status` | ***[StorageShareStatus](#storagesharestatus)***||
## StorageShareSpec

Appears on:[StorageShare](#storageshare), [StorageShareStatus](#storagesharestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `acl` | ***[[]StorageShareSpecAcl](#storagesharespecacl)***| ***(Optional)*** |
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `name` | ***string***||
| `quota` | ***int***| ***(Optional)*** |
| `resourceGroupName` | ***string***| ***(Optional)*** Deprecated|
| `storageAccountName` | ***string***||
| `url` | ***string***| ***(Optional)*** |
## StorageShareSpecAcl

Appears on:[StorageShareSpec](#storagesharespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `accessPolicy` | ***[[]StorageShareSpecAclAccessPolicy](#storagesharespecaclaccesspolicy)***| ***(Optional)*** |
| `ID` | ***string***||
## StorageShareSpecAclAccessPolicy

Appears on:[StorageShareSpecAcl](#storagesharespecacl)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `expiry` | ***string***||
| `permissions` | ***string***||
| `start` | ***string***||
## StorageShareStatus

Appears on:[StorageShare](#storageshare)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageShareSpec](#storagesharespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
