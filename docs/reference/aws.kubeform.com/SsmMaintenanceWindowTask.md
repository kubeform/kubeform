---
title: SsmMaintenanceWindowTask
menu:
  docs_{{ .version }}:
    identifier: ssmmaintenancewindowtask-aws.kubeform.com
    name: SsmMaintenanceWindowTask
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SsmMaintenanceWindowTask
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SsmMaintenanceWindowTask` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SsmMaintenanceWindowTaskSpec](#ssmmaintenancewindowtaskspec)***||
| `status` | ***[SsmMaintenanceWindowTaskStatus](#ssmmaintenancewindowtaskstatus)***||
## Phase(`string` alias)

Appears on:[SsmMaintenanceWindowTaskStatus](#ssmmaintenancewindowtaskstatus)

## SsmMaintenanceWindowTaskSpec

Appears on:[SsmMaintenanceWindowTask](#ssmmaintenancewindowtask), [SsmMaintenanceWindowTaskStatus](#ssmmaintenancewindowtaskstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `loggingInfo` | ***[[]SsmMaintenanceWindowTaskSpecLoggingInfo](#ssmmaintenancewindowtaskspeclogginginfo)***| ***(Optional)*** |
| `maxConcurrency` | ***string***||
| `maxErrors` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `priority` | ***int***| ***(Optional)*** |
| `serviceRoleArn` | ***string***||
| `targets` | ***[[]SsmMaintenanceWindowTaskSpecTargets](#ssmmaintenancewindowtaskspectargets)***||
| `taskArn` | ***string***||
| `taskParameters` | ***[[]SsmMaintenanceWindowTaskSpecTaskParameters](#ssmmaintenancewindowtaskspectaskparameters)***| ***(Optional)*** |
| `taskType` | ***string***||
| `windowID` | ***string***||
## SsmMaintenanceWindowTaskSpecLoggingInfo

Appears on:[SsmMaintenanceWindowTaskSpec](#ssmmaintenancewindowtaskspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `s3BucketName` | ***string***||
| `s3BucketPrefix` | ***string***| ***(Optional)*** |
| `s3Region` | ***string***||
## SsmMaintenanceWindowTaskSpecTargets

Appears on:[SsmMaintenanceWindowTaskSpec](#ssmmaintenancewindowtaskspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***||
| `values` | ***[]string***||
## SsmMaintenanceWindowTaskSpecTaskParameters

Appears on:[SsmMaintenanceWindowTaskSpec](#ssmmaintenancewindowtaskspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `values` | ***[]string***||
## SsmMaintenanceWindowTaskStatus

Appears on:[SsmMaintenanceWindowTask](#ssmmaintenancewindowtask)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SsmMaintenanceWindowTaskSpec](#ssmmaintenancewindowtaskspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
