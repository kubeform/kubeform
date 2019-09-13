---
title: BillingAccountIamPolicy
menu:
  docs_{{ .version }}:
    identifier: billingaccountiampolicy-google.kubeform.com
    name: BillingAccountIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BillingAccountIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `BillingAccountIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BillingAccountIamPolicySpec](#BillingAccountIamPolicySpec)***||
| `status` | ***[BillingAccountIamPolicyStatus](#BillingAccountIamPolicyStatus)***||
## BillingAccountIamPolicySpec
##### (Appears on:[BillingAccountIamPolicy](#BillingAccountIamPolicy), [BillingAccountIamPolicyStatus](#BillingAccountIamPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `billingAccountID` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `policyData` | ***string***||
## BillingAccountIamPolicyStatus
##### (Appears on:[BillingAccountIamPolicy](#BillingAccountIamPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BillingAccountIamPolicySpec](#BillingAccountIamPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
