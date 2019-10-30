---
title: KeyPair
menu:
  docs_{{ .version }}:
    identifier: keypair-aws.kubeform.com
    name: KeyPair
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KeyPair
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `KeyPair` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KeyPairSpec](#keypairspec)***||
| `status` | ***[KeyPairStatus](#keypairstatus)***||
## KeyPairSpec

Appears on:[KeyPair](#keypair), [KeyPairStatus](#keypairstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `fingerprint` | ***string***| ***(Optional)*** |
| `keyName` | ***string***| ***(Optional)*** |
| `keyNamePrefix` | ***string***| ***(Optional)*** |
| `publicKey` | ***string***||
## KeyPairStatus

Appears on:[KeyPair](#keypair)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KeyPairSpec](#keypairspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[KeyPairStatus](#keypairstatus)

---
