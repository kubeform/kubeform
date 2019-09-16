---
title: AlbListenerCertificate
menu:
  docs_{{ .version }}:
    identifier: alblistenercertificate-aws.kubeform.com
    name: AlbListenerCertificate
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AlbListenerCertificate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AlbListenerCertificate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AlbListenerCertificateSpec](#alblistenercertificatespec)***||
| `status` | ***[AlbListenerCertificateStatus](#alblistenercertificatestatus)***||
## AlbListenerCertificateSpec

Appears on:[AlbListenerCertificate](#alblistenercertificate), [AlbListenerCertificateStatus](#alblistenercertificatestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `certificateArn` | ***string***||
| `listenerArn` | ***string***||
## AlbListenerCertificateStatus

Appears on:[AlbListenerCertificate](#alblistenercertificate)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AlbListenerCertificateSpec](#alblistenercertificatespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
