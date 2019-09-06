## VpcPeeringConnection
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `VpcPeeringConnection` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpcPeeringConnectionSpec](#VpcPeeringConnectionSpec)***||
| `status` | ***[VpcPeeringConnectionStatus](#VpcPeeringConnectionStatus)***||
## VpcPeeringConnectionSpec
##### (Appears on:[VpcPeeringConnection](#VpcPeeringConnection), [VpcPeeringConnectionStatus](#VpcPeeringConnectionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `acceptStatus` | ***string***| ***(Optional)*** |
| `accepter` | ***[[]VpcPeeringConnectionSpecAccepter](#VpcPeeringConnectionSpecAccepter)***| ***(Optional)*** |
| `autoAccept` | ***bool***| ***(Optional)*** |
| `peerOwnerID` | ***string***| ***(Optional)*** |
| `peerRegion` | ***string***| ***(Optional)*** |
| `peerVpcID` | ***string***||
| `requester` | ***[[]VpcPeeringConnectionSpecRequester](#VpcPeeringConnectionSpecRequester)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***||
## VpcPeeringConnectionSpecAccepter
##### (Appears on:[VpcPeeringConnectionSpec](#VpcPeeringConnectionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowClassicLinkToRemoteVpc` | ***bool***| ***(Optional)*** |
| `allowRemoteVpcDNSResolution` | ***bool***| ***(Optional)*** |
| `allowVpcToRemoteClassicLink` | ***bool***| ***(Optional)*** |
## VpcPeeringConnectionSpecRequester
##### (Appears on:[VpcPeeringConnectionSpec](#VpcPeeringConnectionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowClassicLinkToRemoteVpc` | ***bool***| ***(Optional)*** |
| `allowRemoteVpcDNSResolution` | ***bool***| ***(Optional)*** |
| `allowVpcToRemoteClassicLink` | ***bool***| ***(Optional)*** |
## VpcPeeringConnectionStatus
##### (Appears on:[VpcPeeringConnection](#VpcPeeringConnection))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpcPeeringConnectionSpec](#VpcPeeringConnectionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
