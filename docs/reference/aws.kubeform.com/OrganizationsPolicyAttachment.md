---
title: OrganizationsPolicyAttachment
menu:
  docs_{{ .version }}:
    identifier: organizationspolicyattachment-aws.kubeform.com
    name: OrganizationsPolicyAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## OrganizationsPolicyAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OrganizationsPolicyAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OrganizationsPolicyAttachmentSpec](#organizationspolicyattachmentspec)***||
| `status` | ***[OrganizationsPolicyAttachmentStatus](#organizationspolicyattachmentstatus)***||
## OrganizationsPolicyAttachmentSpec

Appears on:[OrganizationsPolicyAttachment](#organizationspolicyattachment), [OrganizationsPolicyAttachmentStatus](#organizationspolicyattachmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `policyID` | ***string***||
| `targetID` | ***string***||
## OrganizationsPolicyAttachmentStatus

Appears on:[OrganizationsPolicyAttachment](#organizationspolicyattachment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OrganizationsPolicyAttachmentSpec](#organizationspolicyattachmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[OrganizationsPolicyAttachmentStatus](#organizationspolicyattachmentstatus)

---
