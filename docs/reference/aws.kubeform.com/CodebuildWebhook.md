---
title: CodebuildWebhook
menu:
  docs_{{ .version }}:
    identifier: codebuildwebhook-aws.kubeform.com
    name: CodebuildWebhook
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CodebuildWebhook
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CodebuildWebhook` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CodebuildWebhookSpec](#CodebuildWebhookSpec)***||
| `status` | ***[CodebuildWebhookStatus](#CodebuildWebhookStatus)***||
## CodebuildWebhookSpec

Appears on:[CodebuildWebhook](#CodebuildWebhook), [CodebuildWebhookStatus](#CodebuildWebhookStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `branchFilter` | ***string***| ***(Optional)*** |
| `payloadURL` | ***string***| ***(Optional)*** |
| `projectName` | ***string***||
| `url` | ***string***| ***(Optional)*** |
## CodebuildWebhookStatus

Appears on:[CodebuildWebhook](#CodebuildWebhook)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CodebuildWebhookSpec](#CodebuildWebhookSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `secret` | ***string*** ||
