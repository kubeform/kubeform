---
title: DataprocJobIamMember
menu:
  docs_{{ .version }}:
    identifier: dataprocjobiammember-google.kubeform.com
    name: DataprocJobIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DataprocJobIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `DataprocJobIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataprocJobIamMemberSpec](#dataprocjobiammemberspec)***||
| `status` | ***[DataprocJobIamMemberStatus](#dataprocjobiammemberstatus)***||
## DataprocJobIamMemberSpec

Appears on:[DataprocJobIamMember](#dataprocjobiammember), [DataprocJobIamMemberStatus](#dataprocjobiammemberstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `jobID` | ***string***||
| `member` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `role` | ***string***||
## DataprocJobIamMemberStatus

Appears on:[DataprocJobIamMember](#dataprocjobiammember)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataprocJobIamMemberSpec](#dataprocjobiammemberspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DataprocJobIamMemberStatus](#dataprocjobiammemberstatus)

---
