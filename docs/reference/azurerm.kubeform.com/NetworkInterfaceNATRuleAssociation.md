## NetworkInterfaceNATRuleAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkInterfaceNATRuleAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkInterfaceNATRuleAssociationSpec](#NetworkInterfaceNATRuleAssociationSpec)***||
| `status` | ***[NetworkInterfaceNATRuleAssociationStatus](#NetworkInterfaceNATRuleAssociationStatus)***||
## NetworkInterfaceNATRuleAssociationSpec
##### (Appears on:[NetworkInterfaceNATRuleAssociation](#NetworkInterfaceNATRuleAssociation), [NetworkInterfaceNATRuleAssociationStatus](#NetworkInterfaceNATRuleAssociationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `ipConfigurationName` | ***string***||
| `natRuleID` | ***string***||
| `networkInterfaceID` | ***string***||
## NetworkInterfaceNATRuleAssociationStatus
##### (Appears on:[NetworkInterfaceNATRuleAssociation](#NetworkInterfaceNATRuleAssociation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkInterfaceNATRuleAssociationSpec](#NetworkInterfaceNATRuleAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
