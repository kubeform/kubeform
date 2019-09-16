---
title: SsmMaintenanceWindowTarget
menu:
  docs_{{ .version }}:
    identifier: ssmmaintenancewindowtarget-aws.kubeform.com
    name: SsmMaintenanceWindowTarget
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SsmMaintenanceWindowTarget
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SsmMaintenanceWindowTarget` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SsmMaintenanceWindowTargetSpec](#ssmmaintenancewindowtargetspec)***||
| `status` | ***[SsmMaintenanceWindowTargetStatus](#ssmmaintenancewindowtargetstatus)***||
## SsmMaintenanceWindowTargetSpec

Appears on:[SsmMaintenanceWindowTarget](#ssmmaintenancewindowtarget), [SsmMaintenanceWindowTargetStatus](#ssmmaintenancewindowtargetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `ownerInformation` | ***string***| ***(Optional)*** |
| `resourceType` | ***string***||
| `targets` | ***[[]SsmMaintenanceWindowTargetSpecTargets](#ssmmaintenancewindowtargetspectargets)***||
| `windowID` | ***string***||
## SsmMaintenanceWindowTargetSpecTargets

Appears on:[SsmMaintenanceWindowTargetSpec](#ssmmaintenancewindowtargetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***||
| `values` | ***[]string***||
## SsmMaintenanceWindowTargetStatus

Appears on:[SsmMaintenanceWindowTarget](#ssmmaintenancewindowtarget)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SsmMaintenanceWindowTargetSpec](#ssmmaintenancewindowtargetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
