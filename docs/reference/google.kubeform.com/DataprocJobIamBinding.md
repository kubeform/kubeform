---
title: DataprocJobIamBinding
menu:
  docs_{{ .version }}:
    identifier: dataprocjobiambinding-google.kubeform.com
    name: DataprocJobIamBinding
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DataprocJobIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `DataprocJobIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataprocJobIamBindingSpec](#dataprocjobiambindingspec)***||
| `status` | ***[DataprocJobIamBindingStatus](#dataprocjobiambindingstatus)***||
## DataprocJobIamBindingSpec

Appears on:[DataprocJobIamBinding](#dataprocjobiambinding), [DataprocJobIamBindingStatus](#dataprocjobiambindingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `jobID` | ***string***||
| `members` | ***[]string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `role` | ***string***||
## DataprocJobIamBindingStatus

Appears on:[DataprocJobIamBinding](#dataprocjobiambinding)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataprocJobIamBindingSpec](#dataprocjobiambindingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DataprocJobIamBindingStatus](#dataprocjobiambindingstatus)

---
