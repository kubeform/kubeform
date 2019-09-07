---
title: Subnet
menu:
  docs_v0.0.1:
    identifier: subnet-azurerm.kubeform.com
    name: Subnet
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## Subnet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `Subnet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SubnetSpec](#SubnetSpec)***||
| `status` | ***[SubnetStatus](#SubnetStatus)***||
## SubnetSpec
##### (Appears on:[Subnet](#Subnet), [SubnetStatus](#SubnetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `addressPrefix` | ***string***||
| `delegation` | ***[[]SubnetSpecDelegation](#SubnetSpecDelegation)***| ***(Optional)*** |
| `ipConfigurations` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `networkSecurityGroupID` | ***string***| ***(Optional)*** Deprecated|
| `resourceGroupName` | ***string***||
| `routeTableID` | ***string***| ***(Optional)*** Deprecated|
| `serviceEndpoints` | ***[]string***| ***(Optional)*** |
| `virtualNetworkName` | ***string***||
## SubnetSpecDelegation
##### (Appears on:[SubnetSpec](#SubnetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `serviceDelegation` | ***[[]SubnetSpecDelegationServiceDelegation](#SubnetSpecDelegationServiceDelegation)***||
## SubnetSpecDelegationServiceDelegation
##### (Appears on:[SubnetSpecDelegation](#SubnetSpecDelegation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `actions` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
## SubnetStatus
##### (Appears on:[Subnet](#Subnet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SubnetSpec](#SubnetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
