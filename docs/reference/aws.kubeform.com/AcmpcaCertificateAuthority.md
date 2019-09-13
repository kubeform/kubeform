---
title: AcmpcaCertificateAuthority
menu:
  docs_{{ .version }}:
    identifier: acmpcacertificateauthority-aws.kubeform.com
    name: AcmpcaCertificateAuthority
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AcmpcaCertificateAuthority
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AcmpcaCertificateAuthority` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AcmpcaCertificateAuthoritySpec](#AcmpcaCertificateAuthoritySpec)***||
| `status` | ***[AcmpcaCertificateAuthorityStatus](#AcmpcaCertificateAuthorityStatus)***||
## AcmpcaCertificateAuthoritySpec
##### (Appears on:[AcmpcaCertificateAuthority](#AcmpcaCertificateAuthority), [AcmpcaCertificateAuthorityStatus](#AcmpcaCertificateAuthorityStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `certificate` | ***string***| ***(Optional)*** |
| `certificateAuthorityConfiguration` | ***[[]AcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration](#AcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration)***||
| `certificateChain` | ***string***| ***(Optional)*** |
| `certificateSigningRequest` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `notAfter` | ***string***| ***(Optional)*** |
| `notBefore` | ***string***| ***(Optional)*** |
| `permanentDeletionTimeInDays` | ***int***| ***(Optional)*** |
| `revocationConfiguration` | ***[[]AcmpcaCertificateAuthoritySpecRevocationConfiguration](#AcmpcaCertificateAuthoritySpecRevocationConfiguration)***| ***(Optional)*** |
| `serial` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## AcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration
##### (Appears on:[AcmpcaCertificateAuthoritySpec](#AcmpcaCertificateAuthoritySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `keyAlgorithm` | ***string***||
| `signingAlgorithm` | ***string***||
| `subject` | ***[[]AcmpcaCertificateAuthoritySpecCertificateAuthorityConfigurationSubject](#AcmpcaCertificateAuthoritySpecCertificateAuthorityConfigurationSubject)***||
## AcmpcaCertificateAuthoritySpecCertificateAuthorityConfigurationSubject
##### (Appears on:[AcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration](#AcmpcaCertificateAuthoritySpecCertificateAuthorityConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `commonName` | ***string***| ***(Optional)*** |
| `country` | ***string***| ***(Optional)*** |
| `distinguishedNameQualifier` | ***string***| ***(Optional)*** |
| `generationQualifier` | ***string***| ***(Optional)*** |
| `givenName` | ***string***| ***(Optional)*** |
| `initials` | ***string***| ***(Optional)*** |
| `locality` | ***string***| ***(Optional)*** |
| `organization` | ***string***| ***(Optional)*** |
| `organizationalUnit` | ***string***| ***(Optional)*** |
| `pseudonym` | ***string***| ***(Optional)*** |
| `state` | ***string***| ***(Optional)*** |
| `surname` | ***string***| ***(Optional)*** |
| `title` | ***string***| ***(Optional)*** |
## AcmpcaCertificateAuthoritySpecRevocationConfiguration
##### (Appears on:[AcmpcaCertificateAuthoritySpec](#AcmpcaCertificateAuthoritySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `crlConfiguration` | ***[[]AcmpcaCertificateAuthoritySpecRevocationConfigurationCrlConfiguration](#AcmpcaCertificateAuthoritySpecRevocationConfigurationCrlConfiguration)***| ***(Optional)*** |
## AcmpcaCertificateAuthoritySpecRevocationConfigurationCrlConfiguration
##### (Appears on:[AcmpcaCertificateAuthoritySpecRevocationConfiguration](#AcmpcaCertificateAuthoritySpecRevocationConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `customCname` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `expirationInDays` | ***int***||
| `s3BucketName` | ***string***| ***(Optional)*** |
## AcmpcaCertificateAuthorityStatus
##### (Appears on:[AcmpcaCertificateAuthority](#AcmpcaCertificateAuthority))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AcmpcaCertificateAuthoritySpec](#AcmpcaCertificateAuthoritySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
