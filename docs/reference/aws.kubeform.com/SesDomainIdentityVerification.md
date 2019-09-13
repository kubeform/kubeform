---
title: SesDomainIdentityVerification
menu:
  docs_{{ .version }}:
    identifier: sesdomainidentityverification-aws.kubeform.com
    name: SesDomainIdentityVerification
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SesDomainIdentityVerification
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SesDomainIdentityVerification` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SesDomainIdentityVerificationSpec](#SesDomainIdentityVerificationSpec)***||
| `status` | ***[SesDomainIdentityVerificationStatus](#SesDomainIdentityVerificationStatus)***||
## SesDomainIdentityVerificationSpec
##### (Appears on:[SesDomainIdentityVerification](#SesDomainIdentityVerification), [SesDomainIdentityVerificationStatus](#SesDomainIdentityVerificationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `domain` | ***string***||
## SesDomainIdentityVerificationStatus
##### (Appears on:[SesDomainIdentityVerification](#SesDomainIdentityVerification))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SesDomainIdentityVerificationSpec](#SesDomainIdentityVerificationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
