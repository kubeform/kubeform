---
title: RecoveryServicesProtectionPolicyVm
menu:
  docs_v0.0.1:
    identifier: recoveryservicesprotectionpolicyvm-azurerm.kubeform.com
    name: RecoveryServicesProtectionPolicyVm
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## RecoveryServicesProtectionPolicyVm
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `RecoveryServicesProtectionPolicyVm` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RecoveryServicesProtectionPolicyVmSpec](#RecoveryServicesProtectionPolicyVmSpec)***||
| `status` | ***[RecoveryServicesProtectionPolicyVmStatus](#RecoveryServicesProtectionPolicyVmStatus)***||
## RecoveryServicesProtectionPolicyVmSpec
##### (Appears on:[RecoveryServicesProtectionPolicyVm](#RecoveryServicesProtectionPolicyVm), [RecoveryServicesProtectionPolicyVmStatus](#RecoveryServicesProtectionPolicyVmStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backup` | ***[[]RecoveryServicesProtectionPolicyVmSpecBackup](#RecoveryServicesProtectionPolicyVmSpecBackup)***||
| `name` | ***string***||
| `recoveryVaultName` | ***string***||
| `resourceGroupName` | ***string***||
| `retentionDaily` | ***[[]RecoveryServicesProtectionPolicyVmSpecRetentionDaily](#RecoveryServicesProtectionPolicyVmSpecRetentionDaily)***| ***(Optional)*** |
| `retentionMonthly` | ***[[]RecoveryServicesProtectionPolicyVmSpecRetentionMonthly](#RecoveryServicesProtectionPolicyVmSpecRetentionMonthly)***| ***(Optional)*** |
| `retentionWeekly` | ***[[]RecoveryServicesProtectionPolicyVmSpecRetentionWeekly](#RecoveryServicesProtectionPolicyVmSpecRetentionWeekly)***| ***(Optional)*** |
| `retentionYearly` | ***[[]RecoveryServicesProtectionPolicyVmSpecRetentionYearly](#RecoveryServicesProtectionPolicyVmSpecRetentionYearly)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `timezone` | ***string***| ***(Optional)*** |
## RecoveryServicesProtectionPolicyVmSpecBackup
##### (Appears on:[RecoveryServicesProtectionPolicyVmSpec](#RecoveryServicesProtectionPolicyVmSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `frequency` | ***string***||
| `time` | ***string***||
| `weekdays` | ***[]string***| ***(Optional)*** |
## RecoveryServicesProtectionPolicyVmSpecRetentionDaily
##### (Appears on:[RecoveryServicesProtectionPolicyVmSpec](#RecoveryServicesProtectionPolicyVmSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***||
## RecoveryServicesProtectionPolicyVmSpecRetentionMonthly
##### (Appears on:[RecoveryServicesProtectionPolicyVmSpec](#RecoveryServicesProtectionPolicyVmSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***||
| `weekdays` | ***[]string***||
| `weeks` | ***[]string***||
## RecoveryServicesProtectionPolicyVmSpecRetentionWeekly
##### (Appears on:[RecoveryServicesProtectionPolicyVmSpec](#RecoveryServicesProtectionPolicyVmSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***||
| `weekdays` | ***[]string***||
## RecoveryServicesProtectionPolicyVmSpecRetentionYearly
##### (Appears on:[RecoveryServicesProtectionPolicyVmSpec](#RecoveryServicesProtectionPolicyVmSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***||
| `months` | ***[]string***||
| `weekdays` | ***[]string***||
| `weeks` | ***[]string***||
## RecoveryServicesProtectionPolicyVmStatus
##### (Appears on:[RecoveryServicesProtectionPolicyVm](#RecoveryServicesProtectionPolicyVm))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RecoveryServicesProtectionPolicyVmSpec](#RecoveryServicesProtectionPolicyVmSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
