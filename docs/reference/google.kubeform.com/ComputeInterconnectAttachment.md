---
title: ComputeInterconnectAttachment
menu:
  docs_v0.0.1:
    identifier: computeinterconnectattachment-google.kubeform.com
    name: ComputeInterconnectAttachment
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ComputeInterconnectAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeInterconnectAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeInterconnectAttachmentSpec](#ComputeInterconnectAttachmentSpec)***||
| `status` | ***[ComputeInterconnectAttachmentStatus](#ComputeInterconnectAttachmentStatus)***||
## ComputeInterconnectAttachmentSpec
##### (Appears on:[ComputeInterconnectAttachment](#ComputeInterconnectAttachment), [ComputeInterconnectAttachmentStatus](#ComputeInterconnectAttachmentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cloudRouterIPAddress` | ***string***| ***(Optional)*** |
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `customerRouterIPAddress` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `googleReferenceID` | ***string***| ***(Optional)*** |
| `interconnect` | ***string***||
| `name` | ***string***||
| `privateInterconnectInfo` | ***[[]ComputeInterconnectAttachmentSpecPrivateInterconnectInfo](#ComputeInterconnectAttachmentSpecPrivateInterconnectInfo)***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `router` | ***string***||
| `selfLink` | ***string***| ***(Optional)*** |
## ComputeInterconnectAttachmentSpecPrivateInterconnectInfo
##### (Appears on:[ComputeInterconnectAttachmentSpec](#ComputeInterconnectAttachmentSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `tag8021q` | ***int***| ***(Optional)*** |
## ComputeInterconnectAttachmentStatus
##### (Appears on:[ComputeInterconnectAttachment](#ComputeInterconnectAttachment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeInterconnectAttachmentSpec](#ComputeInterconnectAttachmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
