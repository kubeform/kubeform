---
title: Subnet
menu:
  docs_{{ .version }}:
    identifier: subnet-azurerm.kubeform.com
    name: Subnet
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Subnet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `Subnet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SubnetSpec](#subnetspec)***||
| `status` | ***[SubnetStatus](#subnetstatus)***||
## Phase(`string` alias)

Appears on:[SubnetStatus](#subnetstatus)

## SubnetSpec

Appears on:[Subnet](#subnet), [SubnetStatus](#subnetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `addressPrefix` | ***string***||
| `delegation` | ***[[]SubnetSpecDelegation](#subnetspecdelegation)***| ***(Optional)*** |
| `enforcePrivateLinkEndpointNetworkPolicies` | ***bool***| ***(Optional)*** |
| `enforcePrivateLinkServiceNetworkPolicies` | ***bool***| ***(Optional)*** |
| `ipConfigurations` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `networkSecurityGroupID` | ***string***| ***(Optional)*** Deprecated|
| `resourceGroupName` | ***string***||
| `routeTableID` | ***string***| ***(Optional)*** Deprecated|
| `serviceEndpoints` | ***[]string***| ***(Optional)*** |
| `virtualNetworkName` | ***string***||
## SubnetSpecDelegation

Appears on:[SubnetSpec](#subnetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `serviceDelegation` | ***[[]SubnetSpecDelegationServiceDelegation](#subnetspecdelegationservicedelegation)***||
## SubnetSpecDelegationServiceDelegation

Appears on:[SubnetSpecDelegation](#subnetspecdelegation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `actions` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
## SubnetStatus

Appears on:[Subnet](#subnet)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SubnetSpec](#subnetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
