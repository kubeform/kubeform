---
title: IamUserGroupMembership
menu:
  docs_{{ .version }}:
    identifier: iamusergroupmembership-aws.kubeform.com
    name: IamUserGroupMembership
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IamUserGroupMembership
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamUserGroupMembership` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamUserGroupMembershipSpec](#iamusergroupmembershipspec)***||
| `status` | ***[IamUserGroupMembershipStatus](#iamusergroupmembershipstatus)***||
## IamUserGroupMembershipSpec

Appears on:[IamUserGroupMembership](#iamusergroupmembership), [IamUserGroupMembershipStatus](#iamusergroupmembershipstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `groups` | ***[]string***||
| `user` | ***string***||
## IamUserGroupMembershipStatus

Appears on:[IamUserGroupMembership](#iamusergroupmembership)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamUserGroupMembershipSpec](#iamusergroupmembershipspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IamUserGroupMembershipStatus](#iamusergroupmembershipstatus)

---
