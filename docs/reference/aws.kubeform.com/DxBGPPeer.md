---
title: DxBGPPeer
menu:
  docs_v0.0.1:
    identifier: dxbgppeer-aws.kubeform.com
    name: DxBGPPeer
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DxBGPPeer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DxBGPPeer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DxBGPPeerSpec](#DxBGPPeerSpec)***||
| `status` | ***[DxBGPPeerStatus](#DxBGPPeerStatus)***||
## DxBGPPeerSpec
##### (Appears on:[DxBGPPeer](#DxBGPPeer), [DxBGPPeerStatus](#DxBGPPeerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `addressFamily` | ***string***||
| `amazonAddress` | ***string***| ***(Optional)*** |
| `awsDevice` | ***string***| ***(Optional)*** |
| `bgpAsn` | ***int***||
| `bgpAuthKey` | ***string***| ***(Optional)*** |
| `bgpPeerID` | ***string***| ***(Optional)*** |
| `bgpStatus` | ***string***| ***(Optional)*** |
| `customerAddress` | ***string***| ***(Optional)*** |
| `virtualInterfaceID` | ***string***||
## DxBGPPeerStatus
##### (Appears on:[DxBGPPeer](#DxBGPPeer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DxBGPPeerSpec](#DxBGPPeerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
