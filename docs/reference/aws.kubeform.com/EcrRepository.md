---
title: EcrRepository
menu:
  docs_{{ .version }}:
    identifier: ecrrepository-aws.kubeform.com
    name: EcrRepository
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## EcrRepository
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `EcrRepository` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EcrRepositorySpec](#ecrrepositoryspec)***||
| `status` | ***[EcrRepositoryStatus](#ecrrepositorystatus)***||
## EcrRepositorySpec

Appears on:[EcrRepository](#ecrrepository), [EcrRepositoryStatus](#ecrrepositorystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `imageTagMutability` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `registryID` | ***string***| ***(Optional)*** |
| `repositoryURL` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## EcrRepositoryStatus

Appears on:[EcrRepository](#ecrrepository)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EcrRepositorySpec](#ecrrepositoryspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[EcrRepositoryStatus](#ecrrepositorystatus)

---
