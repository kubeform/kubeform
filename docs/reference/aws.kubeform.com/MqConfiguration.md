---
title: MqConfiguration
menu:
  docs_{{ .version }}:
    identifier: mqconfiguration-aws.kubeform.com
    name: MqConfiguration
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MqConfiguration
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `MqConfiguration` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MqConfigurationSpec](#mqconfigurationspec)***||
| `status` | ***[MqConfigurationStatus](#mqconfigurationstatus)***||
## MqConfigurationSpec

Appears on:[MqConfiguration](#mqconfiguration), [MqConfigurationStatus](#mqconfigurationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `data` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `engineType` | ***string***||
| `engineVersion` | ***string***||
| `latestRevision` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## MqConfigurationStatus

Appears on:[MqConfiguration](#mqconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MqConfigurationSpec](#mqconfigurationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[MqConfigurationStatus](#mqconfigurationstatus)

---
