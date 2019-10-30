---
title: AmiLaunchPermission
menu:
  docs_{{ .version }}:
    identifier: amilaunchpermission-aws.kubeform.com
    name: AmiLaunchPermission
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AmiLaunchPermission
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AmiLaunchPermission` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AmiLaunchPermissionSpec](#amilaunchpermissionspec)***||
| `status` | ***[AmiLaunchPermissionStatus](#amilaunchpermissionstatus)***||
## AmiLaunchPermissionSpec

Appears on:[AmiLaunchPermission](#amilaunchpermission), [AmiLaunchPermissionStatus](#amilaunchpermissionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountID` | ***string***||
| `imageID` | ***string***||
## AmiLaunchPermissionStatus

Appears on:[AmiLaunchPermission](#amilaunchpermission)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AmiLaunchPermissionSpec](#amilaunchpermissionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AmiLaunchPermissionStatus](#amilaunchpermissionstatus)

---
