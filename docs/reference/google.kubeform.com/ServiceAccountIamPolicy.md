---
title: ServiceAccountIamPolicy
menu:
  docs_{{ .version }}:
    identifier: serviceaccountiampolicy-google.kubeform.com
    name: ServiceAccountIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ServiceAccountIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ServiceAccountIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServiceAccountIamPolicySpec](#serviceaccountiampolicyspec)***||
| `status` | ***[ServiceAccountIamPolicyStatus](#serviceaccountiampolicystatus)***||
## Phase(`string` alias)

Appears on:[ServiceAccountIamPolicyStatus](#serviceaccountiampolicystatus)

## ServiceAccountIamPolicySpec

Appears on:[ServiceAccountIamPolicy](#serviceaccountiampolicy), [ServiceAccountIamPolicyStatus](#serviceaccountiampolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `policyData` | ***string***||
| `serviceAccountID` | ***string***||
## ServiceAccountIamPolicyStatus

Appears on:[ServiceAccountIamPolicy](#serviceaccountiampolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServiceAccountIamPolicySpec](#serviceaccountiampolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
