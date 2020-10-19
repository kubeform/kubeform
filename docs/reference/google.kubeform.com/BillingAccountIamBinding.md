---
title: BillingAccountIamBinding
menu:
  docs_{{ .version }}:
    identifier: billingaccountiambinding-google.kubeform.com
    name: BillingAccountIamBinding
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BillingAccountIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `BillingAccountIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BillingAccountIamBindingSpec](#billingaccountiambindingspec)***||
| `status` | ***[BillingAccountIamBindingStatus](#billingaccountiambindingstatus)***||
## BillingAccountIamBindingSpec

Appears on:[BillingAccountIamBinding](#billingaccountiambinding), [BillingAccountIamBindingStatus](#billingaccountiambindingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `billingAccountID` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `members` | ***[]string***||
| `role` | ***string***||
## BillingAccountIamBindingStatus

Appears on:[BillingAccountIamBinding](#billingaccountiambinding)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BillingAccountIamBindingSpec](#billingaccountiambindingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BillingAccountIamBindingStatus](#billingaccountiambindingstatus)

---
