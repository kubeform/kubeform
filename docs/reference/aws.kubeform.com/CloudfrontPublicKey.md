---
title: CloudfrontPublicKey
menu:
  docs_v0.0.1:
    identifier: cloudfrontpublickey-aws.kubeform.com
    name: CloudfrontPublicKey
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## CloudfrontPublicKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudfrontPublicKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudfrontPublicKeySpec](#CloudfrontPublicKeySpec)***||
| `status` | ***[CloudfrontPublicKeyStatus](#CloudfrontPublicKeyStatus)***||
## CloudfrontPublicKeySpec
##### (Appears on:[CloudfrontPublicKey](#CloudfrontPublicKey), [CloudfrontPublicKeyStatus](#CloudfrontPublicKeyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `callerReference` | ***string***| ***(Optional)*** |
| `comment` | ***string***| ***(Optional)*** |
| `encodedKey` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
## CloudfrontPublicKeyStatus
##### (Appears on:[CloudfrontPublicKey](#CloudfrontPublicKey))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudfrontPublicKeySpec](#CloudfrontPublicKeySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
