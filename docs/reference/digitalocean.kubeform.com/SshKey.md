---
title: SshKey
menu:
  docs_{{ .version }}:
    identifier: sshkey-digitalocean.kubeform.com
    name: SshKey
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SshKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `SshKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SshKeySpec](#sshkeyspec)***||
| `status` | ***[SshKeyStatus](#sshkeystatus)***||
## Phase(`string` alias)

Appears on:[SshKeyStatus](#sshkeystatus)

## SshKeySpec

Appears on:[SshKey](#sshkey), [SshKeyStatus](#sshkeystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `fingerprint` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `publicKey` | ***string***||
## SshKeyStatus

Appears on:[SshKey](#sshkey)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SshKeySpec](#sshkeyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
