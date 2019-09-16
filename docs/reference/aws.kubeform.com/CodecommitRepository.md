---
title: CodecommitRepository
menu:
  docs_{{ .version }}:
    identifier: codecommitrepository-aws.kubeform.com
    name: CodecommitRepository
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CodecommitRepository
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CodecommitRepository` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CodecommitRepositorySpec](#codecommitrepositoryspec)***||
| `status` | ***[CodecommitRepositoryStatus](#codecommitrepositorystatus)***||
## CodecommitRepositorySpec

Appears on:[CodecommitRepository](#codecommitrepository), [CodecommitRepositoryStatus](#codecommitrepositorystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `cloneURLHTTP` | ***string***| ***(Optional)*** |
| `cloneURLSSH` | ***string***| ***(Optional)*** |
| `defaultBranch` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `repositoryID` | ***string***| ***(Optional)*** |
| `repositoryName` | ***string***||
## CodecommitRepositoryStatus

Appears on:[CodecommitRepository](#codecommitrepository)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CodecommitRepositorySpec](#codecommitrepositoryspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
