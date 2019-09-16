---
title: OrganizationIamCustomRole
menu:
  docs_{{ .version }}:
    identifier: organizationiamcustomrole-google.kubeform.com
    name: OrganizationIamCustomRole
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## OrganizationIamCustomRole
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `OrganizationIamCustomRole` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OrganizationIamCustomRoleSpec](#OrganizationIamCustomRoleSpec)***||
| `status` | ***[OrganizationIamCustomRoleStatus](#OrganizationIamCustomRoleStatus)***||
## OrganizationIamCustomRoleSpec

Appears on:[OrganizationIamCustomRole](#OrganizationIamCustomRole), [OrganizationIamCustomRoleStatus](#OrganizationIamCustomRoleStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `deleted` | ***bool***| ***(Optional)*** Deprecated|
| `description` | ***string***| ***(Optional)*** |
| `orgID` | ***string***||
| `permissions` | ***[]string***||
| `roleID` | ***string***||
| `stage` | ***string***| ***(Optional)*** |
| `title` | ***string***||
## OrganizationIamCustomRoleStatus

Appears on:[OrganizationIamCustomRole](#OrganizationIamCustomRole)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OrganizationIamCustomRoleSpec](#OrganizationIamCustomRoleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
