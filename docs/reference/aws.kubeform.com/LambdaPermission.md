---
title: LambdaPermission
menu:
  docs_{{ .version }}:
    identifier: lambdapermission-aws.kubeform.com
    name: LambdaPermission
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LambdaPermission
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LambdaPermission` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LambdaPermissionSpec](#lambdapermissionspec)***||
| `status` | ***[LambdaPermissionStatus](#lambdapermissionstatus)***||
## LambdaPermissionSpec

Appears on:[LambdaPermission](#lambdapermission), [LambdaPermissionStatus](#lambdapermissionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `action` | ***string***||
| `eventSourceToken` | ***string***| ***(Optional)*** |
| `functionName` | ***string***||
| `principal` | ***string***||
| `qualifier` | ***string***| ***(Optional)*** |
| `sourceAccount` | ***string***| ***(Optional)*** |
| `sourceArn` | ***string***| ***(Optional)*** |
| `statementID` | ***string***| ***(Optional)*** |
| `statementIDPrefix` | ***string***| ***(Optional)*** |
## LambdaPermissionStatus

Appears on:[LambdaPermission](#lambdapermission)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LambdaPermissionSpec](#lambdapermissionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---