---
title: Iothub
menu:
  docs_{{ .version }}:
    identifier: iothub-azurerm.kubeform.com
    name: Iothub
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Iothub
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `Iothub` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IothubSpec](#iothubspec)***||
| `status` | ***[IothubStatus](#iothubstatus)***||
## IothubSpec

Appears on:[Iothub](#iothub), [IothubStatus](#iothubstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `endpoint` | ***[[]IothubSpecEndpoint](#iothubspecendpoint)***| ***(Optional)*** |
| `eventHubEventsEndpoint` | ***string***| ***(Optional)*** |
| `eventHubEventsPath` | ***string***| ***(Optional)*** |
| `eventHubOperationsEndpoint` | ***string***| ***(Optional)*** |
| `eventHubOperationsPath` | ***string***| ***(Optional)*** |
| `fallbackRoute` | ***[[]IothubSpecFallbackRoute](#iothubspecfallbackroute)***| ***(Optional)*** |
| `fileUpload` | ***[[]IothubSpecFileUpload](#iothubspecfileupload)***| ***(Optional)*** |
| `hostname` | ***string***| ***(Optional)*** |
| `ipFilterRule` | ***[[]IothubSpecIpFilterRule](#iothubspecipfilterrule)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `route` | ***[[]IothubSpecRoute](#iothubspecroute)***| ***(Optional)*** |
| `sharedAccessPolicy` | ***[[]IothubSpecSharedAccessPolicy](#iothubspecsharedaccesspolicy)***| ***(Optional)*** |
| `sku` | ***[[]IothubSpecSku](#iothubspecsku)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## IothubSpecEndpoint

Appears on:[IothubSpec](#iothubspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `batchFrequencyInSeconds` | ***int***| ***(Optional)*** |
| `containerName` | ***string***| ***(Optional)*** |
| `encoding` | ***string***| ***(Optional)*** |
| `fileNameFormat` | ***string***| ***(Optional)*** |
| `maxChunkSizeInBytes` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `type` | ***string***||
## IothubSpecFallbackRoute

Appears on:[IothubSpec](#iothubspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `condition` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `endpointNames` | ***[]string***| ***(Optional)*** |
| `source` | ***string***| ***(Optional)*** |
## IothubSpecFileUpload

Appears on:[IothubSpec](#iothubspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `containerName` | ***string***||
| `defaultTtl` | ***string***| ***(Optional)*** |
| `lockDuration` | ***string***| ***(Optional)*** |
| `maxDeliveryCount` | ***int***| ***(Optional)*** |
| `notifications` | ***bool***| ***(Optional)*** |
| `sasTtl` | ***string***| ***(Optional)*** |
## IothubSpecIpFilterRule

Appears on:[IothubSpec](#iothubspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***||
| `ipMask` | ***string***||
| `name` | ***string***||
## IothubSpecRoute

Appears on:[IothubSpec](#iothubspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `condition` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***||
| `endpointNames` | ***[]string***||
| `name` | ***string***||
| `source` | ***string***||
## IothubSpecSharedAccessPolicy

Appears on:[IothubSpec](#iothubspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `keyName` | ***string***| ***(Optional)*** |
| `permissions` | ***string***| ***(Optional)*** |
## IothubSpecSku

Appears on:[IothubSpec](#iothubspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int***||
| `name` | ***string***||
| `tier` | ***string***||
## IothubStatus

Appears on:[Iothub](#iothub)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IothubSpec](#iothubspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IothubStatus](#iothubstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `endpoint.<index>.connection_string` | ***string*** ||
| `file_upload.<index>.connection_string` | ***string*** ||
| `shared_access_policy.<index>.primary_key` | ***string*** ||
| `shared_access_policy.<index>.secondary_key` | ***string*** ||
