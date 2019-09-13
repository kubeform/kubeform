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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SshKeySpec](#SshKeySpec)***||
| `status` | ***[SshKeyStatus](#SshKeyStatus)***||
## SshKeySpec
##### (Appears on:[SshKey](#SshKey), [SshKeyStatus](#SshKeyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `fingerprint` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `publicKey` | ***string***||
## SshKeyStatus
##### (Appears on:[SshKey](#SshKey))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SshKeySpec](#SshKeySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
