---
title: IamServiceLinkedRole
menu:
  docs_{{ .version }}:
    identifier: iamservicelinkedrole-aws.kubeform.com
    name: IamServiceLinkedRole
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IamServiceLinkedRole
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamServiceLinkedRole` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamServiceLinkedRoleSpec](#iamservicelinkedrolespec)***||
| `status` | ***[IamServiceLinkedRoleStatus](#iamservicelinkedrolestatus)***||
## IamServiceLinkedRoleSpec

Appears on:[IamServiceLinkedRole](#iamservicelinkedrole), [IamServiceLinkedRoleStatus](#iamservicelinkedrolestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `awsServiceName` | ***string***||
| `createDate` | ***string***| ***(Optional)*** |
| `customSuffix` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `uniqueID` | ***string***| ***(Optional)*** |
## IamServiceLinkedRoleStatus

Appears on:[IamServiceLinkedRole](#iamservicelinkedrole)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamServiceLinkedRoleSpec](#iamservicelinkedrolespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
