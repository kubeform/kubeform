---
title: ApiManagementGroupUser
menu:
  docs_{{ .version }}:
    identifier: apimanagementgroupuser-azurerm.kubeform.com
    name: ApiManagementGroupUser
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiManagementGroupUser
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementGroupUser` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementGroupUserSpec](#apimanagementgroupuserspec)***||
| `status` | ***[ApiManagementGroupUserStatus](#apimanagementgroupuserstatus)***||
## ApiManagementGroupUserSpec

Appears on:[ApiManagementGroupUser](#apimanagementgroupuser), [ApiManagementGroupUserStatus](#apimanagementgroupuserstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiManagementName` | ***string***||
| `groupName` | ***string***||
| `resourceGroupName` | ***string***||
| `userID` | ***string***||
## ApiManagementGroupUserStatus

Appears on:[ApiManagementGroupUser](#apimanagementgroupuser)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementGroupUserSpec](#apimanagementgroupuserspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
