---
title: KmsCryptoKeyIamBinding
menu:
  docs_{{ .version }}:
    identifier: kmscryptokeyiambinding-google.kubeform.com
    name: KmsCryptoKeyIamBinding
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KmsCryptoKeyIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `KmsCryptoKeyIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KmsCryptoKeyIamBindingSpec](#kmscryptokeyiambindingspec)***||
| `status` | ***[KmsCryptoKeyIamBindingStatus](#kmscryptokeyiambindingstatus)***||
## KmsCryptoKeyIamBindingSpec

Appears on:[KmsCryptoKeyIamBinding](#kmscryptokeyiambinding), [KmsCryptoKeyIamBindingStatus](#kmscryptokeyiambindingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cryptoKeyID` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `members` | ***[]string***||
| `role` | ***string***||
## KmsCryptoKeyIamBindingStatus

Appears on:[KmsCryptoKeyIamBinding](#kmscryptokeyiambinding)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KmsCryptoKeyIamBindingSpec](#kmscryptokeyiambindingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
