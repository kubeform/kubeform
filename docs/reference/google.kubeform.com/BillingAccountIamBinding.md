---
title: BillingAccountIamBinding
menu:
  docs_v0.0.1:
    identifier: billingaccountiambinding-google.kubeform.com
    name: BillingAccountIamBinding
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## BillingAccountIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `BillingAccountIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BillingAccountIamBindingSpec](#BillingAccountIamBindingSpec)***||
| `status` | ***[BillingAccountIamBindingStatus](#BillingAccountIamBindingStatus)***||
## BillingAccountIamBindingSpec
##### (Appears on:[BillingAccountIamBinding](#BillingAccountIamBinding), [BillingAccountIamBindingStatus](#BillingAccountIamBindingStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `billingAccountID` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `members` | ***[]string***||
| `role` | ***string***||
## BillingAccountIamBindingStatus
##### (Appears on:[BillingAccountIamBinding](#BillingAccountIamBinding))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BillingAccountIamBindingSpec](#BillingAccountIamBindingSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
