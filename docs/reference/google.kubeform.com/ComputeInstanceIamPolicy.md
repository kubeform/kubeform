---
title: ComputeInstanceIamPolicy
menu:
  docs_{{ .version }}:
    identifier: computeinstanceiampolicy-google.kubeform.com
    name: ComputeInstanceIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeInstanceIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeInstanceIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeInstanceIamPolicySpec](#computeinstanceiampolicyspec)***||
| `status` | ***[ComputeInstanceIamPolicyStatus](#computeinstanceiampolicystatus)***||
## ComputeInstanceIamPolicySpec

Appears on:[ComputeInstanceIamPolicy](#computeinstanceiampolicy), [ComputeInstanceIamPolicyStatus](#computeinstanceiampolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `instanceName` | ***string***||
| `policyData` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## ComputeInstanceIamPolicyStatus

Appears on:[ComputeInstanceIamPolicy](#computeinstanceiampolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeInstanceIamPolicySpec](#computeinstanceiampolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeInstanceIamPolicyStatus](#computeinstanceiampolicystatus)

---
