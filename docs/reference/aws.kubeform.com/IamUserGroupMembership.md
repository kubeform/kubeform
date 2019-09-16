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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamUserGroupMembershipSpec](#IamUserGroupMembershipSpec)***||
| `status` | ***[IamUserGroupMembershipStatus](#IamUserGroupMembershipStatus)***||
## IamUserGroupMembershipSpec

Appears on:[IamUserGroupMembership](#IamUserGroupMembership), [IamUserGroupMembershipStatus](#IamUserGroupMembershipStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `groups` | ***[]string***||
| `user` | ***string***||
## IamUserGroupMembershipStatus

Appears on:[IamUserGroupMembership](#IamUserGroupMembership)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamUserGroupMembershipSpec](#IamUserGroupMembershipSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
