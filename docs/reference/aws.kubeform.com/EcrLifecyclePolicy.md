---
title: EcrLifecyclePolicy
menu:
  docs_{{ .version }}:
    identifier: ecrlifecyclepolicy-aws.kubeform.com
    name: EcrLifecyclePolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## EcrLifecyclePolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `EcrLifecyclePolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EcrLifecyclePolicySpec](#ecrlifecyclepolicyspec)***||
| `status` | ***[EcrLifecyclePolicyStatus](#ecrlifecyclepolicystatus)***||
## EcrLifecyclePolicySpec

Appears on:[EcrLifecyclePolicy](#ecrlifecyclepolicy), [EcrLifecyclePolicyStatus](#ecrlifecyclepolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `policy` | ***string***||
| `registryID` | ***string***| ***(Optional)*** |
| `repository` | ***string***||
## EcrLifecyclePolicyStatus

Appears on:[EcrLifecyclePolicy](#ecrlifecyclepolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EcrLifecyclePolicySpec](#ecrlifecyclepolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[EcrLifecyclePolicyStatus](#ecrlifecyclepolicystatus)

---
