---
title: NetworkInterfaceSgAttachment
menu:
  docs_{{ .version }}:
    identifier: networkinterfacesgattachment-aws.kubeform.com
    name: NetworkInterfaceSgAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NetworkInterfaceSgAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkInterfaceSgAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkInterfaceSgAttachmentSpec](#networkinterfacesgattachmentspec)***||
| `status` | ***[NetworkInterfaceSgAttachmentStatus](#networkinterfacesgattachmentstatus)***||
## NetworkInterfaceSgAttachmentSpec

Appears on:[NetworkInterfaceSgAttachment](#networkinterfacesgattachment), [NetworkInterfaceSgAttachmentStatus](#networkinterfacesgattachmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `networkInterfaceID` | ***string***||
| `securityGroupID` | ***string***||
## NetworkInterfaceSgAttachmentStatus

Appears on:[NetworkInterfaceSgAttachment](#networkinterfacesgattachment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkInterfaceSgAttachmentSpec](#networkinterfacesgattachmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
