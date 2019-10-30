---
title: KmsKeyRingIamBinding
menu:
  docs_{{ .version }}:
    identifier: kmskeyringiambinding-google.kubeform.com
    name: KmsKeyRingIamBinding
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KmsKeyRingIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `KmsKeyRingIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KmsKeyRingIamBindingSpec](#kmskeyringiambindingspec)***||
| `status` | ***[KmsKeyRingIamBindingStatus](#kmskeyringiambindingstatus)***||
## KmsKeyRingIamBindingSpec

Appears on:[KmsKeyRingIamBinding](#kmskeyringiambinding), [KmsKeyRingIamBindingStatus](#kmskeyringiambindingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `keyRingID` | ***string***||
| `members` | ***[]string***||
| `role` | ***string***||
## KmsKeyRingIamBindingStatus

Appears on:[KmsKeyRingIamBinding](#kmskeyringiambinding)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KmsKeyRingIamBindingSpec](#kmskeyringiambindingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[KmsKeyRingIamBindingStatus](#kmskeyringiambindingstatus)

---
