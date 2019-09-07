---
title: KmsKey
menu:
  docs_v0.0.1:
    identifier: kmskey-aws.kubeform.com
    name: KmsKey
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## KmsKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `KmsKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KmsKeySpec](#KmsKeySpec)***||
| `status` | ***[KmsKeyStatus](#KmsKeyStatus)***||
## KmsKeySpec
##### (Appears on:[KmsKey](#KmsKey), [KmsKeyStatus](#KmsKeyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `deletionWindowInDays` | ***int***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `enableKeyRotation` | ***bool***| ***(Optional)*** |
| `isEnabled` | ***bool***| ***(Optional)*** |
| `keyID` | ***string***| ***(Optional)*** |
| `keyUsage` | ***string***| ***(Optional)*** |
| `policy` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## KmsKeyStatus
##### (Appears on:[KmsKey](#KmsKey))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KmsKeySpec](#KmsKeySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
