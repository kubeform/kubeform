---
title: BackupPolicyVm
menu:
  docs_{{ .version }}:
    identifier: backuppolicyvm-azurerm.kubeform.com
    name: BackupPolicyVm
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BackupPolicyVm
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `BackupPolicyVm` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BackupPolicyVmSpec](#backuppolicyvmspec)***||
| `status` | ***[BackupPolicyVmStatus](#backuppolicyvmstatus)***||
## BackupPolicyVmSpec

Appears on:[BackupPolicyVm](#backuppolicyvm), [BackupPolicyVmStatus](#backuppolicyvmstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backup` | ***[[]BackupPolicyVmSpecBackup](#backuppolicyvmspecbackup)***||
| `name` | ***string***||
| `recoveryVaultName` | ***string***||
| `resourceGroupName` | ***string***||
| `retentionDaily` | ***[[]BackupPolicyVmSpecRetentionDaily](#backuppolicyvmspecretentiondaily)***| ***(Optional)*** |
| `retentionMonthly` | ***[[]BackupPolicyVmSpecRetentionMonthly](#backuppolicyvmspecretentionmonthly)***| ***(Optional)*** |
| `retentionWeekly` | ***[[]BackupPolicyVmSpecRetentionWeekly](#backuppolicyvmspecretentionweekly)***| ***(Optional)*** |
| `retentionYearly` | ***[[]BackupPolicyVmSpecRetentionYearly](#backuppolicyvmspecretentionyearly)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `timezone` | ***string***| ***(Optional)*** |
## BackupPolicyVmSpecBackup

Appears on:[BackupPolicyVmSpec](#backuppolicyvmspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `frequency` | ***string***||
| `time` | ***string***||
| `weekdays` | ***[]string***| ***(Optional)*** |
## BackupPolicyVmSpecRetentionDaily

Appears on:[BackupPolicyVmSpec](#backuppolicyvmspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int64***||
## BackupPolicyVmSpecRetentionMonthly

Appears on:[BackupPolicyVmSpec](#backuppolicyvmspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int64***||
| `weekdays` | ***[]string***||
| `weeks` | ***[]string***||
## BackupPolicyVmSpecRetentionWeekly

Appears on:[BackupPolicyVmSpec](#backuppolicyvmspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int64***||
| `weekdays` | ***[]string***||
## BackupPolicyVmSpecRetentionYearly

Appears on:[BackupPolicyVmSpec](#backuppolicyvmspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int64***||
| `months` | ***[]string***||
| `weekdays` | ***[]string***||
| `weeks` | ***[]string***||
## BackupPolicyVmStatus

Appears on:[BackupPolicyVm](#backuppolicyvm)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BackupPolicyVmSpec](#backuppolicyvmspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BackupPolicyVmStatus](#backuppolicyvmstatus)

---
