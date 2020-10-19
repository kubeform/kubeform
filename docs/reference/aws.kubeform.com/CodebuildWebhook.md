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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CodebuildWebhookSpec](#codebuildwebhookspec)***||
| `status` | ***[CodebuildWebhookStatus](#codebuildwebhookstatus)***||
## CodebuildWebhookSpec

Appears on:[CodebuildWebhook](#codebuildwebhook), [CodebuildWebhookStatus](#codebuildwebhookstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `branchFilter` | ***string***| ***(Optional)*** |
| `filterGroup` | ***[[]CodebuildWebhookSpecFilterGroup](#codebuildwebhookspecfiltergroup)***| ***(Optional)*** |
| `payloadURL` | ***string***| ***(Optional)*** |
| `projectName` | ***string***||
| `url` | ***string***| ***(Optional)*** |
## CodebuildWebhookSpecFilterGroup

Appears on:[CodebuildWebhookSpec](#codebuildwebhookspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `filter` | ***[[]CodebuildWebhookSpecFilterGroupFilter](#codebuildwebhookspecfiltergroupfilter)***| ***(Optional)*** |
## CodebuildWebhookSpecFilterGroupFilter

Appears on:[CodebuildWebhookSpecFilterGroup](#codebuildwebhookspecfiltergroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `excludeMatchedPattern` | ***bool***| ***(Optional)*** |
| `pattern` | ***string***||
| `type` | ***string***||
## CodebuildWebhookStatus

Appears on:[CodebuildWebhook](#codebuildwebhook)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CodebuildWebhookSpec](#codebuildwebhookspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CodebuildWebhookStatus](#codebuildwebhookstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `secret` | ***string*** ||
