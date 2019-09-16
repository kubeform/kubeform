---
title: AthenaNamedQuery
menu:
  docs_{{ .version }}:
    identifier: athenanamedquery-aws.kubeform.com
    name: AthenaNamedQuery
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AthenaNamedQuery
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AthenaNamedQuery` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AthenaNamedQuerySpec](#athenanamedqueryspec)***||
| `status` | ***[AthenaNamedQueryStatus](#athenanamedquerystatus)***||
## AthenaNamedQuerySpec

Appears on:[AthenaNamedQuery](#athenanamedquery), [AthenaNamedQueryStatus](#athenanamedquerystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `database` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `query` | ***string***||
## AthenaNamedQueryStatus

Appears on:[AthenaNamedQuery](#athenanamedquery)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AthenaNamedQuerySpec](#athenanamedqueryspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---