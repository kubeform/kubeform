---
title: IapWebBackendServiceIamPolicy
menu:
  docs_{{ .version }}:
    identifier: iapwebbackendserviceiampolicy-google.kubeform.com
    name: IapWebBackendServiceIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IapWebBackendServiceIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `IapWebBackendServiceIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IapWebBackendServiceIamPolicySpec](#iapwebbackendserviceiampolicyspec)***||
| `status` | ***[IapWebBackendServiceIamPolicyStatus](#iapwebbackendserviceiampolicystatus)***||
## IapWebBackendServiceIamPolicySpec

Appears on:[IapWebBackendServiceIamPolicy](#iapwebbackendserviceiampolicy), [IapWebBackendServiceIamPolicyStatus](#iapwebbackendserviceiampolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `policyData` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `webBackendService` | ***string***||
## IapWebBackendServiceIamPolicyStatus

Appears on:[IapWebBackendServiceIamPolicy](#iapwebbackendserviceiampolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IapWebBackendServiceIamPolicySpec](#iapwebbackendserviceiampolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IapWebBackendServiceIamPolicyStatus](#iapwebbackendserviceiampolicystatus)

---
