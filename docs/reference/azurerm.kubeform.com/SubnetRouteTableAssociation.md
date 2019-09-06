## SubnetRouteTableAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SubnetRouteTableAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SubnetRouteTableAssociationSpec](#SubnetRouteTableAssociationSpec)***||
| `status` | ***[SubnetRouteTableAssociationStatus](#SubnetRouteTableAssociationStatus)***||
## SubnetRouteTableAssociationSpec
##### (Appears on:[SubnetRouteTableAssociation](#SubnetRouteTableAssociation), [SubnetRouteTableAssociationStatus](#SubnetRouteTableAssociationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `routeTableID` | ***string***||
| `subnetID` | ***string***||
## SubnetRouteTableAssociationStatus
##### (Appears on:[SubnetRouteTableAssociation](#SubnetRouteTableAssociation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SubnetRouteTableAssociationSpec](#SubnetRouteTableAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
