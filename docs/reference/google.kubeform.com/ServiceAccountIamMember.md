---
title: ServiceAccountIamMember
menu:
  docs_{{ .version }}:
    identifier: serviceaccountiammember-google.kubeform.com
    name: ServiceAccountIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ServiceAccountIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ServiceAccountIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServiceAccountIamMemberSpec](#serviceaccountiammemberspec)***||
| `status` | ***[ServiceAccountIamMemberStatus](#serviceaccountiammemberstatus)***||
## Phase(`string` alias)

Appears on:[ServiceAccountIamMemberStatus](#serviceaccountiammemberstatus)

## ServiceAccountIamMemberSpec

Appears on:[ServiceAccountIamMember](#serviceaccountiammember), [ServiceAccountIamMemberStatus](#serviceaccountiammemberstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `member` | ***string***||
| `role` | ***string***||
| `serviceAccountID` | ***string***||
## ServiceAccountIamMemberStatus

Appears on:[ServiceAccountIamMember](#serviceaccountiammember)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServiceAccountIamMemberSpec](#serviceaccountiammemberspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
