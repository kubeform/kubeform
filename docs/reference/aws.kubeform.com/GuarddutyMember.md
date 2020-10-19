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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GuarddutyMemberSpec](#guarddutymemberspec)***||
| `status` | ***[GuarddutyMemberStatus](#guarddutymemberstatus)***||
## GuarddutyMemberSpec

Appears on:[GuarddutyMember](#guarddutymember), [GuarddutyMemberStatus](#guarddutymemberstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountID` | ***string***||
| `detectorID` | ***string***||
| `disableEmailNotification` | ***bool***| ***(Optional)*** |
| `email` | ***string***||
| `invitationMessage` | ***string***| ***(Optional)*** |
| `invite` | ***bool***| ***(Optional)*** |
| `relationshipStatus` | ***string***| ***(Optional)*** |
## GuarddutyMemberStatus

Appears on:[GuarddutyMember](#guarddutymember)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GuarddutyMemberSpec](#guarddutymemberspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[GuarddutyMemberStatus](#guarddutymemberstatus)

---
