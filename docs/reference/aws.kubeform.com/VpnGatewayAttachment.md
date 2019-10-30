---
title: VpnGatewayAttachment
menu:
  docs_{{ .version }}:
    identifier: vpngatewayattachment-aws.kubeform.com
    name: VpnGatewayAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VpnGatewayAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `VpnGatewayAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpnGatewayAttachmentSpec](#vpngatewayattachmentspec)***||
| `status` | ***[VpnGatewayAttachmentStatus](#vpngatewayattachmentstatus)***||
## Phase(`string` alias)

Appears on:[VpnGatewayAttachmentStatus](#vpngatewayattachmentstatus)

## VpnGatewayAttachmentSpec

Appears on:[VpnGatewayAttachment](#vpngatewayattachment), [VpnGatewayAttachmentStatus](#vpngatewayattachmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `vpcID` | ***string***||
| `vpnGatewayID` | ***string***||
## VpnGatewayAttachmentStatus

Appears on:[VpnGatewayAttachment](#vpngatewayattachment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpnGatewayAttachmentSpec](#vpngatewayattachmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
