---
title: RecoveryNetworkMapping
menu:
  docs_{{ .version }}:
    identifier: recoverynetworkmapping-azurerm.kubeform.com
    name: RecoveryNetworkMapping
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RecoveryNetworkMapping
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `RecoveryNetworkMapping` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RecoveryNetworkMappingSpec](#recoverynetworkmappingspec)***||
| `status` | ***[RecoveryNetworkMappingStatus](#recoverynetworkmappingstatus)***||
## Phase(`string` alias)

Appears on:[RecoveryNetworkMappingStatus](#recoverynetworkmappingstatus)

## RecoveryNetworkMappingSpec

Appears on:[RecoveryNetworkMapping](#recoverynetworkmapping), [RecoveryNetworkMappingStatus](#recoverynetworkmappingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `recoveryVaultName` | ***string***||
| `resourceGroupName` | ***string***||
| `sourceNetworkID` | ***string***||
| `sourceRecoveryFabricName` | ***string***||
| `targetNetworkID` | ***string***||
| `targetRecoveryFabricName` | ***string***||
## RecoveryNetworkMappingStatus

Appears on:[RecoveryNetworkMapping](#recoverynetworkmapping)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RecoveryNetworkMappingSpec](#recoverynetworkmappingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
