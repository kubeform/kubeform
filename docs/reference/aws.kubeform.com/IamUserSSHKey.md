---
title: IamUserSSHKey
menu:
  docs_{{ .version }}:
    identifier: iamusersshkey-aws.kubeform.com
    name: IamUserSSHKey
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IamUserSSHKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamUserSSHKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamUserSSHKeySpec](#IamUserSSHKeySpec)***||
| `status` | ***[IamUserSSHKeyStatus](#IamUserSSHKeyStatus)***||
## IamUserSSHKeySpec

Appears on:[IamUserSSHKey](#IamUserSSHKey), [IamUserSSHKeyStatus](#IamUserSSHKeyStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `encoding` | ***string***||
| `fingerprint` | ***string***| ***(Optional)*** |
| `publicKey` | ***string***||
| `sshPublicKeyID` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `username` | ***string***||
## IamUserSSHKeyStatus

Appears on:[IamUserSSHKey](#IamUserSSHKey)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamUserSSHKeySpec](#IamUserSSHKeySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
