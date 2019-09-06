## DxGatewayAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DxGatewayAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DxGatewayAssociationSpec](#DxGatewayAssociationSpec)***||
| `status` | ***[DxGatewayAssociationStatus](#DxGatewayAssociationStatus)***||
## DxGatewayAssociationSpec
##### (Appears on:[DxGatewayAssociation](#DxGatewayAssociation), [DxGatewayAssociationStatus](#DxGatewayAssociationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allowedPrefixes` | ***[]string***| ***(Optional)*** |
| `dxGatewayAssociationID` | ***string***| ***(Optional)*** |
| `dxGatewayID` | ***string***||
| `vpnGatewayID` | ***string***||
## DxGatewayAssociationStatus
##### (Appears on:[DxGatewayAssociation](#DxGatewayAssociation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DxGatewayAssociationSpec](#DxGatewayAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
