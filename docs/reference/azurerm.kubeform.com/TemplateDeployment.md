---
title: TemplateDeployment
menu:
  docs_v0.0.1:
    identifier: templatedeployment-azurerm.kubeform.com
    name: TemplateDeployment
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## TemplateDeployment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `TemplateDeployment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[TemplateDeploymentSpec](#TemplateDeploymentSpec)***||
| `status` | ***[TemplateDeploymentStatus](#TemplateDeploymentStatus)***||
## TemplateDeploymentSpec
##### (Appears on:[TemplateDeployment](#TemplateDeployment), [TemplateDeploymentStatus](#TemplateDeploymentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `deploymentMode` | ***string***||
| `name` | ***string***||
| `outputs` | ***map[string]string***| ***(Optional)*** |
| `parameters` | ***map[string]string***| ***(Optional)*** |
| `parametersBody` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `templateBody` | ***string***| ***(Optional)*** |
## TemplateDeploymentStatus
##### (Appears on:[TemplateDeployment](#TemplateDeployment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[TemplateDeploymentSpec](#TemplateDeploymentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
