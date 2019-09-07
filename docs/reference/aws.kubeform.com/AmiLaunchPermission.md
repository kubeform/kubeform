---
title: AmiLaunchPermission
menu:
  docs_v0.0.1:
    identifier: amilaunchpermission-aws.kubeform.com
    name: AmiLaunchPermission
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## AmiLaunchPermission
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AmiLaunchPermission` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AmiLaunchPermissionSpec](#AmiLaunchPermissionSpec)***||
| `status` | ***[AmiLaunchPermissionStatus](#AmiLaunchPermissionStatus)***||
## AmiLaunchPermissionSpec
##### (Appears on:[AmiLaunchPermission](#AmiLaunchPermission), [AmiLaunchPermissionStatus](#AmiLaunchPermissionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountID` | ***string***||
| `imageID` | ***string***||
## AmiLaunchPermissionStatus
##### (Appears on:[AmiLaunchPermission](#AmiLaunchPermission))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AmiLaunchPermissionSpec](#AmiLaunchPermissionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
