---
title: ComputeSubnetworkIamMember
menu:
  docs_{{ .version }}:
    identifier: computesubnetworkiammember-google.kubeform.com
    name: ComputeSubnetworkIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeSubnetworkIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeSubnetworkIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeSubnetworkIamMemberSpec](#computesubnetworkiammemberspec)***||
| `status` | ***[ComputeSubnetworkIamMemberStatus](#computesubnetworkiammemberstatus)***||
## ComputeSubnetworkIamMemberSpec

Appears on:[ComputeSubnetworkIamMember](#computesubnetworkiammember), [ComputeSubnetworkIamMemberStatus](#computesubnetworkiammemberstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `member` | ***string***||
| `project` | ***string***| ***(Optional)*** Deprecated|
| `region` | ***string***| ***(Optional)*** Deprecated|
| `role` | ***string***||
| `subnetwork` | ***string***|Deprecated|
## ComputeSubnetworkIamMemberStatus

Appears on:[ComputeSubnetworkIamMember](#computesubnetworkiammember)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeSubnetworkIamMemberSpec](#computesubnetworkiammemberspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeSubnetworkIamMemberStatus](#computesubnetworkiammemberstatus)

---
