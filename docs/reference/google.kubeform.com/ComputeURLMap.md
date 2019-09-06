## ComputeURLMap
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeURLMap` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeURLMapSpec](#ComputeURLMapSpec)***||
| `status` | ***[ComputeURLMapStatus](#ComputeURLMapStatus)***||
## ComputeURLMapSpec
##### (Appears on:[ComputeURLMap](#ComputeURLMap), [ComputeURLMapStatus](#ComputeURLMapStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `defaultService` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `fingerprint` | ***string***| ***(Optional)*** |
| `hostRule` | ***[[]ComputeURLMapSpecHostRule](#ComputeURLMapSpecHostRule)***| ***(Optional)*** |
| `mapID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `pathMatcher` | ***[[]ComputeURLMapSpecPathMatcher](#ComputeURLMapSpecPathMatcher)***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `test` | ***[[]ComputeURLMapSpecTest](#ComputeURLMapSpecTest)***| ***(Optional)*** |
## ComputeURLMapSpecHostRule
##### (Appears on:[ComputeURLMapSpec](#ComputeURLMapSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `description` | ***string***| ***(Optional)*** |
| `hosts` | ***[]string***||
| `pathMatcher` | ***string***||
## ComputeURLMapSpecPathMatcher
##### (Appears on:[ComputeURLMapSpec](#ComputeURLMapSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultService` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `pathRule` | ***[[]ComputeURLMapSpecPathMatcherPathRule](#ComputeURLMapSpecPathMatcherPathRule)***| ***(Optional)*** |
## ComputeURLMapSpecPathMatcherPathRule
##### (Appears on:[ComputeURLMapSpecPathMatcher](#ComputeURLMapSpecPathMatcher))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `paths` | ***[]string***||
| `service` | ***string***||
## ComputeURLMapSpecTest
##### (Appears on:[ComputeURLMapSpec](#ComputeURLMapSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `description` | ***string***| ***(Optional)*** |
| `host` | ***string***||
| `path` | ***string***||
| `service` | ***string***||
## ComputeURLMapStatus
##### (Appears on:[ComputeURLMap](#ComputeURLMap))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeURLMapSpec](#ComputeURLMapSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
