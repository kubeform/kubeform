---
title: KmsCiphertext
menu:
  docs_{{ .version }}:
    identifier: kmsciphertext-aws.kubeform.com
    name: KmsCiphertext
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KmsCiphertext
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `KmsCiphertext` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KmsCiphertextSpec](#kmsciphertextspec)***||
| `status` | ***[KmsCiphertextStatus](#kmsciphertextstatus)***||
## KmsCiphertextSpec

Appears on:[KmsCiphertext](#kmsciphertext), [KmsCiphertextStatus](#kmsciphertextstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `ciphertextBlob` | ***string***| ***(Optional)*** |
| `context` | ***map[string]string***| ***(Optional)*** |
| `keyID` | ***string***||
## KmsCiphertextStatus

Appears on:[KmsCiphertext](#kmsciphertext)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KmsCiphertextSpec](#kmsciphertextspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `plaintext` | ***string*** ||
