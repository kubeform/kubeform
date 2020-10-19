---
title: DiskEncryptionSet
menu:
  docs_{{ .version }}:
    identifier: diskencryptionset-azurerm.kubeform.com
    name: DiskEncryptionSet
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DiskEncryptionSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DiskEncryptionSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DiskEncryptionSetSpec](#diskencryptionsetspec)***||
| `status` | ***[DiskEncryptionSetStatus](#diskencryptionsetstatus)***||
## DiskEncryptionSetSpec

Appears on:[DiskEncryptionSet](#diskencryptionset), [DiskEncryptionSetStatus](#diskencryptionsetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `identity` | ***[[]DiskEncryptionSetSpecIdentity](#diskencryptionsetspecidentity)***||
| `keyVaultKeyID` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## DiskEncryptionSetSpecIdentity

Appears on:[DiskEncryptionSetSpec](#diskencryptionsetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `principalID` | ***string***| ***(Optional)*** |
| `tenantID` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## DiskEncryptionSetStatus

Appears on:[DiskEncryptionSet](#diskencryptionset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DiskEncryptionSetSpec](#diskencryptionsetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DiskEncryptionSetStatus](#diskencryptionsetstatus)

---
