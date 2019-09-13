---
title: ProjectIamCustomRole
menu:
  docs_{{ .version }}:
    identifier: projectiamcustomrole-google.kubeform.com
    name: ProjectIamCustomRole
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ProjectIamCustomRole
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ProjectIamCustomRole` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ProjectIamCustomRoleSpec](#ProjectIamCustomRoleSpec)***||
| `status` | ***[ProjectIamCustomRoleStatus](#ProjectIamCustomRoleStatus)***||
## ProjectIamCustomRoleSpec
##### (Appears on:[ProjectIamCustomRole](#ProjectIamCustomRole), [ProjectIamCustomRoleStatus](#ProjectIamCustomRoleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `deleted` | ***bool***| ***(Optional)*** Deprecated|
| `description` | ***string***| ***(Optional)*** |
| `permissions` | ***[]string***||
| `project` | ***string***| ***(Optional)*** |
| `roleID` | ***string***||
| `stage` | ***string***| ***(Optional)*** |
| `title` | ***string***||
## ProjectIamCustomRoleStatus
##### (Appears on:[ProjectIamCustomRole](#ProjectIamCustomRole))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ProjectIamCustomRoleSpec](#ProjectIamCustomRoleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
