---
title: OrganizationIamPolicy
menu:
  docs_{{ .version }}:
    identifier: organizationiampolicy-google.kubeform.com
    name: OrganizationIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## OrganizationIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `OrganizationIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OrganizationIamPolicySpec](#organizationiampolicyspec)***||
| `status` | ***[OrganizationIamPolicyStatus](#organizationiampolicystatus)***||
## OrganizationIamPolicySpec

Appears on:[OrganizationIamPolicy](#organizationiampolicy), [OrganizationIamPolicyStatus](#organizationiampolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `orgID` | ***string***||
| `policyData` | ***string***||
## OrganizationIamPolicyStatus

Appears on:[OrganizationIamPolicy](#organizationiampolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OrganizationIamPolicySpec](#organizationiampolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[OrganizationIamPolicyStatus](#organizationiampolicystatus)

---
