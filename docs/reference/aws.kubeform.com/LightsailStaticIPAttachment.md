---
title: LightsailStaticIPAttachment
menu:
  docs_{{ .version }}:
    identifier: lightsailstaticipattachment-aws.kubeform.com
    name: LightsailStaticIPAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LightsailStaticIPAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LightsailStaticIPAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LightsailStaticIPAttachmentSpec](#lightsailstaticipattachmentspec)***||
| `status` | ***[LightsailStaticIPAttachmentStatus](#lightsailstaticipattachmentstatus)***||
## LightsailStaticIPAttachmentSpec

Appears on:[LightsailStaticIPAttachment](#lightsailstaticipattachment), [LightsailStaticIPAttachmentStatus](#lightsailstaticipattachmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `instanceName` | ***string***||
| `ipAddress` | ***string***| ***(Optional)*** |
| `staticIPName` | ***string***||
## LightsailStaticIPAttachmentStatus

Appears on:[LightsailStaticIPAttachment](#lightsailstaticipattachment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LightsailStaticIPAttachmentSpec](#lightsailstaticipattachmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LightsailStaticIPAttachmentStatus](#lightsailstaticipattachmentstatus)

---
