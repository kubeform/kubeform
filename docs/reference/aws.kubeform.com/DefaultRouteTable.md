## DefaultRouteTable
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DefaultRouteTable` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DefaultRouteTableSpec](#DefaultRouteTableSpec)***||
| `status` | ***[DefaultRouteTableStatus](#DefaultRouteTableStatus)***||
## DefaultRouteTableSpec
##### (Appears on:[DefaultRouteTable](#DefaultRouteTable), [DefaultRouteTableStatus](#DefaultRouteTableStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `defaultRouteTableID` | ***string***||
| `ownerID` | ***string***| ***(Optional)*** |
| `propagatingVgws` | ***[]string***| ***(Optional)*** |
| `route` | ***[[]DefaultRouteTableSpecRoute](#DefaultRouteTableSpecRoute)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
## DefaultRouteTableSpecRoute
##### (Appears on:[DefaultRouteTableSpec](#DefaultRouteTableSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cidrBlock` | ***string***| ***(Optional)*** |
| `egressOnlyGatewayID` | ***string***| ***(Optional)*** |
| `gatewayID` | ***string***| ***(Optional)*** |
| `instanceID` | ***string***| ***(Optional)*** |
| `ipv6CIDRBlock` | ***string***| ***(Optional)*** |
| `natGatewayID` | ***string***| ***(Optional)*** |
| `networkInterfaceID` | ***string***| ***(Optional)*** |
| `transitGatewayID` | ***string***| ***(Optional)*** |
| `vpcPeeringConnectionID` | ***string***| ***(Optional)*** |
## DefaultRouteTableStatus
##### (Appears on:[DefaultRouteTable](#DefaultRouteTable))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DefaultRouteTableSpec](#DefaultRouteTableSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
