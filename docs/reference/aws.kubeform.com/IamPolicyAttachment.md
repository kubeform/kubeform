---
title: IamPolicyAttachment
menu:
  docs_{{ .version }}:
    identifier: iampolicyattachment-aws.kubeform.com
    name: IamPolicyAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IamPolicyAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamPolicyAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamPolicyAttachmentSpec](#IamPolicyAttachmentSpec)***||
| `status` | ***[IamPolicyAttachmentStatus](#IamPolicyAttachmentStatus)***||
## IamPolicyAttachmentSpec
##### (Appears on:[IamPolicyAttachment](#IamPolicyAttachment), [IamPolicyAttachmentStatus](#IamPolicyAttachmentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `groups` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `policyArn` | ***string***||
| `roles` | ***[]string***| ***(Optional)*** |
| `users` | ***[]string***| ***(Optional)*** |
## IamPolicyAttachmentStatus
##### (Appears on:[IamPolicyAttachment](#IamPolicyAttachment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamPolicyAttachmentSpec](#IamPolicyAttachmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
