## AcmCertificate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AcmCertificate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AcmCertificateSpec](#AcmCertificateSpec)***||
| `status` | ***[AcmCertificateStatus](#AcmCertificateStatus)***||
## AcmCertificateSpec
##### (Appears on:[AcmCertificate](#AcmCertificate), [AcmCertificateStatus](#AcmCertificateStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `arn` | ***string***| ***(Optional)*** |
| `certificateBody` | ***string***| ***(Optional)*** |
| `certificateChain` | ***string***| ***(Optional)*** |
| `domainName` | ***string***| ***(Optional)*** |
| `domainValidationOptions` | ***[[]AcmCertificateSpecDomainValidationOptions](#AcmCertificateSpecDomainValidationOptions)***| ***(Optional)*** |
| `subjectAlternativeNames` | ***[]string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `validationEmails` | ***[]string***| ***(Optional)*** |
| `validationMethod` | ***string***| ***(Optional)*** |
## AcmCertificateSpecDomainValidationOptions
##### (Appears on:[AcmCertificateSpec](#AcmCertificateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `domainName` | ***string***| ***(Optional)*** |
| `resourceRecordName` | ***string***| ***(Optional)*** |
| `resourceRecordType` | ***string***| ***(Optional)*** |
| `resourceRecordValue` | ***string***| ***(Optional)*** |
## AcmCertificateStatus
##### (Appears on:[AcmCertificate](#AcmCertificate))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AcmCertificateSpec](#AcmCertificateSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `private_key` | ***string*** ||
