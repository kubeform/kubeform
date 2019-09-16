---
title: GuarddutyMember
menu:
  docs_{{ .version }}:
    identifier: guarddutymember-aws.kubeform.com
    name: GuarddutyMember
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GuarddutyMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GuarddutyMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GuarddutyMemberSpec](#GuarddutyMemberSpec)***||
| `status` | ***[GuarddutyMemberStatus](#GuarddutyMemberStatus)***||
## GuarddutyMemberSpec

Appears on:[GuarddutyMember](#GuarddutyMember), [GuarddutyMemberStatus](#GuarddutyMemberStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountID` | ***string***||
| `detectorID` | ***string***||
| `disableEmailNotification` | ***bool***| ***(Optional)*** |
| `email` | ***string***||
| `invitationMessage` | ***string***| ***(Optional)*** |
| `invite` | ***bool***| ***(Optional)*** |
| `relationshipStatus` | ***string***| ***(Optional)*** |
## GuarddutyMemberStatus

Appears on:[GuarddutyMember](#GuarddutyMember)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GuarddutyMemberSpec](#GuarddutyMemberSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
