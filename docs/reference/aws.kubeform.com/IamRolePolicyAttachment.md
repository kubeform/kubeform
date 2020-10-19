---
title: IamRolePolicyAttachment
menu:
  docs_{{ .version }}:
    identifier: iamrolepolicyattachment-aws.kubeform.com
    name: IamRolePolicyAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IamRolePolicyAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamRolePolicyAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamRolePolicyAttachmentSpec](#iamrolepolicyattachmentspec)***||
| `status` | ***[IamRolePolicyAttachmentStatus](#iamrolepolicyattachmentstatus)***||
## IamRolePolicyAttachmentSpec

Appears on:[IamRolePolicyAttachment](#iamrolepolicyattachment), [IamRolePolicyAttachmentStatus](#iamrolepolicyattachmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `policyArn` | ***string***||
| `role` | ***string***||
## IamRolePolicyAttachmentStatus

Appears on:[IamRolePolicyAttachment](#iamrolepolicyattachment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamRolePolicyAttachmentSpec](#iamrolepolicyattachmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IamRolePolicyAttachmentStatus](#iamrolepolicyattachmentstatus)

---
