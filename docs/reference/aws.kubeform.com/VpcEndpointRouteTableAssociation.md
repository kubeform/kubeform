## VpcEndpointRouteTableAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `VpcEndpointRouteTableAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpcEndpointRouteTableAssociationSpec](#VpcEndpointRouteTableAssociationSpec)***||
| `status` | ***[VpcEndpointRouteTableAssociationStatus](#VpcEndpointRouteTableAssociationStatus)***||
## VpcEndpointRouteTableAssociationSpec
##### (Appears on:[VpcEndpointRouteTableAssociation](#VpcEndpointRouteTableAssociation), [VpcEndpointRouteTableAssociationStatus](#VpcEndpointRouteTableAssociationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `routeTableID` | ***string***||
| `vpcEndpointID` | ***string***||
## VpcEndpointRouteTableAssociationStatus
##### (Appears on:[VpcEndpointRouteTableAssociation](#VpcEndpointRouteTableAssociation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpcEndpointRouteTableAssociationSpec](#VpcEndpointRouteTableAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
