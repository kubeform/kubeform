---
title: RecoveryServicesReplicationPolicy
menu:
  docs_{{ .version }}:
    identifier: recoveryservicesreplicationpolicy-azurerm.kubeform.com
    name: RecoveryServicesReplicationPolicy
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RecoveryServicesReplicationPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `RecoveryServicesReplicationPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RecoveryServicesReplicationPolicySpec](#recoveryservicesreplicationpolicyspec)***||
| `status` | ***[RecoveryServicesReplicationPolicyStatus](#recoveryservicesreplicationpolicystatus)***||
## Phase(`string` alias)

Appears on:[RecoveryServicesReplicationPolicyStatus](#recoveryservicesreplicationpolicystatus)

## RecoveryServicesReplicationPolicySpec

Appears on:[RecoveryServicesReplicationPolicy](#recoveryservicesreplicationpolicy), [RecoveryServicesReplicationPolicyStatus](#recoveryservicesreplicationpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `applicationConsistentSnapshotFrequencyInMinutes` | ***int64***||
| `name` | ***string***||
| `recoveryPointRetentionInMinutes` | ***int64***||
| `recoveryVaultName` | ***string***||
| `resourceGroupName` | ***string***||
## RecoveryServicesReplicationPolicyStatus

Appears on:[RecoveryServicesReplicationPolicy](#recoveryservicesreplicationpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RecoveryServicesReplicationPolicySpec](#recoveryservicesreplicationpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
