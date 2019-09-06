## Ec2TransitGatewayRouteTableAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Ec2TransitGatewayRouteTableAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Ec2TransitGatewayRouteTableAssociationSpec](#Ec2TransitGatewayRouteTableAssociationSpec)***||
| `status` | ***[Ec2TransitGatewayRouteTableAssociationStatus](#Ec2TransitGatewayRouteTableAssociationStatus)***||
## Ec2TransitGatewayRouteTableAssociationSpec
##### (Appears on:[Ec2TransitGatewayRouteTableAssociation](#Ec2TransitGatewayRouteTableAssociation), [Ec2TransitGatewayRouteTableAssociationStatus](#Ec2TransitGatewayRouteTableAssociationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `resourceID` | ***string***| ***(Optional)*** |
| `resourceType` | ***string***| ***(Optional)*** |
| `transitGatewayAttachmentID` | ***string***||
| `transitGatewayRouteTableID` | ***string***||
## Ec2TransitGatewayRouteTableAssociationStatus
##### (Appears on:[Ec2TransitGatewayRouteTableAssociation](#Ec2TransitGatewayRouteTableAssociation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Ec2TransitGatewayRouteTableAssociationSpec](#Ec2TransitGatewayRouteTableAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
