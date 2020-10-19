---
title: OpsworksPermission
menu:
  docs_{{ .version }}:
    identifier: opsworkspermission-aws.kubeform.com
    name: OpsworksPermission
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## OpsworksPermission
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OpsworksPermission` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OpsworksPermissionSpec](#opsworkspermissionspec)***||
| `status` | ***[OpsworksPermissionStatus](#opsworkspermissionstatus)***||
## OpsworksPermissionSpec

Appears on:[OpsworksPermission](#opsworkspermission), [OpsworksPermissionStatus](#opsworkspermissionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allowSSH` | ***bool***| ***(Optional)*** |
| `allowSudo` | ***bool***| ***(Optional)*** |
| `level` | ***string***| ***(Optional)*** |
| `stackID` | ***string***| ***(Optional)*** |
| `userArn` | ***string***||
## OpsworksPermissionStatus

Appears on:[OpsworksPermission](#opsworkspermission)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OpsworksPermissionSpec](#opsworkspermissionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[OpsworksPermissionStatus](#opsworkspermissionstatus)

---
