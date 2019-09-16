---
title: ApiManagement
menu:
  docs_{{ .version }}:
    identifier: apimanagement-azurerm.kubeform.com
    name: ApiManagement
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiManagement
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagement` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementSpec](#ApiManagementSpec)***||
| `status` | ***[ApiManagementStatus](#ApiManagementStatus)***||
## ApiManagementSpec

Appears on:[ApiManagement](#ApiManagement), [ApiManagementStatus](#ApiManagementStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `additionalLocation` | ***[[]ApiManagementSpecAdditionalLocation](#ApiManagementSpecAdditionalLocation)***| ***(Optional)*** |
| `certificate` | ***[[]ApiManagementSpecCertificate](#ApiManagementSpecCertificate)***| ***(Optional)*** |
| `gatewayRegionalURL` | ***string***| ***(Optional)*** |
| `gatewayURL` | ***string***| ***(Optional)*** |
| `hostnameConfiguration` | ***[[]ApiManagementSpecHostnameConfiguration](#ApiManagementSpecHostnameConfiguration)***| ***(Optional)*** |
| `identity` | ***[[]ApiManagementSpecIdentity](#ApiManagementSpecIdentity)***| ***(Optional)*** |
| `location` | ***string***||
| `managementAPIURL` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `notificationSenderEmail` | ***string***| ***(Optional)*** |
| `policy` | ***[[]ApiManagementSpecPolicy](#ApiManagementSpecPolicy)***| ***(Optional)*** |
| `portalURL` | ***string***| ***(Optional)*** |
| `publicIPAddresses` | ***[]string***| ***(Optional)*** |
| `publisherEmail` | ***string***||
| `publisherName` | ***string***||
| `resourceGroupName` | ***string***||
| `scmURL` | ***string***| ***(Optional)*** |
| `security` | ***[[]ApiManagementSpecSecurity](#ApiManagementSpecSecurity)***| ***(Optional)*** |
| `signIn` | ***[[]ApiManagementSpecSignIn](#ApiManagementSpecSignIn)***| ***(Optional)*** |
| `signUp` | ***[[]ApiManagementSpecSignUp](#ApiManagementSpecSignUp)***| ***(Optional)*** |
| `sku` | ***[[]ApiManagementSpecSku](#ApiManagementSpecSku)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## ApiManagementSpecAdditionalLocation

Appears on:[ApiManagementSpec](#ApiManagementSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `gatewayRegionalURL` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `publicIPAddresses` | ***[]string***| ***(Optional)*** |
## ApiManagementSpecCertificate

Appears on:[ApiManagementSpec](#ApiManagementSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `storeName` | ***string***||
## ApiManagementSpecHostnameConfiguration

Appears on:[ApiManagementSpec](#ApiManagementSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `management` | ***[[]ApiManagementSpecHostnameConfigurationManagement](#ApiManagementSpecHostnameConfigurationManagement)***| ***(Optional)*** |
| `portal` | ***[[]ApiManagementSpecHostnameConfigurationPortal](#ApiManagementSpecHostnameConfigurationPortal)***| ***(Optional)*** |
| `proxy` | ***[[]ApiManagementSpecHostnameConfigurationProxy](#ApiManagementSpecHostnameConfigurationProxy)***| ***(Optional)*** |
| `scm` | ***[[]ApiManagementSpecHostnameConfigurationScm](#ApiManagementSpecHostnameConfigurationScm)***| ***(Optional)*** |
## ApiManagementSpecHostnameConfigurationManagement

Appears on:[ApiManagementSpecHostnameConfiguration](#ApiManagementSpecHostnameConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hostName` | ***string***||
| `keyVaultID` | ***string***| ***(Optional)*** |
| `negotiateClientCertificate` | ***bool***| ***(Optional)*** |
## ApiManagementSpecHostnameConfigurationPortal

Appears on:[ApiManagementSpecHostnameConfiguration](#ApiManagementSpecHostnameConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hostName` | ***string***||
| `keyVaultID` | ***string***| ***(Optional)*** |
| `negotiateClientCertificate` | ***bool***| ***(Optional)*** |
## ApiManagementSpecHostnameConfigurationProxy

Appears on:[ApiManagementSpecHostnameConfiguration](#ApiManagementSpecHostnameConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultSSLBinding` | ***bool***| ***(Optional)*** |
| `hostName` | ***string***||
| `keyVaultID` | ***string***| ***(Optional)*** |
| `negotiateClientCertificate` | ***bool***| ***(Optional)*** |
## ApiManagementSpecHostnameConfigurationScm

Appears on:[ApiManagementSpecHostnameConfiguration](#ApiManagementSpecHostnameConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hostName` | ***string***||
| `keyVaultID` | ***string***| ***(Optional)*** |
| `negotiateClientCertificate` | ***bool***| ***(Optional)*** |
## ApiManagementSpecIdentity

Appears on:[ApiManagementSpec](#ApiManagementSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `principalID` | ***string***| ***(Optional)*** |
| `tenantID` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## ApiManagementSpecPolicy

Appears on:[ApiManagementSpec](#ApiManagementSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `xmlContent` | ***string***| ***(Optional)*** |
| `xmlLink` | ***string***| ***(Optional)*** |
## ApiManagementSpecSecurity

Appears on:[ApiManagementSpec](#ApiManagementSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `disableBackendSSL30` | ***bool***| ***(Optional)*** |
| `disableBackendTLS10` | ***bool***| ***(Optional)*** |
| `disableBackendTLS11` | ***bool***| ***(Optional)*** |
| `disableFrontendSSL30` | ***bool***| ***(Optional)*** |
| `disableFrontendTLS10` | ***bool***| ***(Optional)*** |
| `disableFrontendTLS11` | ***bool***| ***(Optional)*** |
| `disableTripleDESChipers` | ***bool***| ***(Optional)*** Deprecated|
| `disableTripleDESCiphers` | ***bool***| ***(Optional)*** |
## ApiManagementSpecSignIn

Appears on:[ApiManagementSpec](#ApiManagementSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
## ApiManagementSpecSignUp

Appears on:[ApiManagementSpec](#ApiManagementSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `termsOfService` | ***[[]ApiManagementSpecSignUpTermsOfService](#ApiManagementSpecSignUpTermsOfService)***||
## ApiManagementSpecSignUpTermsOfService

Appears on:[ApiManagementSpecSignUp](#ApiManagementSpecSignUp)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `consentRequired` | ***bool***||
| `enabled` | ***bool***||
| `text` | ***string***| ***(Optional)*** |
## ApiManagementSpecSku

Appears on:[ApiManagementSpec](#ApiManagementSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int***||
| `name` | ***string***||
## ApiManagementStatus

Appears on:[ApiManagement](#ApiManagement)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementSpec](#ApiManagementSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `certificate.<index>.certificate_password` | ***string*** ||
| `certificate.<index>.encoded_certificate` | ***string*** ||
| `hostname_configuration.<index>.management.<index>.certificate` | ***string*** ||
| `hostname_configuration.<index>.management.<index>.certificate_password` | ***string*** ||
| `hostname_configuration.<index>.portal.<index>.certificate` | ***string*** ||
| `hostname_configuration.<index>.portal.<index>.certificate_password` | ***string*** ||
| `hostname_configuration.<index>.proxy.<index>.certificate` | ***string*** ||
| `hostname_configuration.<index>.proxy.<index>.certificate_password` | ***string*** ||
| `hostname_configuration.<index>.scm.<index>.certificate` | ***string*** ||
| `hostname_configuration.<index>.scm.<index>.certificate_password` | ***string*** ||
