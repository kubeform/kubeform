---
title: KmsCryptoKeyIamMember
menu:
  docs_v0.0.1:
    identifier: kmscryptokeyiammember-google.kubeform.com
    name: KmsCryptoKeyIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## KmsCryptoKeyIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `KmsCryptoKeyIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KmsCryptoKeyIamMemberSpec](#KmsCryptoKeyIamMemberSpec)***||
| `status` | ***[KmsCryptoKeyIamMemberStatus](#KmsCryptoKeyIamMemberStatus)***||
## KmsCryptoKeyIamMemberSpec
##### (Appears on:[KmsCryptoKeyIamMember](#KmsCryptoKeyIamMember), [KmsCryptoKeyIamMemberStatus](#KmsCryptoKeyIamMemberStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cryptoKeyID` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `member` | ***string***||
| `role` | ***string***||
## KmsCryptoKeyIamMemberStatus
##### (Appears on:[KmsCryptoKeyIamMember](#KmsCryptoKeyIamMember))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KmsCryptoKeyIamMemberSpec](#KmsCryptoKeyIamMemberSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
