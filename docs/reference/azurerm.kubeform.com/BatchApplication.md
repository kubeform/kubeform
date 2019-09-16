---
title: BatchApplication
menu:
  docs_{{ .version }}:
    identifier: batchapplication-azurerm.kubeform.com
    name: BatchApplication
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BatchApplication
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `BatchApplication` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BatchApplicationSpec](#batchapplicationspec)***||
| `status` | ***[BatchApplicationStatus](#batchapplicationstatus)***||
## BatchApplicationSpec

Appears on:[BatchApplication](#batchapplication), [BatchApplicationStatus](#batchapplicationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***||
| `allowUpdates` | ***bool***| ***(Optional)*** |
| `defaultVersion` | ***string***| ***(Optional)*** |
| `displayName` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
## BatchApplicationStatus

Appears on:[BatchApplication](#batchapplication)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BatchApplicationSpec](#batchapplicationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
