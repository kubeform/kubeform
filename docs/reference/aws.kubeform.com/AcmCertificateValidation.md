---
title: AcmCertificateValidation
menu:
  docs_v0.0.1:
    identifier: acmcertificatevalidation-aws.kubeform.com
    name: AcmCertificateValidation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## AcmCertificateValidation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AcmCertificateValidation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AcmCertificateValidationSpec](#AcmCertificateValidationSpec)***||
| `status` | ***[AcmCertificateValidationStatus](#AcmCertificateValidationStatus)***||
## AcmCertificateValidationSpec
##### (Appears on:[AcmCertificateValidation](#AcmCertificateValidation), [AcmCertificateValidationStatus](#AcmCertificateValidationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `certificateArn` | ***string***||
| `validationRecordFqdns` | ***[]string***| ***(Optional)*** |
## AcmCertificateValidationStatus
##### (Appears on:[AcmCertificateValidation](#AcmCertificateValidation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AcmCertificateValidationSpec](#AcmCertificateValidationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
