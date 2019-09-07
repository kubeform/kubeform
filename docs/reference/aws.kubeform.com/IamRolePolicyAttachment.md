---
title: IamRolePolicyAttachment
menu:
  docs_v0.0.1:
    identifier: iamrolepolicyattachment-aws.kubeform.com
    name: IamRolePolicyAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## IamRolePolicyAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamRolePolicyAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamRolePolicyAttachmentSpec](#IamRolePolicyAttachmentSpec)***||
| `status` | ***[IamRolePolicyAttachmentStatus](#IamRolePolicyAttachmentStatus)***||
## IamRolePolicyAttachmentSpec
##### (Appears on:[IamRolePolicyAttachment](#IamRolePolicyAttachment), [IamRolePolicyAttachmentStatus](#IamRolePolicyAttachmentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `policyArn` | ***string***||
| `role` | ***string***||
## IamRolePolicyAttachmentStatus
##### (Appears on:[IamRolePolicyAttachment](#IamRolePolicyAttachment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamRolePolicyAttachmentSpec](#IamRolePolicyAttachmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
