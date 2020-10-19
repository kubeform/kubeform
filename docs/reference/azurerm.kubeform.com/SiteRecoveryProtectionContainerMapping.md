---
title: SiteRecoveryProtectionContainerMapping
menu:
  docs_{{ .version }}:
    identifier: siterecoveryprotectioncontainermapping-azurerm.kubeform.com
    name: SiteRecoveryProtectionContainerMapping
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SiteRecoveryProtectionContainerMapping
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SiteRecoveryProtectionContainerMapping` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SiteRecoveryProtectionContainerMappingSpec](#siterecoveryprotectioncontainermappingspec)***||
| `status` | ***[SiteRecoveryProtectionContainerMappingStatus](#siterecoveryprotectioncontainermappingstatus)***||
## Phase(`string` alias)

Appears on:[SiteRecoveryProtectionContainerMappingStatus](#siterecoveryprotectioncontainermappingstatus)

## SiteRecoveryProtectionContainerMappingSpec

Appears on:[SiteRecoveryProtectionContainerMapping](#siterecoveryprotectioncontainermapping), [SiteRecoveryProtectionContainerMappingStatus](#siterecoveryprotectioncontainermappingstatus)

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
## SiteRecoveryProtectionContainerMappingStatus

Appears on:[SiteRecoveryProtectionContainerMapping](#siterecoveryprotectioncontainermapping)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SiteRecoveryProtectionContainerMappingSpec](#siterecoveryprotectioncontainermappingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
