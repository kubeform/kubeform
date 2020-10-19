---
title: BatchJobDefinition
menu:
  docs_{{ .version }}:
    identifier: batchjobdefinition-aws.kubeform.com
    name: BatchJobDefinition
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BatchJobDefinition
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `BatchJobDefinition` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BatchJobDefinitionSpec](#batchjobdefinitionspec)***||
| `status` | ***[BatchJobDefinitionStatus](#batchjobdefinitionstatus)***||
## BatchJobDefinitionSpec

Appears on:[BatchJobDefinition](#batchjobdefinition), [BatchJobDefinitionStatus](#batchjobdefinitionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `containerProperties` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `parameters` | ***map[string]string***| ***(Optional)*** |
| `retryStrategy` | ***[[]BatchJobDefinitionSpecRetryStrategy](#batchjobdefinitionspecretrystrategy)***| ***(Optional)*** |
| `revision` | ***int64***| ***(Optional)*** |
| `timeout` | ***[[]BatchJobDefinitionSpecTimeout](#batchjobdefinitionspectimeout)***| ***(Optional)*** |
| `type` | ***string***||
## BatchJobDefinitionSpecRetryStrategy

Appears on:[BatchJobDefinitionSpec](#batchjobdefinitionspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `attempts` | ***int64***| ***(Optional)*** |
## BatchJobDefinitionSpecTimeout

Appears on:[BatchJobDefinitionSpec](#batchjobdefinitionspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `attemptDurationSeconds` | ***int64***| ***(Optional)*** |
## BatchJobDefinitionStatus

Appears on:[BatchJobDefinition](#batchjobdefinition)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BatchJobDefinitionSpec](#batchjobdefinitionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BatchJobDefinitionStatus](#batchjobdefinitionstatus)

---
