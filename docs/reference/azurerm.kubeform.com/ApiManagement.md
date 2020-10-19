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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementSpec](#apimanagementspec)***||
| `status` | ***[ApiManagementStatus](#apimanagementstatus)***||
## ApiManagementSpec

Appears on:[ApiManagement](#apimanagement), [ApiManagementStatus](#apimanagementstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `additionalLocation` | ***[[]ApiManagementSpecAdditionalLocation](#apimanagementspecadditionallocation)***| ***(Optional)*** |
| `certificate` | ***[[]ApiManagementSpecCertificate](#apimanagementspeccertificate)***| ***(Optional)*** |
| `gatewayRegionalURL` | ***string***| ***(Optional)*** |
| `gatewayURL` | ***string***| ***(Optional)*** |
| `hostnameConfiguration` | ***[[]ApiManagementSpecHostnameConfiguration](#apimanagementspechostnameconfiguration)***| ***(Optional)*** |
| `identity` | ***[[]ApiManagementSpecIdentity](#apimanagementspecidentity)***| ***(Optional)*** |
| `location` | ***string***||
| `managementAPIURL` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `notificationSenderEmail` | ***string***| ***(Optional)*** |
| `policy` | ***[[]ApiManagementSpecPolicy](#apimanagementspecpolicy)***| ***(Optional)*** |
| `portalURL` | ***string***| ***(Optional)*** |
| `protocols` | ***[[]ApiManagementSpecProtocols](#apimanagementspecprotocols)***| ***(Optional)*** |
| `publicIPAddresses` | ***[]string***| ***(Optional)*** |
| `publisherEmail` | ***string***||
| `publisherName` | ***string***||
| `resourceGroupName` | ***string***||
| `scmURL` | ***string***| ***(Optional)*** |
| `security` | ***[[]ApiManagementSpecSecurity](#apimanagementspecsecurity)***| ***(Optional)*** |
| `signIn` | ***[[]ApiManagementSpecSignIn](#apimanagementspecsignin)***| ***(Optional)*** |
| `signUp` | ***[[]ApiManagementSpecSignUp](#apimanagementspecsignup)***| ***(Optional)*** |
| `sku` | ***[[]ApiManagementSpecSku](#apimanagementspecsku)***| ***(Optional)*** Deprecated|
| `skuName` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## ApiManagementSpecAdditionalLocation

Appears on:[ApiManagementSpec](#apimanagementspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `gatewayRegionalURL` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `publicIPAddresses` | ***[]string***| ***(Optional)*** |
## ApiManagementSpecCertificate

Appears on:[ApiManagementSpec](#apimanagementspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `storeName` | ***string***||
## ApiManagementSpecHostnameConfiguration

Appears on:[ApiManagementSpec](#apimanagementspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `management` | ***[[]ApiManagementSpecHostnameConfigurationManagement](#apimanagementspechostnameconfigurationmanagement)***| ***(Optional)*** |
| `portal` | ***[[]ApiManagementSpecHostnameConfigurationPortal](#apimanagementspechostnameconfigurationportal)***| ***(Optional)*** |
| `proxy` | ***[[]ApiManagementSpecHostnameConfigurationProxy](#apimanagementspechostnameconfigurationproxy)***| ***(Optional)*** |
| `scm` | ***[[]ApiManagementSpecHostnameConfigurationScm](#apimanagementspechostnameconfigurationscm)***| ***(Optional)*** |
## ApiManagementSpecHostnameConfigurationManagement

Appears on:[ApiManagementSpecHostnameConfiguration](#apimanagementspechostnameconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hostName` | ***string***||
| `keyVaultID` | ***string***| ***(Optional)*** |
| `negotiateClientCertificate` | ***bool***| ***(Optional)*** |
## ApiManagementSpecHostnameConfigurationPortal

Appears on:[ApiManagementSpecHostnameConfiguration](#apimanagementspechostnameconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hostName` | ***string***||
| `keyVaultID` | ***string***| ***(Optional)*** |
| `negotiateClientCertificate` | ***bool***| ***(Optional)*** |
## ApiManagementSpecHostnameConfigurationProxy

Appears on:[ApiManagementSpecHostnameConfiguration](#apimanagementspechostnameconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultSSLBinding` | ***bool***| ***(Optional)*** |
| `hostName` | ***string***||
| `keyVaultID` | ***string***| ***(Optional)*** |
| `negotiateClientCertificate` | ***bool***| ***(Optional)*** |
## ApiManagementSpecHostnameConfigurationScm

Appears on:[ApiManagementSpecHostnameConfiguration](#apimanagementspechostnameconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hostName` | ***string***||
| `keyVaultID` | ***string***| ***(Optional)*** |
| `negotiateClientCertificate` | ***bool***| ***(Optional)*** |
## ApiManagementSpecIdentity

Appears on:[ApiManagementSpec](#apimanagementspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `principalID` | ***string***| ***(Optional)*** |
| `tenantID` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## ApiManagementSpecPolicy

Appears on:[ApiManagementSpec](#apimanagementspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `xmlContent` | ***string***| ***(Optional)*** |
| `xmlLink` | ***string***| ***(Optional)*** |
## ApiManagementSpecProtocols

Appears on:[ApiManagementSpec](#apimanagementspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enableHttp2` | ***bool***| ***(Optional)*** |
## ApiManagementSpecSecurity

Appears on:[ApiManagementSpec](#apimanagementspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `disableBackendSSL30` | ***bool***| ***(Optional)*** Deprecated|
| `disableBackendTLS10` | ***bool***| ***(Optional)*** Deprecated|
| `disableBackendTLS11` | ***bool***| ***(Optional)*** Deprecated|
| `disableFrontendSSL30` | ***bool***| ***(Optional)*** Deprecated|
| `disableFrontendTLS10` | ***bool***| ***(Optional)*** Deprecated|
| `disableFrontendTLS11` | ***bool***| ***(Optional)*** Deprecated|
| `disableTripleDESChipers` | ***bool***| ***(Optional)*** Deprecated|
| `disableTripleDESCiphers` | ***bool***| ***(Optional)*** Deprecated|
| `enableBackendSSL30` | ***bool***| ***(Optional)*** |
| `enableBackendTLS10` | ***bool***| ***(Optional)*** |
| `enableBackendTLS11` | ***bool***| ***(Optional)*** |
| `enableFrontendSSL30` | ***bool***| ***(Optional)*** |
| `enableFrontendTLS10` | ***bool***| ***(Optional)*** |
| `enableFrontendTLS11` | ***bool***| ***(Optional)*** |
| `enableTripleDESCiphers` | ***bool***| ***(Optional)*** |
## ApiManagementSpecSignIn

Appears on:[ApiManagementSpec](#apimanagementspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
## ApiManagementSpecSignUp

Appears on:[ApiManagementSpec](#apimanagementspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `termsOfService` | ***[[]ApiManagementSpecSignUpTermsOfService](#apimanagementspecsignuptermsofservice)***||
## ApiManagementSpecSignUpTermsOfService

Appears on:[ApiManagementSpecSignUp](#apimanagementspecsignup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `consentRequired` | ***bool***||
| `enabled` | ***bool***||
| `text` | ***string***| ***(Optional)*** |
## ApiManagementSpecSku

Appears on:[ApiManagementSpec](#apimanagementspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
## ApiManagementStatus

Appears on:[ApiManagement](#apimanagement)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementSpec](#apimanagementspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ApiManagementStatus](#apimanagementstatus)

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
