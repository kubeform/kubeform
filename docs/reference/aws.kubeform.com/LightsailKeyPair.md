---
title: LightsailKeyPair
menu:
  docs_{{ .version }}:
    identifier: lightsailkeypair-aws.kubeform.com
    name: LightsailKeyPair
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LightsailKeyPair
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LightsailKeyPair` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LightsailKeyPairSpec](#LightsailKeyPairSpec)***||
| `status` | ***[LightsailKeyPairStatus](#LightsailKeyPairStatus)***||
## LightsailKeyPairSpec
##### (Appears on:[LightsailKeyPair](#LightsailKeyPair), [LightsailKeyPairStatus](#LightsailKeyPairStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `encryptedFingerprint` | ***string***| ***(Optional)*** |
| `encryptedPrivateKey` | ***string***| ***(Optional)*** |
| `fingerprint` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `pgpKey` | ***string***| ***(Optional)*** |
| `privateKey` | ***string***| ***(Optional)*** |
| `publicKey` | ***string***| ***(Optional)*** |
## LightsailKeyPairStatus
##### (Appears on:[LightsailKeyPair](#LightsailKeyPair))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LightsailKeyPairSpec](#LightsailKeyPairSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
