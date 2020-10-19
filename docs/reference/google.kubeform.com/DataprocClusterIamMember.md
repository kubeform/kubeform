---
title: DataprocClusterIamMember
menu:
  docs_{{ .version }}:
    identifier: dataprocclusteriammember-google.kubeform.com
    name: DataprocClusterIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DataprocClusterIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `DataprocClusterIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataprocClusterIamMemberSpec](#dataprocclusteriammemberspec)***||
| `status` | ***[DataprocClusterIamMemberStatus](#dataprocclusteriammemberstatus)***||
## DataprocClusterIamMemberSpec

Appears on:[DataprocClusterIamMember](#dataprocclusteriammember), [DataprocClusterIamMemberStatus](#dataprocclusteriammemberstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cluster` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `member` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `role` | ***string***||
## DataprocClusterIamMemberStatus

Appears on:[DataprocClusterIamMember](#dataprocclusteriammember)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataprocClusterIamMemberSpec](#dataprocclusteriammemberspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DataprocClusterIamMemberStatus](#dataprocclusteriammemberstatus)

---
