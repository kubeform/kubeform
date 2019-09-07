---
title: ServiceAccountKey
menu:
  docs_v0.0.1:
    identifier: serviceaccountkey-google.kubeform.com
    name: ServiceAccountKey
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ServiceAccountKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ServiceAccountKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServiceAccountKeySpec](#ServiceAccountKeySpec)***||
| `status` | ***[ServiceAccountKeyStatus](#ServiceAccountKeyStatus)***||
## ServiceAccountKeySpec
##### (Appears on:[ServiceAccountKey](#ServiceAccountKey), [ServiceAccountKeyStatus](#ServiceAccountKeyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `keyAlgorithm` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `pgpKey` | ***string***| ***(Optional)*** |
| `privateKeyEncrypted` | ***string***| ***(Optional)*** |
| `privateKeyFingerprint` | ***string***| ***(Optional)*** |
| `privateKeyType` | ***string***| ***(Optional)*** |
| `publicKey` | ***string***| ***(Optional)*** |
| `publicKeyType` | ***string***| ***(Optional)*** |
| `serviceAccountID` | ***string***||
| `validAfter` | ***string***| ***(Optional)*** |
| `validBefore` | ***string***| ***(Optional)*** |
## ServiceAccountKeyStatus
##### (Appears on:[ServiceAccountKey](#ServiceAccountKey))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServiceAccountKeySpec](#ServiceAccountKeySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `private_key` | ***string*** ||
