---
title: CloudformationStackSet
menu:
  docs_{{ .version }}:
    identifier: cloudformationstackset-aws.kubeform.com
    name: CloudformationStackSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudformationStackSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudformationStackSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudformationStackSetSpec](#cloudformationstacksetspec)***||
| `status` | ***[CloudformationStackSetStatus](#cloudformationstacksetstatus)***||
## CloudformationStackSetSpec

Appears on:[CloudformationStackSet](#cloudformationstackset), [CloudformationStackSetStatus](#cloudformationstacksetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `administrationRoleArn` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `capabilities` | ***[]string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `executionRoleName` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `parameters` | ***map[string]string***| ***(Optional)*** |
| `stackSetID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `templateBody` | ***string***| ***(Optional)*** |
| `templateURL` | ***string***| ***(Optional)*** |
## CloudformationStackSetStatus

Appears on:[CloudformationStackSet](#cloudformationstackset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudformationStackSetSpec](#cloudformationstacksetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
