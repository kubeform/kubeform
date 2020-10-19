---
title: IamGroupPolicyAttachment
menu:
  docs_{{ .version }}:
    identifier: iamgrouppolicyattachment-aws.kubeform.com
    name: IamGroupPolicyAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IamGroupPolicyAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamGroupPolicyAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamGroupPolicyAttachmentSpec](#iamgrouppolicyattachmentspec)***||
| `status` | ***[IamGroupPolicyAttachmentStatus](#iamgrouppolicyattachmentstatus)***||
## IamGroupPolicyAttachmentSpec

Appears on:[IamGroupPolicyAttachment](#iamgrouppolicyattachment), [IamGroupPolicyAttachmentStatus](#iamgrouppolicyattachmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `group` | ***string***||
| `policyArn` | ***string***||
## IamGroupPolicyAttachmentStatus

Appears on:[IamGroupPolicyAttachment](#iamgrouppolicyattachment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamGroupPolicyAttachmentSpec](#iamgrouppolicyattachmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IamGroupPolicyAttachmentStatus](#iamgrouppolicyattachmentstatus)

---
