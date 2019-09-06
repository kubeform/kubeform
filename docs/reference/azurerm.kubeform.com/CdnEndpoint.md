## CdnEndpoint
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `CdnEndpoint` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CdnEndpointSpec](#CdnEndpointSpec)***||
| `status` | ***[CdnEndpointStatus](#CdnEndpointStatus)***||
## CdnEndpointSpec
##### (Appears on:[CdnEndpoint](#CdnEndpoint), [CdnEndpointStatus](#CdnEndpointStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `contentTypesToCompress` | ***[]string***| ***(Optional)*** |
| `geoFilter` | ***[[]CdnEndpointSpecGeoFilter](#CdnEndpointSpecGeoFilter)***| ***(Optional)*** |
| `hostName` | ***string***| ***(Optional)*** |
| `isCompressionEnabled` | ***bool***| ***(Optional)*** |
| `isHTTPAllowed` | ***bool***| ***(Optional)*** |
| `isHTTPSAllowed` | ***bool***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `optimizationType` | ***string***| ***(Optional)*** |
| `origin` | ***[[]CdnEndpointSpecOrigin](#CdnEndpointSpecOrigin)***||
| `originHostHeader` | ***string***| ***(Optional)*** |
| `originPath` | ***string***| ***(Optional)*** |
| `probePath` | ***string***| ***(Optional)*** |
| `profileName` | ***string***||
| `querystringCachingBehaviour` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## CdnEndpointSpecGeoFilter
##### (Appears on:[CdnEndpointSpec](#CdnEndpointSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***||
| `countryCodes` | ***[]string***||
| `relativePath` | ***string***||
## CdnEndpointSpecOrigin
##### (Appears on:[CdnEndpointSpec](#CdnEndpointSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `hostName` | ***string***||
| `httpPort` | ***int***| ***(Optional)*** |
| `httpsPort` | ***int***| ***(Optional)*** |
| `name` | ***string***||
## CdnEndpointStatus
##### (Appears on:[CdnEndpoint](#CdnEndpoint))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CdnEndpointSpec](#CdnEndpointSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
