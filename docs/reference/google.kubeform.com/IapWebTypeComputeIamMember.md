---
title: IapWebTypeComputeIamMember
menu:
  docs_{{ .version }}:
    identifier: iapwebtypecomputeiammember-google.kubeform.com
    name: IapWebTypeComputeIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IapWebTypeComputeIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `IapWebTypeComputeIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IapWebTypeComputeIamMemberSpec](#iapwebtypecomputeiammemberspec)***||
| `status` | ***[IapWebTypeComputeIamMemberStatus](#iapwebtypecomputeiammemberstatus)***||
## IapWebTypeComputeIamMemberSpec

Appears on:[IapWebTypeComputeIamMember](#iapwebtypecomputeiammember), [IapWebTypeComputeIamMemberStatus](#iapwebtypecomputeiammemberstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `member` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
## IapWebTypeComputeIamMemberStatus

Appears on:[IapWebTypeComputeIamMember](#iapwebtypecomputeiammember)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IapWebTypeComputeIamMemberSpec](#iapwebtypecomputeiammemberspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IapWebTypeComputeIamMemberStatus](#iapwebtypecomputeiammemberstatus)

---
