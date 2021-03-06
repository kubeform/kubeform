---
title: RecoveryServicesProtectionContainerMapping
menu:
  docs_{{ .version }}:
    identifier: recoveryservicesprotectioncontainermapping-azurerm.kubeform.com
    name: RecoveryServicesProtectionContainerMapping
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RecoveryServicesProtectionContainerMapping
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `RecoveryServicesProtectionContainerMapping` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RecoveryServicesProtectionContainerMappingSpec](#recoveryservicesprotectioncontainermappingspec)***||
| `status` | ***[RecoveryServicesProtectionContainerMappingStatus](#recoveryservicesprotectioncontainermappingstatus)***||
## Phase(`string` alias)

Appears on:[RecoveryServicesProtectionContainerMappingStatus](#recoveryservicesprotectioncontainermappingstatus)

## RecoveryServicesProtectionContainerMappingSpec

Appears on:[RecoveryServicesProtectionContainerMapping](#recoveryservicesprotectioncontainermapping), [RecoveryServicesProtectionContainerMappingStatus](#recoveryservicesprotectioncontainermappingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `recoveryFabricName` | ***string***||
| `recoveryReplicationPolicyID` | ***string***||
| `recoverySourceProtectionContainerName` | ***string***||
| `recoveryTargetProtectionContainerID` | ***string***||
| `recoveryVaultName` | ***string***||
| `resourceGroupName` | ***string***||
## RecoveryServicesProtectionContainerMappingStatus

Appears on:[RecoveryServicesProtectionContainerMapping](#recoveryservicesprotectioncontainermapping)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RecoveryServicesProtectionContainerMappingSpec](#recoveryservicesprotectioncontainermappingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
