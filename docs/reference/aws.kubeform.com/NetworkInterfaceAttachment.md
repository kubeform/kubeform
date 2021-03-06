---
title: NetworkInterfaceAttachment
menu:
  docs_{{ .version }}:
    identifier: networkinterfaceattachment-aws.kubeform.com
    name: NetworkInterfaceAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NetworkInterfaceAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkInterfaceAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkInterfaceAttachmentSpec](#networkinterfaceattachmentspec)***||
| `status` | ***[NetworkInterfaceAttachmentStatus](#networkinterfaceattachmentstatus)***||
## NetworkInterfaceAttachmentSpec

Appears on:[NetworkInterfaceAttachment](#networkinterfaceattachment), [NetworkInterfaceAttachmentStatus](#networkinterfaceattachmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `attachmentID` | ***string***| ***(Optional)*** |
| `deviceIndex` | ***int64***||
| `instanceID` | ***string***||
| `networkInterfaceID` | ***string***||
| `status` | ***string***| ***(Optional)*** |
## NetworkInterfaceAttachmentStatus

Appears on:[NetworkInterfaceAttachment](#networkinterfaceattachment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkInterfaceAttachmentSpec](#networkinterfaceattachmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[NetworkInterfaceAttachmentStatus](#networkinterfaceattachmentstatus)

---
