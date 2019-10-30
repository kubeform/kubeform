---
title: VpcPeeringConnection
menu:
  docs_{{ .version }}:
    identifier: vpcpeeringconnection-aws.kubeform.com
    name: VpcPeeringConnection
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VpcPeeringConnection
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `VpcPeeringConnection` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpcPeeringConnectionSpec](#vpcpeeringconnectionspec)***||
| `status` | ***[VpcPeeringConnectionStatus](#vpcpeeringconnectionstatus)***||
## Phase(`string` alias)

Appears on:[VpcPeeringConnectionStatus](#vpcpeeringconnectionstatus)

## VpcPeeringConnectionSpec

Appears on:[VpcPeeringConnection](#vpcpeeringconnection), [VpcPeeringConnectionStatus](#vpcpeeringconnectionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `acceptStatus` | ***string***| ***(Optional)*** |
| `accepter` | ***[[]VpcPeeringConnectionSpecAccepter](#vpcpeeringconnectionspecaccepter)***| ***(Optional)*** |
| `autoAccept` | ***bool***| ***(Optional)*** |
| `peerOwnerID` | ***string***| ***(Optional)*** |
| `peerRegion` | ***string***| ***(Optional)*** |
| `peerVpcID` | ***string***||
| `requester` | ***[[]VpcPeeringConnectionSpecRequester](#vpcpeeringconnectionspecrequester)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***||
## VpcPeeringConnectionSpecAccepter

Appears on:[VpcPeeringConnectionSpec](#vpcpeeringconnectionspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowClassicLinkToRemoteVpc` | ***bool***| ***(Optional)*** |
| `allowRemoteVpcDNSResolution` | ***bool***| ***(Optional)*** |
| `allowVpcToRemoteClassicLink` | ***bool***| ***(Optional)*** |
## VpcPeeringConnectionSpecRequester

Appears on:[VpcPeeringConnectionSpec](#vpcpeeringconnectionspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowClassicLinkToRemoteVpc` | ***bool***| ***(Optional)*** |
| `allowRemoteVpcDNSResolution` | ***bool***| ***(Optional)*** |
| `allowVpcToRemoteClassicLink` | ***bool***| ***(Optional)*** |
## VpcPeeringConnectionStatus

Appears on:[VpcPeeringConnection](#vpcpeeringconnection)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpcPeeringConnectionSpec](#vpcpeeringconnectionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
