---
title: NetworkInterfaceSgAttachment
menu:
  docs_v0.0.1:
    identifier: networkinterfacesgattachment-aws.kubeform.com
    name: NetworkInterfaceSgAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## NetworkInterfaceSgAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkInterfaceSgAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkInterfaceSgAttachmentSpec](#NetworkInterfaceSgAttachmentSpec)***||
| `status` | ***[NetworkInterfaceSgAttachmentStatus](#NetworkInterfaceSgAttachmentStatus)***||
## NetworkInterfaceSgAttachmentSpec
##### (Appears on:[NetworkInterfaceSgAttachment](#NetworkInterfaceSgAttachment), [NetworkInterfaceSgAttachmentStatus](#NetworkInterfaceSgAttachmentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `networkInterfaceID` | ***string***||
| `securityGroupID` | ***string***||
## NetworkInterfaceSgAttachmentStatus
##### (Appears on:[NetworkInterfaceSgAttachment](#NetworkInterfaceSgAttachment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkInterfaceSgAttachmentSpec](#NetworkInterfaceSgAttachmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
