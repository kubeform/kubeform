---
title: OpsworksPermission
menu:
  docs_v0.0.1:
    identifier: opsworkspermission-aws.kubeform.com
    name: OpsworksPermission
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## OpsworksPermission
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OpsworksPermission` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OpsworksPermissionSpec](#OpsworksPermissionSpec)***||
| `status` | ***[OpsworksPermissionStatus](#OpsworksPermissionStatus)***||
## OpsworksPermissionSpec
##### (Appears on:[OpsworksPermission](#OpsworksPermission), [OpsworksPermissionStatus](#OpsworksPermissionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allowSSH` | ***bool***| ***(Optional)*** |
| `allowSudo` | ***bool***| ***(Optional)*** |
| `level` | ***string***| ***(Optional)*** |
| `stackID` | ***string***| ***(Optional)*** |
| `userArn` | ***string***||
## OpsworksPermissionStatus
##### (Appears on:[OpsworksPermission](#OpsworksPermission))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OpsworksPermissionSpec](#OpsworksPermissionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
