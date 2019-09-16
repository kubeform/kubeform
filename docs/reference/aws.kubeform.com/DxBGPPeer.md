---
title: DxBGPPeer
menu:
  docs_{{ .version }}:
    identifier: dxbgppeer-aws.kubeform.com
    name: DxBGPPeer
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DxBGPPeer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DxBGPPeer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DxBGPPeerSpec](#dxbgppeerspec)***||
| `status` | ***[DxBGPPeerStatus](#dxbgppeerstatus)***||
## DxBGPPeerSpec

Appears on:[DxBGPPeer](#dxbgppeer), [DxBGPPeerStatus](#dxbgppeerstatus)

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

Appears on:[DxBGPPeer](#dxbgppeer)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DxBGPPeerSpec](#dxbgppeerspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
