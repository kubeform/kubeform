---
title: KmsCryptoKey
menu:
  docs_{{ .version }}:
    identifier: kmscryptokey-google.kubeform.com
    name: KmsCryptoKey
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KmsCryptoKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `KmsCryptoKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KmsCryptoKeySpec](#kmscryptokeyspec)***||
| `status` | ***[KmsCryptoKeyStatus](#kmscryptokeystatus)***||
## KmsCryptoKeySpec

Appears on:[KmsCryptoKey](#kmscryptokey), [KmsCryptoKeyStatus](#kmscryptokeystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `keyRing` | ***string***||
| `labels` | ***map[string]string***| ***(Optional)*** |
| `name` | ***string***||
| `purpose` | ***string***| ***(Optional)*** |
| `rotationPeriod` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `versionTemplate` | ***[[]KmsCryptoKeySpecVersionTemplate](#kmscryptokeyspecversiontemplate)***| ***(Optional)*** |
## KmsCryptoKeySpecVersionTemplate

Appears on:[KmsCryptoKeySpec](#kmscryptokeyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `algorithm` | ***string***||
| `protectionLevel` | ***string***| ***(Optional)*** |
## KmsCryptoKeyStatus

Appears on:[KmsCryptoKey](#kmscryptokey)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KmsCryptoKeySpec](#kmscryptokeyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[KmsCryptoKeyStatus](#kmscryptokeystatus)

---
