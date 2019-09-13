---
title: FunctionApp
menu:
  docs_{{ .version }}:
    identifier: functionapp-azurerm.kubeform.com
    name: FunctionApp
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## FunctionApp
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `FunctionApp` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FunctionAppSpec](#FunctionAppSpec)***||
| `status` | ***[FunctionAppStatus](#FunctionAppStatus)***||
## FunctionAppSpec
##### (Appears on:[FunctionApp](#FunctionApp), [FunctionAppStatus](#FunctionAppStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `appServicePlanID` | ***string***||
| `appSettings` | ***map[string]string***| ***(Optional)*** |
| `authSettings` | ***[[]FunctionAppSpecAuthSettings](#FunctionAppSpecAuthSettings)***| ***(Optional)*** |
| `clientAffinityEnabled` | ***bool***| ***(Optional)*** |
| `connectionString` | ***[[]FunctionAppSpecConnectionString](#FunctionAppSpecConnectionString)***| ***(Optional)*** |
| `defaultHostname` | ***string***| ***(Optional)*** |
| `enableBuiltinLogging` | ***bool***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `httpsOnly` | ***bool***| ***(Optional)*** |
| `identity` | ***[[]FunctionAppSpecIdentity](#FunctionAppSpecIdentity)***| ***(Optional)*** |
| `kind` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `outboundIPAddresses` | ***string***| ***(Optional)*** |
| `possibleOutboundIPAddresses` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `siteConfig` | ***[[]FunctionAppSpecSiteConfig](#FunctionAppSpecSiteConfig)***| ***(Optional)*** |
| `siteCredential` | ***[[]FunctionAppSpecSiteCredential](#FunctionAppSpecSiteCredential)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
## FunctionAppSpecAuthSettings
##### (Appears on:[FunctionAppSpec](#FunctionAppSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `activeDirectory` | ***[[]FunctionAppSpecAuthSettingsActiveDirectory](#FunctionAppSpecAuthSettingsActiveDirectory)***| ***(Optional)*** |
| `additionalLoginParams` | ***map[string]string***| ***(Optional)*** |
| `allowedExternalRedirectUrls` | ***[]string***| ***(Optional)*** |
| `defaultProvider` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***||
| `facebook` | ***[[]FunctionAppSpecAuthSettingsFacebook](#FunctionAppSpecAuthSettingsFacebook)***| ***(Optional)*** |
| `google` | ***[[]FunctionAppSpecAuthSettingsGoogle](#FunctionAppSpecAuthSettingsGoogle)***| ***(Optional)*** |
| `issuer` | ***string***| ***(Optional)*** |
| `microsoft` | ***[[]FunctionAppSpecAuthSettingsMicrosoft](#FunctionAppSpecAuthSettingsMicrosoft)***| ***(Optional)*** |
| `runtimeVersion` | ***string***| ***(Optional)*** |
| `tokenRefreshExtensionHours` | ***encoding/json.Number***| ***(Optional)*** |
| `tokenStoreEnabled` | ***bool***| ***(Optional)*** |
| `twitter` | ***[[]FunctionAppSpecAuthSettingsTwitter](#FunctionAppSpecAuthSettingsTwitter)***| ***(Optional)*** |
| `unauthenticatedClientAction` | ***string***| ***(Optional)*** |
## FunctionAppSpecAuthSettingsActiveDirectory
##### (Appears on:[FunctionAppSpecAuthSettings](#FunctionAppSpecAuthSettings))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowedAudiences` | ***[]string***| ***(Optional)*** |
| `clientID` | ***string***||
## FunctionAppSpecAuthSettingsFacebook
##### (Appears on:[FunctionAppSpecAuthSettings](#FunctionAppSpecAuthSettings))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `appID` | ***string***||
| `oauthScopes` | ***[]string***| ***(Optional)*** |
## FunctionAppSpecAuthSettingsGoogle
##### (Appears on:[FunctionAppSpecAuthSettings](#FunctionAppSpecAuthSettings))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientID` | ***string***||
| `oauthScopes` | ***[]string***| ***(Optional)*** |
## FunctionAppSpecAuthSettingsMicrosoft
##### (Appears on:[FunctionAppSpecAuthSettings](#FunctionAppSpecAuthSettings))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientID` | ***string***||
| `oauthScopes` | ***[]string***| ***(Optional)*** |
## FunctionAppSpecAuthSettingsTwitter
##### (Appears on:[FunctionAppSpecAuthSettings](#FunctionAppSpecAuthSettings))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `consumerKey` | ***string***||
## FunctionAppSpecConnectionString
##### (Appears on:[FunctionAppSpec](#FunctionAppSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `type` | ***string***||
## FunctionAppSpecIdentity
##### (Appears on:[FunctionAppSpec](#FunctionAppSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `principalID` | ***string***| ***(Optional)*** |
| `tenantID` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## FunctionAppSpecSiteConfig
##### (Appears on:[FunctionAppSpec](#FunctionAppSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `alwaysOn` | ***bool***| ***(Optional)*** |
| `linuxFxVersion` | ***string***| ***(Optional)*** |
| `use32BitWorkerProcess` | ***bool***| ***(Optional)*** |
| `websocketsEnabled` | ***bool***| ***(Optional)*** |
## FunctionAppSpecSiteCredential
##### (Appears on:[FunctionAppSpec](#FunctionAppSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `username` | ***string***| ***(Optional)*** |
## FunctionAppStatus
##### (Appears on:[FunctionApp](#FunctionApp))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FunctionAppSpec](#FunctionAppSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `auth_settings.<index>.active_directory.<index>.client_secret` | ***string*** ||
| `auth_settings.<index>.facebook.<index>.app_secret` | ***string*** ||
| `auth_settings.<index>.google.<index>.client_secret` | ***string*** ||
| `auth_settings.<index>.microsoft.<index>.client_secret` | ***string*** ||
| `auth_settings.<index>.twitter.<index>.consumer_secret` | ***string*** ||
| `connection_string.<index>.value` | ***string*** ||
| `site_credential.<index>.password` | ***string*** ||
| `storage_connection_string` | ***string*** ||
