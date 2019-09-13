---
title: GlueJob
menu:
  docs_{{ .version }}:
    identifier: gluejob-aws.kubeform.com
    name: GlueJob
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GlueJob
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GlueJob` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GlueJobSpec](#GlueJobSpec)***||
| `status` | ***[GlueJobStatus](#GlueJobStatus)***||
## GlueJobSpec
##### (Appears on:[GlueJob](#GlueJob), [GlueJobStatus](#GlueJobStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allocatedCapacity` | ***int***| ***(Optional)*** Deprecated|
| `command` | ***[[]GlueJobSpecCommand](#GlueJobSpecCommand)***||
| `connections` | ***[]string***| ***(Optional)*** |
| `defaultArguments` | ***map[string]string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `executionProperty` | ***[[]GlueJobSpecExecutionProperty](#GlueJobSpecExecutionProperty)***| ***(Optional)*** |
| `maxCapacity` | ***encoding/json.Number***| ***(Optional)*** |
| `maxRetries` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `roleArn` | ***string***||
| `securityConfiguration` | ***string***| ***(Optional)*** |
| `timeout` | ***int***| ***(Optional)*** |
## GlueJobSpecCommand
##### (Appears on:[GlueJobSpec](#GlueJobSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
| `scriptLocation` | ***string***||
## GlueJobSpecExecutionProperty
##### (Appears on:[GlueJobSpec](#GlueJobSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxConcurrentRuns` | ***int***| ***(Optional)*** |
## GlueJobStatus
##### (Appears on:[GlueJob](#GlueJob))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlueJobSpec](#GlueJobSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
