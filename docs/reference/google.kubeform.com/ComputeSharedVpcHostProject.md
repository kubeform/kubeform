---
title: ComputeSharedVpcHostProject
menu:
  docs_{{ .version }}:
    identifier: computesharedvpchostproject-google.kubeform.com
    name: ComputeSharedVpcHostProject
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeSharedVpcHostProject
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeSharedVpcHostProject` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeSharedVpcHostProjectSpec](#computesharedvpchostprojectspec)***||
| `status` | ***[ComputeSharedVpcHostProjectStatus](#computesharedvpchostprojectstatus)***||
## ComputeSharedVpcHostProjectSpec

Appears on:[ComputeSharedVpcHostProject](#computesharedvpchostproject), [ComputeSharedVpcHostProjectStatus](#computesharedvpchostprojectstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `project` | ***string***||
## ComputeSharedVpcHostProjectStatus

Appears on:[ComputeSharedVpcHostProject](#computesharedvpchostproject)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeSharedVpcHostProjectSpec](#computesharedvpchostprojectspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeSharedVpcHostProjectStatus](#computesharedvpchostprojectstatus)

---
