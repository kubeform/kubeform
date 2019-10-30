---
title: TransferSSHKey
menu:
  docs_{{ .version }}:
    identifier: transfersshkey-aws.kubeform.com
    name: TransferSSHKey
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## TransferSSHKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `TransferSSHKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[TransferSSHKeySpec](#transfersshkeyspec)***||
| `status` | ***[TransferSSHKeyStatus](#transfersshkeystatus)***||
## Phase(`string` alias)

Appears on:[TransferSSHKeyStatus](#transfersshkeystatus)

## TransferSSHKeySpec

Appears on:[TransferSSHKey](#transfersshkey), [TransferSSHKeyStatus](#transfersshkeystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `body` | ***string***||
| `serverID` | ***string***||
| `userName` | ***string***||
## TransferSSHKeyStatus

Appears on:[TransferSSHKey](#transfersshkey)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[TransferSSHKeySpec](#transfersshkeyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
