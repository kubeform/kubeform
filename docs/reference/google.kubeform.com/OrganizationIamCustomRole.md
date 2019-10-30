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
| `spec` | ***[OrganizationIamCustomRoleSpec](#organizationiamcustomrolespec)***||
| `status` | ***[OrganizationIamCustomRoleStatus](#organizationiamcustomrolestatus)***||
## OrganizationIamCustomRoleSpec

Appears on:[OrganizationIamCustomRole](#organizationiamcustomrole), [OrganizationIamCustomRoleStatus](#organizationiamcustomrolestatus)

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

Appears on:[OrganizationIamCustomRole](#organizationiamcustomrole)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OrganizationIamCustomRoleSpec](#organizationiamcustomrolespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[OrganizationIamCustomRoleStatus](#organizationiamcustomrolestatus)

---
