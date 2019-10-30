---
title: SagemakerEndpoint
menu:
  docs_{{ .version }}:
    identifier: sagemakerendpoint-aws.kubeform.com
    name: SagemakerEndpoint
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SagemakerEndpoint
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SagemakerEndpoint` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SagemakerEndpointSpec](#sagemakerendpointspec)***||
| `status` | ***[SagemakerEndpointStatus](#sagemakerendpointstatus)***||
## Phase(`string` alias)

Appears on:[SagemakerEndpointStatus](#sagemakerendpointstatus)

## SagemakerEndpointSpec

Appears on:[SagemakerEndpoint](#sagemakerendpoint), [SagemakerEndpointStatus](#sagemakerendpointstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `endpointConfigName` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## SagemakerEndpointStatus

Appears on:[SagemakerEndpoint](#sagemakerendpoint)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SagemakerEndpointSpec](#sagemakerendpointspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
