---
title: ApiManagementProperty
menu:
  docs_{{ .version }}:
    identifier: apimanagementproperty-azurerm.kubeform.com
    name: ApiManagementProperty
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiManagementProperty
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementProperty` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementPropertySpec](#apimanagementpropertyspec)***||
| `status` | ***[ApiManagementPropertyStatus](#apimanagementpropertystatus)***||
## ApiManagementPropertySpec

Appears on:[ApiManagementProperty](#apimanagementproperty), [ApiManagementPropertyStatus](#apimanagementpropertystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiManagementName` | ***string***||
| `displayName` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `secret` | ***bool***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
| `value` | ***string***||
## ApiManagementPropertyStatus

Appears on:[ApiManagementProperty](#apimanagementproperty)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementPropertySpec](#apimanagementpropertyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ApiManagementPropertyStatus](#apimanagementpropertystatus)

---
