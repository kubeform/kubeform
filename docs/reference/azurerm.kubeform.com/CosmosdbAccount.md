---
title: CosmosdbAccount
menu:
  docs_{{ .version }}:
    identifier: cosmosdbaccount-azurerm.kubeform.com
    name: CosmosdbAccount
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CosmosdbAccount
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `CosmosdbAccount` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CosmosdbAccountSpec](#CosmosdbAccountSpec)***||
| `status` | ***[CosmosdbAccountStatus](#CosmosdbAccountStatus)***||
## CosmosdbAccountSpec

Appears on:[CosmosdbAccount](#CosmosdbAccount), [CosmosdbAccountStatus](#CosmosdbAccountStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `capabilities` | ***[[]CosmosdbAccountSpecCapabilities](#CosmosdbAccountSpecCapabilities)***| ***(Optional)*** |
| `consistencyPolicy` | ***[[]CosmosdbAccountSpecConsistencyPolicy](#CosmosdbAccountSpecConsistencyPolicy)***| ***(Optional)*** |
| `enableAutomaticFailover` | ***bool***| ***(Optional)*** |
| `enableMultipleWriteLocations` | ***bool***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `failoverPolicy` | ***[[]CosmosdbAccountSpecFailoverPolicy](#CosmosdbAccountSpecFailoverPolicy)***| ***(Optional)*** Deprecated|
| `geoLocation` | ***[[]CosmosdbAccountSpecGeoLocation](#CosmosdbAccountSpecGeoLocation)***| ***(Optional)*** |
| `ipRangeFilter` | ***string***| ***(Optional)*** |
| `isVirtualNetworkFilterEnabled` | ***bool***| ***(Optional)*** |
| `kind` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `offerType` | ***string***||
| `readEndpoints` | ***[]string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `virtualNetworkRule` | ***[[]CosmosdbAccountSpecVirtualNetworkRule](#CosmosdbAccountSpecVirtualNetworkRule)***| ***(Optional)*** |
| `writeEndpoints` | ***[]string***| ***(Optional)*** |
## CosmosdbAccountSpecCapabilities

Appears on:[CosmosdbAccountSpec](#CosmosdbAccountSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
## CosmosdbAccountSpecConsistencyPolicy

Appears on:[CosmosdbAccountSpec](#CosmosdbAccountSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `consistencyLevel` | ***string***||
| `maxIntervalInSeconds` | ***int***| ***(Optional)*** |
| `maxStalenessPrefix` | ***int***| ***(Optional)*** |
## CosmosdbAccountSpecFailoverPolicy

Appears on:[CosmosdbAccountSpec](#CosmosdbAccountSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `priority` | ***int***||
## CosmosdbAccountSpecGeoLocation

Appears on:[CosmosdbAccountSpec](#CosmosdbAccountSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `failoverPriority` | ***int***||
| `ID` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `prefix` | ***string***| ***(Optional)*** |
## CosmosdbAccountSpecVirtualNetworkRule

Appears on:[CosmosdbAccountSpec](#CosmosdbAccountSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***||
## CosmosdbAccountStatus

Appears on:[CosmosdbAccount](#CosmosdbAccount)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CosmosdbAccountSpec](#CosmosdbAccountSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_master_key` | ***string*** ||
| `primary_readonly_master_key` | ***string*** ||
| `secondary_master_key` | ***string*** ||
| `secondary_readonly_master_key` | ***string*** ||
