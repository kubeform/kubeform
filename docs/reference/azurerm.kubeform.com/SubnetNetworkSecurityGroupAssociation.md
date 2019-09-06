## SubnetNetworkSecurityGroupAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SubnetNetworkSecurityGroupAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SubnetNetworkSecurityGroupAssociationSpec](#SubnetNetworkSecurityGroupAssociationSpec)***||
| `status` | ***[SubnetNetworkSecurityGroupAssociationStatus](#SubnetNetworkSecurityGroupAssociationStatus)***||
## SubnetNetworkSecurityGroupAssociationSpec
##### (Appears on:[SubnetNetworkSecurityGroupAssociation](#SubnetNetworkSecurityGroupAssociation), [SubnetNetworkSecurityGroupAssociationStatus](#SubnetNetworkSecurityGroupAssociationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `networkSecurityGroupID` | ***string***||
| `subnetID` | ***string***||
## SubnetNetworkSecurityGroupAssociationStatus
##### (Appears on:[SubnetNetworkSecurityGroupAssociation](#SubnetNetworkSecurityGroupAssociation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SubnetNetworkSecurityGroupAssociationSpec](#SubnetNetworkSecurityGroupAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
