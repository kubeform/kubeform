---
title: KeyVault
menu:
  docs_{{ .version }}:
    identifier: keyvault-azurerm.kubeform.com
    name: KeyVault
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KeyVault
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `KeyVault` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KeyVaultSpec](#keyvaultspec)***||
| `status` | ***[KeyVaultStatus](#keyvaultstatus)***||
## KeyVaultSpec

Appears on:[KeyVault](#keyvault), [KeyVaultStatus](#keyvaultstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accessPolicy` | ***[[]KeyVaultSpecAccessPolicy](#keyvaultspecaccesspolicy)***| ***(Optional)*** |
| `enabledForDeployment` | ***bool***| ***(Optional)*** |
| `enabledForDiskEncryption` | ***bool***| ***(Optional)*** |
| `enabledForTemplateDeployment` | ***bool***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `networkAcls` | ***[[]KeyVaultSpecNetworkAcls](#keyvaultspecnetworkacls)***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `sku` | ***[[]KeyVaultSpecSku](#keyvaultspecsku)***| ***(Optional)*** Deprecated|
| `skuName` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `tenantID` | ***string***||
| `vaultURI` | ***string***| ***(Optional)*** |
## KeyVaultSpecAccessPolicy

Appears on:[KeyVaultSpec](#keyvaultspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `applicationID` | ***string***| ***(Optional)*** |
| `certificatePermissions` | ***[]string***| ***(Optional)*** |
| `keyPermissions` | ***[]string***| ***(Optional)*** |
| `objectID` | ***string***||
| `secretPermissions` | ***[]string***| ***(Optional)*** |
| `storagePermissions` | ***[]string***| ***(Optional)*** |
| `tenantID` | ***string***||
## KeyVaultSpecNetworkAcls

Appears on:[KeyVaultSpec](#keyvaultspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bypass` | ***string***||
| `defaultAction` | ***string***||
| `ipRules` | ***[]string***| ***(Optional)*** |
| `virtualNetworkSubnetIDS` | ***[]string***| ***(Optional)*** |
## KeyVaultSpecSku

Appears on:[KeyVaultSpec](#keyvaultspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
## KeyVaultStatus

Appears on:[KeyVault](#keyvault)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KeyVaultSpec](#keyvaultspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[KeyVaultStatus](#keyvaultstatus)

---
