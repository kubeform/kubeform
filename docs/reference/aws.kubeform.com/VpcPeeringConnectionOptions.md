---
title: VpcPeeringConnectionOptions
menu:
  docs_{{ .version }}:
    identifier: vpcpeeringconnectionoptions-aws.kubeform.com
    name: VpcPeeringConnectionOptions
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VpcPeeringConnectionOptions
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `VpcPeeringConnectionOptions` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpcPeeringConnectionOptionsSpec](#VpcPeeringConnectionOptionsSpec)***||
| `status` | ***[VpcPeeringConnectionOptionsStatus](#VpcPeeringConnectionOptionsStatus)***||
## VpcPeeringConnectionOptionsSpec

Appears on:[VpcPeeringConnectionOptions](#VpcPeeringConnectionOptions), [VpcPeeringConnectionOptionsStatus](#VpcPeeringConnectionOptionsStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accepter` | ***[[]VpcPeeringConnectionOptionsSpecAccepter](#VpcPeeringConnectionOptionsSpecAccepter)***| ***(Optional)*** |
| `requester` | ***[[]VpcPeeringConnectionOptionsSpecRequester](#VpcPeeringConnectionOptionsSpecRequester)***| ***(Optional)*** |
| `vpcPeeringConnectionID` | ***string***||
## VpcPeeringConnectionOptionsSpecAccepter

Appears on:[VpcPeeringConnectionOptionsSpec](#VpcPeeringConnectionOptionsSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowClassicLinkToRemoteVpc` | ***bool***| ***(Optional)*** |
| `allowRemoteVpcDNSResolution` | ***bool***| ***(Optional)*** |
| `allowVpcToRemoteClassicLink` | ***bool***| ***(Optional)*** |
## VpcPeeringConnectionOptionsSpecRequester

Appears on:[VpcPeeringConnectionOptionsSpec](#VpcPeeringConnectionOptionsSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowClassicLinkToRemoteVpc` | ***bool***| ***(Optional)*** |
| `allowRemoteVpcDNSResolution` | ***bool***| ***(Optional)*** |
| `allowVpcToRemoteClassicLink` | ***bool***| ***(Optional)*** |
## VpcPeeringConnectionOptionsStatus

Appears on:[VpcPeeringConnectionOptions](#VpcPeeringConnectionOptions)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpcPeeringConnectionOptionsSpec](#VpcPeeringConnectionOptionsSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
