---
title: ComputeTargetSSLProxy
menu:
  docs_{{ .version }}:
    identifier: computetargetsslproxy-google.kubeform.com
    name: ComputeTargetSSLProxy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeTargetSSLProxy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeTargetSSLProxy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeTargetSSLProxySpec](#ComputeTargetSSLProxySpec)***||
| `status` | ***[ComputeTargetSSLProxyStatus](#ComputeTargetSSLProxyStatus)***||
## ComputeTargetSSLProxySpec
##### (Appears on:[ComputeTargetSSLProxy](#ComputeTargetSSLProxy), [ComputeTargetSSLProxyStatus](#ComputeTargetSSLProxyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backendService` | ***string***||
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `proxyHeader` | ***string***| ***(Optional)*** |
| `proxyID` | ***int***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `sslCertificates` | ***[]string***||
| `sslPolicy` | ***string***| ***(Optional)*** |
## ComputeTargetSSLProxyStatus
##### (Appears on:[ComputeTargetSSLProxy](#ComputeTargetSSLProxy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeTargetSSLProxySpec](#ComputeTargetSSLProxySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
