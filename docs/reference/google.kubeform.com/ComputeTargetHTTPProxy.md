---
title: ComputeTargetHTTPProxy
menu:
  docs_{{ .version }}:
    identifier: computetargethttpproxy-google.kubeform.com
    name: ComputeTargetHTTPProxy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeTargetHTTPProxy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeTargetHTTPProxy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeTargetHTTPProxySpec](#ComputeTargetHTTPProxySpec)***||
| `status` | ***[ComputeTargetHTTPProxyStatus](#ComputeTargetHTTPProxyStatus)***||
## ComputeTargetHTTPProxySpec
##### (Appears on:[ComputeTargetHTTPProxy](#ComputeTargetHTTPProxy), [ComputeTargetHTTPProxyStatus](#ComputeTargetHTTPProxyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `proxyID` | ***int***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `urlMap` | ***string***||
## ComputeTargetHTTPProxyStatus
##### (Appears on:[ComputeTargetHTTPProxy](#ComputeTargetHTTPProxy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeTargetHTTPProxySpec](#ComputeTargetHTTPProxySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
