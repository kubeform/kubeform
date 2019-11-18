---
title: IamRole
menu:
  docs_{{ .version }}:
    identifier: iamrole-aws.kubeform.com
    name: IamRole
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IamRole
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamRole` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamRoleSpec](#iamrolespec)***||
| `status` | ***[IamRoleStatus](#iamrolestatus)***||
## IamRoleSpec

Appears on:[IamRole](#iamrole), [IamRoleStatus](#iamrolestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `assumeRolePolicy` | ***string***||
| `createDate` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `forceDetachPolicies` | ***bool***| ***(Optional)*** |
| `maxSessionDuration` | ***int64***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `permissionsBoundary` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `uniqueID` | ***string***| ***(Optional)*** |
## IamRoleStatus

Appears on:[IamRole](#iamrole)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamRoleSpec](#iamrolespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IamRoleStatus](#iamrolestatus)

---
