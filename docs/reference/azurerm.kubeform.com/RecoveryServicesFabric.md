---
title: RecoveryServicesFabric
menu:
  docs_{{ .version }}:
    identifier: recoveryservicesfabric-azurerm.kubeform.com
    name: RecoveryServicesFabric
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RecoveryServicesFabric
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `RecoveryServicesFabric` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RecoveryServicesFabricSpec](#recoveryservicesfabricspec)***||
| `status` | ***[RecoveryServicesFabricStatus](#recoveryservicesfabricstatus)***||
## Phase(`string` alias)

Appears on:[RecoveryServicesFabricStatus](#recoveryservicesfabricstatus)

## RecoveryServicesFabricSpec

Appears on:[RecoveryServicesFabric](#recoveryservicesfabric), [RecoveryServicesFabricStatus](#recoveryservicesfabricstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `recoveryVaultName` | ***string***||
| `resourceGroupName` | ***string***||
## RecoveryServicesFabricStatus

Appears on:[RecoveryServicesFabric](#recoveryservicesfabric)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RecoveryServicesFabricSpec](#recoveryservicesfabricspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
