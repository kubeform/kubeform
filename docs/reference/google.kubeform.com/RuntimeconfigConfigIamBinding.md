---
title: RuntimeconfigConfigIamBinding
menu:
  docs_{{ .version }}:
    identifier: runtimeconfigconfigiambinding-google.kubeform.com
    name: RuntimeconfigConfigIamBinding
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RuntimeconfigConfigIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `RuntimeconfigConfigIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RuntimeconfigConfigIamBindingSpec](#runtimeconfigconfigiambindingspec)***||
| `status` | ***[RuntimeconfigConfigIamBindingStatus](#runtimeconfigconfigiambindingstatus)***||
## Phase(`string` alias)

Appears on:[RuntimeconfigConfigIamBindingStatus](#runtimeconfigconfigiambindingstatus)

## RuntimeconfigConfigIamBindingSpec

Appears on:[RuntimeconfigConfigIamBinding](#runtimeconfigconfigiambinding), [RuntimeconfigConfigIamBindingStatus](#runtimeconfigconfigiambindingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `config` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `members` | ***[]string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
## RuntimeconfigConfigIamBindingStatus

Appears on:[RuntimeconfigConfigIamBinding](#runtimeconfigconfigiambinding)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RuntimeconfigConfigIamBindingSpec](#runtimeconfigconfigiambindingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
