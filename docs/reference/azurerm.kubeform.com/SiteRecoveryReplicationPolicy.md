---
title: SiteRecoveryReplicationPolicy
menu:
  docs_{{ .version }}:
    identifier: siterecoveryreplicationpolicy-azurerm.kubeform.com
    name: SiteRecoveryReplicationPolicy
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SiteRecoveryReplicationPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SiteRecoveryReplicationPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SiteRecoveryReplicationPolicySpec](#siterecoveryreplicationpolicyspec)***||
| `status` | ***[SiteRecoveryReplicationPolicyStatus](#siterecoveryreplicationpolicystatus)***||
## Phase(`string` alias)

Appears on:[SiteRecoveryReplicationPolicyStatus](#siterecoveryreplicationpolicystatus)

## SiteRecoveryReplicationPolicySpec

Appears on:[SiteRecoveryReplicationPolicy](#siterecoveryreplicationpolicy), [SiteRecoveryReplicationPolicyStatus](#siterecoveryreplicationpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `applicationConsistentSnapshotFrequencyInMinutes` | ***int64***||
| `name` | ***string***||
| `recoveryPointRetentionInMinutes` | ***int64***||
| `recoveryVaultName` | ***string***||
| `resourceGroupName` | ***string***||
## SiteRecoveryReplicationPolicyStatus

Appears on:[SiteRecoveryReplicationPolicy](#siterecoveryreplicationpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SiteRecoveryReplicationPolicySpec](#siterecoveryreplicationpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
