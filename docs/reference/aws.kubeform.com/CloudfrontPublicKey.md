---
title: CloudfrontPublicKey
menu:
  docs_{{ .version }}:
    identifier: cloudfrontpublickey-aws.kubeform.com
    name: CloudfrontPublicKey
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudfrontPublicKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudfrontPublicKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudfrontPublicKeySpec](#cloudfrontpublickeyspec)***||
| `status` | ***[CloudfrontPublicKeyStatus](#cloudfrontpublickeystatus)***||
## CloudfrontPublicKeySpec

Appears on:[CloudfrontPublicKey](#cloudfrontpublickey), [CloudfrontPublicKeyStatus](#cloudfrontpublickeystatus)

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

Appears on:[CloudfrontPublicKey](#cloudfrontpublickey)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudfrontPublicKeySpec](#cloudfrontpublickeyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CloudfrontPublicKeyStatus](#cloudfrontpublickeystatus)

---
