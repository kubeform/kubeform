---
title: KeyVaultKey
menu:
  docs_{{ .version }}:
    identifier: keyvaultkey-azurerm.kubeform.com
    name: KeyVaultKey
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KeyVaultKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `KeyVaultKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KeyVaultKeySpec](#keyvaultkeyspec)***||
| `status` | ***[KeyVaultKeyStatus](#keyvaultkeystatus)***||
## KeyVaultKeySpec

Appears on:[KeyVaultKey](#keyvaultkey), [KeyVaultKeyStatus](#keyvaultkeystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `curve` | ***string***| ***(Optional)*** |
| `e` | ***string***| ***(Optional)*** |
| `expirationDate` | ***string***| ***(Optional)*** |
| `keyOpts` | ***[]string***||
| `keySize` | ***int64***| ***(Optional)*** |
| `keyType` | ***string***||
| `keyVaultID` | ***string***| ***(Optional)*** |
| `n` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `notBeforeDate` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vaultURI` | ***string***| ***(Optional)*** Deprecated|
| `version` | ***string***| ***(Optional)*** |
| `x` | ***string***| ***(Optional)*** |
| `y` | ***string***| ***(Optional)*** |
## KeyVaultKeyStatus

Appears on:[KeyVaultKey](#keyvaultkey)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KeyVaultKeySpec](#keyvaultkeyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[KeyVaultKeyStatus](#keyvaultkeystatus)

---
