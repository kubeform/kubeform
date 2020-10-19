---
title: RuntimeconfigConfigIamPolicy
menu:
  docs_{{ .version }}:
    identifier: runtimeconfigconfigiampolicy-google.kubeform.com
    name: RuntimeconfigConfigIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RuntimeconfigConfigIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `RuntimeconfigConfigIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RuntimeconfigConfigIamPolicySpec](#runtimeconfigconfigiampolicyspec)***||
| `status` | ***[RuntimeconfigConfigIamPolicyStatus](#runtimeconfigconfigiampolicystatus)***||
## Phase(`string` alias)

Appears on:[RuntimeconfigConfigIamPolicyStatus](#runtimeconfigconfigiampolicystatus)

## RuntimeconfigConfigIamPolicySpec

Appears on:[RuntimeconfigConfigIamPolicy](#runtimeconfigconfigiampolicy), [RuntimeconfigConfigIamPolicyStatus](#runtimeconfigconfigiampolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `config` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `policyData` | ***string***||
| `project` | ***string***| ***(Optional)*** |
## RuntimeconfigConfigIamPolicyStatus

Appears on:[RuntimeconfigConfigIamPolicy](#runtimeconfigconfigiampolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RuntimeconfigConfigIamPolicySpec](#runtimeconfigconfigiampolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
