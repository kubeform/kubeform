---
title: IapWebBackendServiceIamMember
menu:
  docs_{{ .version }}:
    identifier: iapwebbackendserviceiammember-google.kubeform.com
    name: IapWebBackendServiceIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IapWebBackendServiceIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `IapWebBackendServiceIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IapWebBackendServiceIamMemberSpec](#iapwebbackendserviceiammemberspec)***||
| `status` | ***[IapWebBackendServiceIamMemberStatus](#iapwebbackendserviceiammemberstatus)***||
## IapWebBackendServiceIamMemberSpec

Appears on:[IapWebBackendServiceIamMember](#iapwebbackendserviceiammember), [IapWebBackendServiceIamMemberStatus](#iapwebbackendserviceiammemberstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `member` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
| `webBackendService` | ***string***||
## IapWebBackendServiceIamMemberStatus

Appears on:[IapWebBackendServiceIamMember](#iapwebbackendserviceiammember)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IapWebBackendServiceIamMemberSpec](#iapwebbackendserviceiammemberspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IapWebBackendServiceIamMemberStatus](#iapwebbackendserviceiammemberstatus)

---
