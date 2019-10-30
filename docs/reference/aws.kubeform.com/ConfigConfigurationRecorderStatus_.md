---
title: ConfigConfigurationRecorderStatus 
menu:
  docs_{{ .version }}:
    identifier: configconfigurationrecorderstatus--aws.kubeform.com
    name: ConfigConfigurationRecorderStatus 
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ConfigConfigurationRecorderStatus_
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ConfigConfigurationRecorderStatus_` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ConfigConfigurationRecorderStatus_Spec](#configconfigurationrecorderstatus_spec)***||
| `status` | ***[ConfigConfigurationRecorderStatus_Status](#configconfigurationrecorderstatus_status)***||
## ConfigConfigurationRecorderStatus_Spec

Appears on:[ConfigConfigurationRecorderStatus_](#configconfigurationrecorderstatus_), [ConfigConfigurationRecorderStatus_Status](#configconfigurationrecorderstatus_status)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `isEnabled` | ***bool***||
| `name` | ***string***||
## ConfigConfigurationRecorderStatus_Status

Appears on:[ConfigConfigurationRecorderStatus_](#configconfigurationrecorderstatus_)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ConfigConfigurationRecorderStatus_Spec](#configconfigurationrecorderstatus_spec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ConfigConfigurationRecorderStatus_Status](#configconfigurationrecorderstatus_status)

---
