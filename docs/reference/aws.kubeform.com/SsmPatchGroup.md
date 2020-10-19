---
title: SsmPatchGroup
menu:
  docs_{{ .version }}:
    identifier: ssmpatchgroup-aws.kubeform.com
    name: SsmPatchGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SsmPatchGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SsmPatchGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SsmPatchGroupSpec](#ssmpatchgroupspec)***||
| `status` | ***[SsmPatchGroupStatus](#ssmpatchgroupstatus)***||
## Phase(`string` alias)

Appears on:[SsmPatchGroupStatus](#ssmpatchgroupstatus)

## SsmPatchGroupSpec

Appears on:[SsmPatchGroup](#ssmpatchgroup), [SsmPatchGroupStatus](#ssmpatchgroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `baselineID` | ***string***||
| `patchGroup` | ***string***||
## SsmPatchGroupStatus

Appears on:[SsmPatchGroup](#ssmpatchgroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SsmPatchGroupSpec](#ssmpatchgroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
