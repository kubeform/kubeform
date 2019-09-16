---
title: SsmDocument
menu:
  docs_{{ .version }}:
    identifier: ssmdocument-aws.kubeform.com
    name: SsmDocument
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SsmDocument
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SsmDocument` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SsmDocumentSpec](#SsmDocumentSpec)***||
| `status` | ***[SsmDocumentStatus](#SsmDocumentStatus)***||
## SsmDocumentSpec

Appears on:[SsmDocument](#SsmDocument), [SsmDocumentStatus](#SsmDocumentStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `content` | ***string***||
| `createdDate` | ***string***| ***(Optional)*** |
| `defaultVersion` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `documentFormat` | ***string***| ***(Optional)*** |
| `documentType` | ***string***||
| `hash` | ***string***| ***(Optional)*** |
| `hashType` | ***string***| ***(Optional)*** |
| `latestVersion` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `owner` | ***string***| ***(Optional)*** |
| `parameter` | ***[[]SsmDocumentSpecParameter](#SsmDocumentSpecParameter)***| ***(Optional)*** |
| `permissions` | ***map[string]kubeform.dev/kubeform/apis/aws/v1alpha1.SsmDocumentSpecPermissions***| ***(Optional)*** |
| `platformTypes` | ***[]string***| ***(Optional)*** |
| `schemaVersion` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## SsmDocumentSpecParameter

Appears on:[SsmDocumentSpec](#SsmDocumentSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultValue` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## SsmDocumentStatus

Appears on:[SsmDocument](#SsmDocument)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SsmDocumentSpec](#SsmDocumentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
