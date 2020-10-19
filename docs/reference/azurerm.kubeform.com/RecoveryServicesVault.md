---
title: RecoveryServicesVault
menu:
  docs_{{ .version }}:
    identifier: recoveryservicesvault-azurerm.kubeform.com
    name: RecoveryServicesVault
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RecoveryServicesVault
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `RecoveryServicesVault` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RecoveryServicesVaultSpec](#recoveryservicesvaultspec)***||
| `status` | ***[RecoveryServicesVaultStatus](#recoveryservicesvaultstatus)***||
## Phase(`string` alias)

Appears on:[RecoveryServicesVaultStatus](#recoveryservicesvaultstatus)

## RecoveryServicesVaultSpec

Appears on:[RecoveryServicesVault](#recoveryservicesvault), [RecoveryServicesVaultStatus](#recoveryservicesvaultstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sku` | ***string***||
| `softDeleteEnabled` | ***bool***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## RecoveryServicesVaultStatus

Appears on:[RecoveryServicesVault](#recoveryservicesvault)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RecoveryServicesVaultSpec](#recoveryservicesvaultspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
