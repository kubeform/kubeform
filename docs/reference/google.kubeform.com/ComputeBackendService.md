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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeBackendServiceSpec](#computebackendservicespec)***||
| `status` | ***[ComputeBackendServiceStatus](#computebackendservicestatus)***||
## ComputeBackendServiceSpec

Appears on:[ComputeBackendService](#computebackendservice), [ComputeBackendServiceStatus](#computebackendservicestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `backend` | ***[[]ComputeBackendServiceSpecBackend](#computebackendservicespecbackend)***| ***(Optional)*** |
| `cdnPolicy` | ***[[]ComputeBackendServiceSpecCdnPolicy](#computebackendservicespeccdnpolicy)***| ***(Optional)*** |
| `connectionDrainingTimeoutSec` | ***int***| ***(Optional)*** |
| `customRequestHeaders` | ***[]string***| ***(Optional)*** Deprecated|
| `description` | ***string***| ***(Optional)*** |
| `enableCdn` | ***bool***| ***(Optional)*** |
| `fingerprint` | ***string***| ***(Optional)*** |
| `healthChecks` | ***[]string***||
| `iap` | ***[[]ComputeBackendServiceSpecIap](#computebackendservicespeciap)***| ***(Optional)*** |
| `name` | ***string***||
| `portName` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** |
| `securityPolicy` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `sessionAffinity` | ***string***| ***(Optional)*** |
| `timeoutSec` | ***int***| ***(Optional)*** |
## ComputeBackendServiceSpecBackend

Appears on:[ComputeBackendServiceSpec](#computebackendservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `balancingMode` | ***string***| ***(Optional)*** |
| `capacityScaler` | ***encoding/json.Number***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `group` | ***string***| ***(Optional)*** |
| `maxConnections` | ***int***| ***(Optional)*** |
| `maxConnectionsPerInstance` | ***int***| ***(Optional)*** |
| `maxRate` | ***int***| ***(Optional)*** |
| `maxRatePerInstance` | ***encoding/json.Number***| ***(Optional)*** |
| `maxUtilization` | ***encoding/json.Number***| ***(Optional)*** |
## ComputeBackendServiceSpecCdnPolicy

Appears on:[ComputeBackendServiceSpec](#computebackendservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cacheKeyPolicy` | ***[[]ComputeBackendServiceSpecCdnPolicyCacheKeyPolicy](#computebackendservicespeccdnpolicycachekeypolicy)***| ***(Optional)*** |
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
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `iap.<index>.oauth2_client_secret` | ***string*** ||
