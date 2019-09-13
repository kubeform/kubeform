---
title: MskConfiguration
menu:
  docs_{{ .version }}:
    identifier: mskconfiguration-aws.kubeform.com
    name: MskConfiguration
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MskConfiguration
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `MskConfiguration` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MskConfigurationSpec](#MskConfigurationSpec)***||
| `status` | ***[MskConfigurationStatus](#MskConfigurationStatus)***||
## MskConfigurationSpec
##### (Appears on:[MskConfiguration](#MskConfiguration), [MskConfigurationStatus](#MskConfigurationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `kafkaVersions` | ***[]string***||
| `latestRevision` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `serverProperties` | ***string***||
## MskConfigurationStatus
##### (Appears on:[MskConfiguration](#MskConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MskConfigurationSpec](#MskConfigurationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
