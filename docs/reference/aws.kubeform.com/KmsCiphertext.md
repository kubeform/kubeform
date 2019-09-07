---
title: KmsCiphertext
menu:
  docs_v0.0.1:
    identifier: kmsciphertext-aws.kubeform.com
    name: KmsCiphertext
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## KmsCiphertext
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `KmsCiphertext` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KmsCiphertextSpec](#KmsCiphertextSpec)***||
| `status` | ***[KmsCiphertextStatus](#KmsCiphertextStatus)***||
## KmsCiphertextSpec
##### (Appears on:[KmsCiphertext](#KmsCiphertext), [KmsCiphertextStatus](#KmsCiphertextStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `ciphertextBlob` | ***string***| ***(Optional)*** |
| `context` | ***map[string]string***| ***(Optional)*** |
| `keyID` | ***string***||
## KmsCiphertextStatus
##### (Appears on:[KmsCiphertext](#KmsCiphertext))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KmsCiphertextSpec](#KmsCiphertextSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `plaintext` | ***string*** ||
