---
title: KmsCryptoKey
menu:
  docs_v0.0.1:
    identifier: kmscryptokey-google.kubeform.com
    name: KmsCryptoKey
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## KmsCryptoKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `KmsCryptoKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KmsCryptoKeySpec](#KmsCryptoKeySpec)***||
| `status` | ***[KmsCryptoKeyStatus](#KmsCryptoKeyStatus)***||
## KmsCryptoKeySpec
##### (Appears on:[KmsCryptoKey](#KmsCryptoKey), [KmsCryptoKeyStatus](#KmsCryptoKeyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `keyRing` | ***string***||
| `name` | ***string***||
| `rotationPeriod` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
## KmsCryptoKeyStatus
##### (Appears on:[KmsCryptoKey](#KmsCryptoKey))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KmsCryptoKeySpec](#KmsCryptoKeySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
