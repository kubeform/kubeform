---
title: CdnEndpoint
menu:
  docs_{{ .version }}:
    identifier: cdnendpoint-azurerm.kubeform.com
    name: CdnEndpoint
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CdnEndpoint
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `CdnEndpoint` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CdnEndpointSpec](#cdnendpointspec)***||
| `status` | ***[CdnEndpointStatus](#cdnendpointstatus)***||
## CdnEndpointSpec

Appears on:[CdnEndpoint](#cdnendpoint), [CdnEndpointStatus](#cdnendpointstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `contentTypesToCompress` | ***[]string***| ***(Optional)*** |
| `geoFilter` | ***[[]CdnEndpointSpecGeoFilter](#cdnendpointspecgeofilter)***| ***(Optional)*** |
| `hostName` | ***string***| ***(Optional)*** |
| `isCompressionEnabled` | ***bool***| ***(Optional)*** |
| `isHTTPAllowed` | ***bool***| ***(Optional)*** |
| `isHTTPSAllowed` | ***bool***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `optimizationType` | ***string***| ***(Optional)*** |
| `origin` | ***[[]CdnEndpointSpecOrigin](#cdnendpointspecorigin)***||
| `originHostHeader` | ***string***| ***(Optional)*** |
| `originPath` | ***string***| ***(Optional)*** |
| `probePath` | ***string***| ***(Optional)*** |
| `profileName` | ***string***||
| `querystringCachingBehaviour` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## CdnEndpointSpecGeoFilter

Appears on:[CdnEndpointSpec](#cdnendpointspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***||
| `countryCodes` | ***[]string***||
| `relativePath` | ***string***||
## CdnEndpointSpecOrigin

Appears on:[CdnEndpointSpec](#cdnendpointspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hostName` | ***string***||
| `httpPort` | ***int***| ***(Optional)*** |
| `httpsPort` | ***int***| ***(Optional)*** |
| `name` | ***string***||
## CdnEndpointStatus

Appears on:[CdnEndpoint](#cdnendpoint)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CdnEndpointSpec](#cdnendpointspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CdnEndpointStatus](#cdnendpointstatus)

---
