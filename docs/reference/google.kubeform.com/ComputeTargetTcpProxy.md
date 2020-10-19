---
title: ComputeTargetTcpProxy
menu:
  docs_{{ .version }}:
    identifier: computetargettcpproxy-google.kubeform.com
    name: ComputeTargetTcpProxy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeTargetTcpProxy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeTargetTcpProxy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeTargetTcpProxySpec](#computetargettcpproxyspec)***||
| `status` | ***[ComputeTargetTcpProxyStatus](#computetargettcpproxystatus)***||
## ComputeTargetTcpProxySpec

Appears on:[ComputeTargetTcpProxy](#computetargettcpproxy), [ComputeTargetTcpProxyStatus](#computetargettcpproxystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backendService` | ***string***||
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `proxyHeader` | ***string***| ***(Optional)*** |
| `proxyID` | ***int64***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
## ComputeTargetTcpProxyStatus

Appears on:[ComputeTargetTcpProxy](#computetargettcpproxy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeTargetTcpProxySpec](#computetargettcpproxyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeTargetTcpProxyStatus](#computetargettcpproxystatus)

---
