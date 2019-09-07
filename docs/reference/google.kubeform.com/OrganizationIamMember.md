---
title: OrganizationIamMember
menu:
  docs_v0.0.1:
    identifier: organizationiammember-google.kubeform.com
    name: OrganizationIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## OrganizationIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `OrganizationIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OrganizationIamMemberSpec](#OrganizationIamMemberSpec)***||
| `status` | ***[OrganizationIamMemberStatus](#OrganizationIamMemberStatus)***||
## OrganizationIamMemberSpec
##### (Appears on:[OrganizationIamMember](#OrganizationIamMember), [OrganizationIamMemberStatus](#OrganizationIamMemberStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `member` | ***string***||
| `orgID` | ***string***||
| `role` | ***string***||
## OrganizationIamMemberStatus
##### (Appears on:[OrganizationIamMember](#OrganizationIamMember))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OrganizationIamMemberSpec](#OrganizationIamMemberSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
