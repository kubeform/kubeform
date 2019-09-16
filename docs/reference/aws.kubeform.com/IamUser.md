---
title: IamUser
menu:
  docs_{{ .version }}:
    identifier: iamuser-aws.kubeform.com
    name: IamUser
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IamUser
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamUser` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamUserSpec](#IamUserSpec)***||
| `status` | ***[IamUserStatus](#IamUserStatus)***||
## IamUserSpec

Appears on:[IamUser](#IamUser), [IamUserStatus](#IamUserStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `forceDestroy` | ***bool***| ***(Optional)*** Delete user even if it has non-Terraform-managed IAM access keys, login profile or MFA devices|
| `name` | ***string***||
| `path` | ***string***| ***(Optional)*** |
| `permissionsBoundary` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `uniqueID` | ***string***| ***(Optional)*** |
## IamUserStatus

Appears on:[IamUser](#IamUser)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamUserSpec](#IamUserSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
