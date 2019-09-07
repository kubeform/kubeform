---
title: NetworkInterfaceBackendAddressPoolAssociation
menu:
  docs_v0.0.1:
    identifier: networkinterfacebackendaddresspoolassociation-azurerm.kubeform.com
    name: NetworkInterfaceBackendAddressPoolAssociation
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## NetworkInterfaceBackendAddressPoolAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkInterfaceBackendAddressPoolAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkInterfaceBackendAddressPoolAssociationSpec](#NetworkInterfaceBackendAddressPoolAssociationSpec)***||
| `status` | ***[NetworkInterfaceBackendAddressPoolAssociationStatus](#NetworkInterfaceBackendAddressPoolAssociationStatus)***||
## NetworkInterfaceBackendAddressPoolAssociationSpec
##### (Appears on:[NetworkInterfaceBackendAddressPoolAssociation](#NetworkInterfaceBackendAddressPoolAssociation), [NetworkInterfaceBackendAddressPoolAssociationStatus](#NetworkInterfaceBackendAddressPoolAssociationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backendAddressPoolID` | ***string***||
| `ipConfigurationName` | ***string***||
| `networkInterfaceID` | ***string***||
## NetworkInterfaceBackendAddressPoolAssociationStatus
##### (Appears on:[NetworkInterfaceBackendAddressPoolAssociation](#NetworkInterfaceBackendAddressPoolAssociation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkInterfaceBackendAddressPoolAssociationSpec](#NetworkInterfaceBackendAddressPoolAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
