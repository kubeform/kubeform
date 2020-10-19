---
title: ComputeInstanceIamBinding
menu:
  docs_{{ .version }}:
    identifier: computeinstanceiambinding-google.kubeform.com
    name: ComputeInstanceIamBinding
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeInstanceIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeInstanceIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeInstanceIamBindingSpec](#computeinstanceiambindingspec)***||
| `status` | ***[ComputeInstanceIamBindingStatus](#computeinstanceiambindingstatus)***||
## ComputeInstanceIamBindingSpec

Appears on:[ComputeInstanceIamBinding](#computeinstanceiambinding), [ComputeInstanceIamBindingStatus](#computeinstanceiambindingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `instanceName` | ***string***||
| `members` | ***[]string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
| `zone` | ***string***| ***(Optional)*** |
## ComputeInstanceIamBindingStatus

Appears on:[ComputeInstanceIamBinding](#computeinstanceiambinding)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeInstanceIamBindingSpec](#computeinstanceiambindingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeInstanceIamBindingStatus](#computeinstanceiambindingstatus)

---
