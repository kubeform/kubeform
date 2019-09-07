---
title: Sshkey
menu:
  docs_v0.0.1:
    identifier: sshkey-linode.kubeform.com
    name: Sshkey
    parent: linode.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## Sshkey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `linode.kubeform.com/v1alpha1` |
|    `kind` | string | `Sshkey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SshkeySpec](#SshkeySpec)***||
| `status` | ***[SshkeyStatus](#SshkeyStatus)***||
## SshkeySpec
##### (Appears on:[Sshkey](#Sshkey), [SshkeyStatus](#SshkeyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `created` | ***string***| ***(Optional)*** The date this key was added.|
| `label` | ***string***|The label of the Linode SSH Key.|
| `sshKey` | ***string***|The public SSH Key, which is used to authenticate to the root user of the Linodes you deploy.|
## SshkeyStatus
##### (Appears on:[Sshkey](#Sshkey))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SshkeySpec](#SshkeySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
