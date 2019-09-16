---
title: CognitoUserGroup
menu:
  docs_{{ .version }}:
    identifier: cognitousergroup-aws.kubeform.com
    name: CognitoUserGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CognitoUserGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CognitoUserGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CognitoUserGroupSpec](#CognitoUserGroupSpec)***||
| `status` | ***[CognitoUserGroupStatus](#CognitoUserGroupStatus)***||
## CognitoUserGroupSpec

Appears on:[CognitoUserGroup](#CognitoUserGroup), [CognitoUserGroupStatus](#CognitoUserGroupStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `precedence` | ***int***| ***(Optional)*** |
| `roleArn` | ***string***| ***(Optional)*** |
| `userPoolID` | ***string***||
## CognitoUserGroupStatus

Appears on:[CognitoUserGroup](#CognitoUserGroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CognitoUserGroupSpec](#CognitoUserGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
