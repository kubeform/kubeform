---
title: SourcerepoRepositoryIamMember
menu:
  docs_{{ .version }}:
    identifier: sourcereporepositoryiammember-google.kubeform.com
    name: SourcerepoRepositoryIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SourcerepoRepositoryIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `SourcerepoRepositoryIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SourcerepoRepositoryIamMemberSpec](#sourcereporepositoryiammemberspec)***||
| `status` | ***[SourcerepoRepositoryIamMemberStatus](#sourcereporepositoryiammemberstatus)***||
## Phase(`string` alias)

Appears on:[SourcerepoRepositoryIamMemberStatus](#sourcereporepositoryiammemberstatus)

## SourcerepoRepositoryIamMemberSpec

Appears on:[SourcerepoRepositoryIamMember](#sourcereporepositoryiammember), [SourcerepoRepositoryIamMemberStatus](#sourcereporepositoryiammemberstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `member` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `repository` | ***string***||
| `role` | ***string***||
## SourcerepoRepositoryIamMemberStatus

Appears on:[SourcerepoRepositoryIamMember](#sourcereporepositoryiammember)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SourcerepoRepositoryIamMemberSpec](#sourcereporepositoryiammemberspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
