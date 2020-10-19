---
title: ComputeBackendServiceSignedURLKey
menu:
  docs_{{ .version }}:
    identifier: computebackendservicesignedurlkey-google.kubeform.com
    name: ComputeBackendServiceSignedURLKey
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeBackendServiceSignedURLKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeBackendServiceSignedURLKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeBackendServiceSignedURLKeySpec](#computebackendservicesignedurlkeyspec)***||
| `status` | ***[ComputeBackendServiceSignedURLKeyStatus](#computebackendservicesignedurlkeystatus)***||
## ComputeBackendServiceSignedURLKeySpec

Appears on:[ComputeBackendServiceSignedURLKey](#computebackendservicesignedurlkey), [ComputeBackendServiceSignedURLKeyStatus](#computebackendservicesignedurlkeystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `backendService` | ***string***||
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
## ComputeBackendServiceSignedURLKeyStatus

Appears on:[ComputeBackendServiceSignedURLKey](#computebackendservicesignedurlkey)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeBackendServiceSignedURLKeySpec](#computebackendservicesignedurlkeyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeBackendServiceSignedURLKeyStatus](#computebackendservicesignedurlkeystatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `key_value` | ***string*** ||
