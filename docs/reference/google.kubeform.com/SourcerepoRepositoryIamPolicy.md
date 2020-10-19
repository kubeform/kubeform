---
title: SourcerepoRepositoryIamPolicy
menu:
  docs_{{ .version }}:
    identifier: sourcereporepositoryiampolicy-google.kubeform.com
    name: SourcerepoRepositoryIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SourcerepoRepositoryIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `SourcerepoRepositoryIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SourcerepoRepositoryIamPolicySpec](#sourcereporepositoryiampolicyspec)***||
| `status` | ***[SourcerepoRepositoryIamPolicyStatus](#sourcereporepositoryiampolicystatus)***||
## Phase(`string` alias)

Appears on:[SourcerepoRepositoryIamPolicyStatus](#sourcereporepositoryiampolicystatus)

## SourcerepoRepositoryIamPolicySpec

Appears on:[SourcerepoRepositoryIamPolicy](#sourcereporepositoryiampolicy), [SourcerepoRepositoryIamPolicyStatus](#sourcereporepositoryiampolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `policyData` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `repository` | ***string***||
## SourcerepoRepositoryIamPolicyStatus

Appears on:[SourcerepoRepositoryIamPolicy](#sourcereporepositoryiampolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SourcerepoRepositoryIamPolicySpec](#sourcereporepositoryiampolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
