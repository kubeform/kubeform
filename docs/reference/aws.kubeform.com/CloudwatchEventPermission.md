---
title: CloudwatchEventPermission
menu:
  docs_{{ .version }}:
    identifier: cloudwatcheventpermission-aws.kubeform.com
    name: CloudwatchEventPermission
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudwatchEventPermission
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudwatchEventPermission` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudwatchEventPermissionSpec](#CloudwatchEventPermissionSpec)***||
| `status` | ***[CloudwatchEventPermissionStatus](#CloudwatchEventPermissionStatus)***||
## CloudwatchEventPermissionSpec
##### (Appears on:[CloudwatchEventPermission](#CloudwatchEventPermission), [CloudwatchEventPermissionStatus](#CloudwatchEventPermissionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `action` | ***string***| ***(Optional)*** |
| `condition` | ***[[]CloudwatchEventPermissionSpecCondition](#CloudwatchEventPermissionSpecCondition)***| ***(Optional)*** |
| `principal` | ***string***||
| `statementID` | ***string***||
## CloudwatchEventPermissionSpecCondition
##### (Appears on:[CloudwatchEventPermissionSpec](#CloudwatchEventPermissionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***||
| `type` | ***string***||
| `value` | ***string***||
## CloudwatchEventPermissionStatus
##### (Appears on:[CloudwatchEventPermission](#CloudwatchEventPermission))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudwatchEventPermissionSpec](#CloudwatchEventPermissionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
