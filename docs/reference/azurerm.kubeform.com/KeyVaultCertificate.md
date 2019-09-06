## KeyVaultCertificate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `KeyVaultCertificate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KeyVaultCertificateSpec](#KeyVaultCertificateSpec)***||
| `status` | ***[KeyVaultCertificateStatus](#KeyVaultCertificateStatus)***||
## KeyVaultCertificateSpec
##### (Appears on:[KeyVaultCertificate](#KeyVaultCertificate), [KeyVaultCertificateStatus](#KeyVaultCertificateStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `certificate` | ***[[]KeyVaultCertificateSpecCertificate](#KeyVaultCertificateSpecCertificate)***| ***(Optional)*** |
| `certificateData` | ***string***| ***(Optional)*** |
| `certificatePolicy` | ***[[]KeyVaultCertificateSpecCertificatePolicy](#KeyVaultCertificateSpecCertificatePolicy)***||
| `keyVaultID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `secretID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `thumbprint` | ***string***| ***(Optional)*** |
| `vaultURI` | ***string***| ***(Optional)*** Deprecated|
| `version` | ***string***| ***(Optional)*** |
## KeyVaultCertificateSpecCertificate
##### (Appears on:[KeyVaultCertificateSpec](#KeyVaultCertificateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
## KeyVaultCertificateSpecCertificatePolicy
##### (Appears on:[KeyVaultCertificateSpec](#KeyVaultCertificateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `issuerParameters` | ***[[]KeyVaultCertificateSpecCertificatePolicyIssuerParameters](#KeyVaultCertificateSpecCertificatePolicyIssuerParameters)***||
| `keyProperties` | ***[[]KeyVaultCertificateSpecCertificatePolicyKeyProperties](#KeyVaultCertificateSpecCertificatePolicyKeyProperties)***||
| `lifetimeAction` | ***[[]KeyVaultCertificateSpecCertificatePolicyLifetimeAction](#KeyVaultCertificateSpecCertificatePolicyLifetimeAction)***| ***(Optional)*** |
| `secretProperties` | ***[[]KeyVaultCertificateSpecCertificatePolicySecretProperties](#KeyVaultCertificateSpecCertificatePolicySecretProperties)***||
| `x509CertificateProperties` | ***[[]KeyVaultCertificateSpecCertificatePolicyX509CertificateProperties](#KeyVaultCertificateSpecCertificatePolicyX509CertificateProperties)***| ***(Optional)*** |
## KeyVaultCertificateSpecCertificatePolicyIssuerParameters
##### (Appears on:[KeyVaultCertificateSpecCertificatePolicy](#KeyVaultCertificateSpecCertificatePolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
## KeyVaultCertificateSpecCertificatePolicyKeyProperties
##### (Appears on:[KeyVaultCertificateSpecCertificatePolicy](#KeyVaultCertificateSpecCertificatePolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `exportable` | ***bool***||
| `keySize` | ***int***||
| `keyType` | ***string***||
| `reuseKey` | ***bool***||
## KeyVaultCertificateSpecCertificatePolicyLifetimeAction
##### (Appears on:[KeyVaultCertificateSpecCertificatePolicy](#KeyVaultCertificateSpecCertificatePolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***[[]KeyVaultCertificateSpecCertificatePolicyLifetimeActionAction](#KeyVaultCertificateSpecCertificatePolicyLifetimeActionAction)***||
| `trigger` | ***[[]KeyVaultCertificateSpecCertificatePolicyLifetimeActionTrigger](#KeyVaultCertificateSpecCertificatePolicyLifetimeActionTrigger)***||
## KeyVaultCertificateSpecCertificatePolicyLifetimeActionAction
##### (Appears on:[KeyVaultCertificateSpecCertificatePolicyLifetimeAction](#KeyVaultCertificateSpecCertificatePolicyLifetimeAction))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `actionType` | ***string***||
## KeyVaultCertificateSpecCertificatePolicyLifetimeActionTrigger
##### (Appears on:[KeyVaultCertificateSpecCertificatePolicyLifetimeAction](#KeyVaultCertificateSpecCertificatePolicyLifetimeAction))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `daysBeforeExpiry` | ***int***| ***(Optional)*** |
| `lifetimePercentage` | ***int***| ***(Optional)*** |
## KeyVaultCertificateSpecCertificatePolicySecretProperties
##### (Appears on:[KeyVaultCertificateSpecCertificatePolicy](#KeyVaultCertificateSpecCertificatePolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `contentType` | ***string***||
## KeyVaultCertificateSpecCertificatePolicyX509CertificateProperties
##### (Appears on:[KeyVaultCertificateSpecCertificatePolicy](#KeyVaultCertificateSpecCertificatePolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `extendedKeyUsage` | ***[]string***| ***(Optional)*** |
| `keyUsage` | ***[]string***||
| `subject` | ***string***||
| `subjectAlternativeNames` | ***[[]KeyVaultCertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames](#KeyVaultCertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames)***| ***(Optional)*** |
| `validityInMonths` | ***int***||
## KeyVaultCertificateSpecCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames
##### (Appears on:[KeyVaultCertificateSpecCertificatePolicyX509CertificateProperties](#KeyVaultCertificateSpecCertificatePolicyX509CertificateProperties))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `dnsNames` | ***[]string***| ***(Optional)*** |
| `emails` | ***[]string***| ***(Optional)*** |
| `upns` | ***[]string***| ***(Optional)*** |
## KeyVaultCertificateStatus
##### (Appears on:[KeyVaultCertificate](#KeyVaultCertificate))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KeyVaultCertificateSpec](#KeyVaultCertificateSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `certificate.<index>.contents` | ***string*** ||
| `certificate.<index>.password` | ***string*** ||
