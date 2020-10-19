---
title: ComputeBackendService
menu:
  docs_{{ .version }}:
    identifier: computebackendservice-google.kubeform.com
    name: ComputeBackendService
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeBackendService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeBackendService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeBackendServiceSpec](#computebackendservicespec)***||
| `status` | ***[ComputeBackendServiceStatus](#computebackendservicestatus)***||
## ComputeBackendServiceSpec

Appears on:[ComputeBackendService](#computebackendservice), [ComputeBackendServiceStatus](#computebackendservicestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `affinityCookieTtlSec` | ***int64***| ***(Optional)*** |
| `backend` | ***[[]ComputeBackendServiceSpecBackend](#computebackendservicespecbackend)***| ***(Optional)*** |
| `cdnPolicy` | ***[[]ComputeBackendServiceSpecCdnPolicy](#computebackendservicespeccdnpolicy)***| ***(Optional)*** |
| `connectionDrainingTimeoutSec` | ***int64***| ***(Optional)*** |
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `enableCdn` | ***bool***| ***(Optional)*** |
| `fingerprint` | ***string***| ***(Optional)*** |
| `healthChecks` | ***[]string***||
| `iap` | ***[[]ComputeBackendServiceSpecIap](#computebackendservicespeciap)***| ***(Optional)*** |
| `loadBalancingScheme` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `portName` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** |
| `securityPolicy` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `sessionAffinity` | ***string***| ***(Optional)*** |
| `timeoutSec` | ***int64***| ***(Optional)*** |
## ComputeBackendServiceSpecBackend

Appears on:[ComputeBackendServiceSpec](#computebackendservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `balancingMode` | ***string***| ***(Optional)*** |
| `capacityScaler` | ***float64***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `group` | ***string***| ***(Optional)*** |
| `maxConnections` | ***int64***| ***(Optional)*** |
| `maxConnectionsPerEndpoint` | ***int64***| ***(Optional)*** |
| `maxConnectionsPerInstance` | ***int64***| ***(Optional)*** |
| `maxRate` | ***int64***| ***(Optional)*** |
| `maxRatePerEndpoint` | ***float64***| ***(Optional)*** |
| `maxRatePerInstance` | ***float64***| ***(Optional)*** |
| `maxUtilization` | ***float64***| ***(Optional)*** |
## ComputeBackendServiceSpecCdnPolicy

Appears on:[ComputeBackendServiceSpec](#computebackendservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cacheKeyPolicy` | ***[[]ComputeBackendServiceSpecCdnPolicyCacheKeyPolicy](#computebackendservicespeccdnpolicycachekeypolicy)***| ***(Optional)*** |
| `signedURLCacheMaxAgeSec` | ***int64***| ***(Optional)*** |
## ComputeBackendServiceSpecCdnPolicyCacheKeyPolicy

Appears on:[ComputeBackendServiceSpecCdnPolicy](#computebackendservicespeccdnpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `includeHost` | ***bool***| ***(Optional)*** |
| `includeProtocol` | ***bool***| ***(Optional)*** |
| `includeQueryString` | ***bool***| ***(Optional)*** |
| `queryStringBlacklist` | ***[]string***| ***(Optional)*** |
| `queryStringWhitelist` | ***[]string***| ***(Optional)*** |
## ComputeBackendServiceSpecIap

Appears on:[ComputeBackendServiceSpec](#computebackendservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `oauth2ClientID` | ***string***||
## ComputeBackendServiceStatus

Appears on:[ComputeBackendService](#computebackendservice)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeBackendServiceSpec](#computebackendservicespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeBackendServiceStatus](#computebackendservicestatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `iap.<index>.oauth2_client_secret` | ***string*** ||
| `iap.<index>.oauth2_client_secret_sha256` | ***string*** ||
