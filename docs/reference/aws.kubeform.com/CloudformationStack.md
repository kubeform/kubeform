---
title: CloudformationStack
menu:
  docs_{{ .version }}:
    identifier: cloudformationstack-aws.kubeform.com
    name: CloudformationStack
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudformationStack
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudformationStack` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudformationStackSpec](#cloudformationstackspec)***||
| `status` | ***[CloudformationStackStatus](#cloudformationstackstatus)***||
## CloudformationStackSpec

Appears on:[CloudformationStack](#cloudformationstack), [CloudformationStackStatus](#cloudformationstackstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `capabilities` | ***[]string***| ***(Optional)*** |
| `disableRollback` | ***bool***| ***(Optional)*** |
| `iamRoleArn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `notificationArns` | ***[]string***| ***(Optional)*** |
| `onFailure` | ***string***| ***(Optional)*** |
| `outputs` | ***map[string]string***| ***(Optional)*** |
| `parameters` | ***map[string]string***| ***(Optional)*** |
| `policyBody` | ***string***| ***(Optional)*** |
| `policyURL` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `templateBody` | ***string***| ***(Optional)*** |
| `templateURL` | ***string***| ***(Optional)*** |
| `timeoutInMinutes` | ***int***| ***(Optional)*** |
## CloudformationStackStatus

Appears on:[CloudformationStack](#cloudformationstack)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudformationStackSpec](#cloudformationstackspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
