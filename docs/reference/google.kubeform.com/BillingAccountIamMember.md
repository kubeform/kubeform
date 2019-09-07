---
title: BillingAccountIamMember
menu:
  docs_v0.0.1:
    identifier: billingaccountiammember-google.kubeform.com
    name: BillingAccountIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## BillingAccountIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `BillingAccountIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BillingAccountIamMemberSpec](#BillingAccountIamMemberSpec)***||
| `status` | ***[BillingAccountIamMemberStatus](#BillingAccountIamMemberStatus)***||
## BillingAccountIamMemberSpec
##### (Appears on:[BillingAccountIamMember](#BillingAccountIamMember), [BillingAccountIamMemberStatus](#BillingAccountIamMemberStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `billingAccountID` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `member` | ***string***||
| `role` | ***string***||
## BillingAccountIamMemberStatus
##### (Appears on:[BillingAccountIamMember](#BillingAccountIamMember))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BillingAccountIamMemberSpec](#BillingAccountIamMemberSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
