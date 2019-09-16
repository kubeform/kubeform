---
title: RecoveryServicesProtectedVm
menu:
  docs_{{ .version }}:
    identifier: recoveryservicesprotectedvm-azurerm.kubeform.com
    name: RecoveryServicesProtectedVm
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RecoveryServicesProtectedVm
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `RecoveryServicesProtectedVm` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RecoveryServicesProtectedVmSpec](#recoveryservicesprotectedvmspec)***||
| `status` | ***[RecoveryServicesProtectedVmStatus](#recoveryservicesprotectedvmstatus)***||
## RecoveryServicesProtectedVmSpec

Appears on:[RecoveryServicesProtectedVm](#recoveryservicesprotectedvm), [RecoveryServicesProtectedVmStatus](#recoveryservicesprotectedvmstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backupPolicyID` | ***string***||
| `recoveryVaultName` | ***string***||
| `resourceGroupName` | ***string***||
| `sourceVmID` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## RecoveryServicesProtectedVmStatus

Appears on:[RecoveryServicesProtectedVm](#recoveryservicesprotectedvm)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RecoveryServicesProtectedVmSpec](#recoveryservicesprotectedvmspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---