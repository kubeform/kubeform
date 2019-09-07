---
title: AppService
menu:
  docs_v0.0.1:
    identifier: appservice-azurerm.kubeform.com
    name: AppService
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## AppService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AppService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppServiceSpec](#AppServiceSpec)***||
| `status` | ***[AppServiceStatus](#AppServiceStatus)***||
## AppServiceSpec
##### (Appears on:[AppService](#AppService), [AppServiceStatus](#AppServiceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `appServicePlanID` | ***string***||
| `appSettings` | ***map[string]string***| ***(Optional)*** |
| `authSettings` | ***[[]AppServiceSpecAuthSettings](#AppServiceSpecAuthSettings)***| ***(Optional)*** |
| `backup` | ***[[]AppServiceSpecBackup](#AppServiceSpecBackup)***| ***(Optional)*** |
| `clientAffinityEnabled` | ***bool***| ***(Optional)*** |
| `clientCertEnabled` | ***bool***| ***(Optional)*** |
| `connectionString` | ***[[]AppServiceSpecConnectionString](#AppServiceSpecConnectionString)***| ***(Optional)*** |
| `defaultSiteHostname` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `httpsOnly` | ***bool***| ***(Optional)*** |
| `identity` | ***[[]AppServiceSpecIdentity](#AppServiceSpecIdentity)***| ***(Optional)*** |
| `location` | ***string***||
| `logs` | ***[[]AppServiceSpecLogs](#AppServiceSpecLogs)***| ***(Optional)*** |
| `name` | ***string***||
| `outboundIPAddresses` | ***string***| ***(Optional)*** |
| `possibleOutboundIPAddresses` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `siteConfig` | ***[[]AppServiceSpecSiteConfig](#AppServiceSpecSiteConfig)***| ***(Optional)*** |
| `siteCredential` | ***[[]AppServiceSpecSiteCredential](#AppServiceSpecSiteCredential)***| ***(Optional)*** |
| `sourceControl` | ***[[]AppServiceSpecSourceControl](#AppServiceSpecSourceControl)***| ***(Optional)*** |
| `storageAccount` | ***[[]AppServiceSpecStorageAccount](#AppServiceSpecStorageAccount)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## AppServiceSpecAuthSettings
##### (Appears on:[AppServiceSpec](#AppServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `activeDirectory` | ***[[]AppServiceSpecAuthSettingsActiveDirectory](#AppServiceSpecAuthSettingsActiveDirectory)***| ***(Optional)*** |
| `additionalLoginParams` | ***map[string]string***| ***(Optional)*** |
| `allowedExternalRedirectUrls` | ***[]string***| ***(Optional)*** |
| `defaultProvider` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***||
| `facebook` | ***[[]AppServiceSpecAuthSettingsFacebook](#AppServiceSpecAuthSettingsFacebook)***| ***(Optional)*** |
| `google` | ***[[]AppServiceSpecAuthSettingsGoogle](#AppServiceSpecAuthSettingsGoogle)***| ***(Optional)*** |
| `issuer` | ***string***| ***(Optional)*** |
| `microsoft` | ***[[]AppServiceSpecAuthSettingsMicrosoft](#AppServiceSpecAuthSettingsMicrosoft)***| ***(Optional)*** |
| `runtimeVersion` | ***string***| ***(Optional)*** |
| `tokenRefreshExtensionHours` | ***encoding/json.Number***| ***(Optional)*** |
| `tokenStoreEnabled` | ***bool***| ***(Optional)*** |
| `twitter` | ***[[]AppServiceSpecAuthSettingsTwitter](#AppServiceSpecAuthSettingsTwitter)***| ***(Optional)*** |
| `unauthenticatedClientAction` | ***string***| ***(Optional)*** |
## AppServiceSpecAuthSettingsActiveDirectory
##### (Appears on:[AppServiceSpecAuthSettings](#AppServiceSpecAuthSettings))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowedAudiences` | ***[]string***| ***(Optional)*** |
| `clientID` | ***string***||
## AppServiceSpecAuthSettingsFacebook
##### (Appears on:[AppServiceSpecAuthSettings](#AppServiceSpecAuthSettings))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `appID` | ***string***||
| `oauthScopes` | ***[]string***| ***(Optional)*** |
## AppServiceSpecAuthSettingsGoogle
##### (Appears on:[AppServiceSpecAuthSettings](#AppServiceSpecAuthSettings))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientID` | ***string***||
| `oauthScopes` | ***[]string***| ***(Optional)*** |
## AppServiceSpecAuthSettingsMicrosoft
##### (Appears on:[AppServiceSpecAuthSettings](#AppServiceSpecAuthSettings))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientID` | ***string***||
| `oauthScopes` | ***[]string***| ***(Optional)*** |
## AppServiceSpecAuthSettingsTwitter
##### (Appears on:[AppServiceSpecAuthSettings](#AppServiceSpecAuthSettings))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `consumerKey` | ***string***||
## AppServiceSpecBackup
##### (Appears on:[AppServiceSpec](#AppServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `schedule` | ***[[]AppServiceSpecBackupSchedule](#AppServiceSpecBackupSchedule)***||
## AppServiceSpecBackupSchedule
##### (Appears on:[AppServiceSpecBackup](#AppServiceSpecBackup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `frequencyInterval` | ***int***||
| `frequencyUnit` | ***string***||
| `keepAtLeastOneBackup` | ***bool***| ***(Optional)*** |
| `retentionPeriodInDays` | ***int***| ***(Optional)*** |
| `startTime` | ***string***| ***(Optional)*** |
## AppServiceSpecConnectionString
##### (Appears on:[AppServiceSpec](#AppServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `type` | ***string***||
## AppServiceSpecIdentity
##### (Appears on:[AppServiceSpec](#AppServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `identityIDS` | ***[]string***| ***(Optional)*** |
| `principalID` | ***string***| ***(Optional)*** |
| `tenantID` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## AppServiceSpecLogs
##### (Appears on:[AppServiceSpec](#AppServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `applicationLogs` | ***[[]AppServiceSpecLogsApplicationLogs](#AppServiceSpecLogsApplicationLogs)***| ***(Optional)*** |
## AppServiceSpecLogsApplicationLogs
##### (Appears on:[AppServiceSpecLogs](#AppServiceSpecLogs))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `azureBlobStorage` | ***[[]AppServiceSpecLogsApplicationLogsAzureBlobStorage](#AppServiceSpecLogsApplicationLogsAzureBlobStorage)***| ***(Optional)*** |
## AppServiceSpecLogsApplicationLogsAzureBlobStorage
##### (Appears on:[AppServiceSpecLogsApplicationLogs](#AppServiceSpecLogsApplicationLogs))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `level` | ***string***||
| `retentionInDays` | ***int***||
## AppServiceSpecSiteConfig
##### (Appears on:[AppServiceSpec](#AppServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `alwaysOn` | ***bool***| ***(Optional)*** |
| `appCommandLine` | ***string***| ***(Optional)*** |
| `cors` | ***[[]AppServiceSpecSiteConfigCors](#AppServiceSpecSiteConfigCors)***| ***(Optional)*** |
| `defaultDocuments` | ***[]string***| ***(Optional)*** |
| `dotnetFrameworkVersion` | ***string***| ***(Optional)*** |
| `ftpsState` | ***string***| ***(Optional)*** |
| `http2Enabled` | ***bool***| ***(Optional)*** |
| `ipRestriction` | ***[[]AppServiceSpecSiteConfigIpRestriction](#AppServiceSpecSiteConfigIpRestriction)***| ***(Optional)*** |
| `javaContainer` | ***string***| ***(Optional)*** |
| `javaContainerVersion` | ***string***| ***(Optional)*** |
| `javaVersion` | ***string***| ***(Optional)*** |
| `linuxFxVersion` | ***string***| ***(Optional)*** |
| `localMysqlEnabled` | ***bool***| ***(Optional)*** |
| `managedPipelineMode` | ***string***| ***(Optional)*** |
| `minTLSVersion` | ***string***| ***(Optional)*** |
| `phpVersion` | ***string***| ***(Optional)*** |
| `pythonVersion` | ***string***| ***(Optional)*** |
| `remoteDebuggingEnabled` | ***bool***| ***(Optional)*** |
| `remoteDebuggingVersion` | ***string***| ***(Optional)*** |
| `scmType` | ***string***| ***(Optional)*** |
| `use32BitWorkerProcess` | ***bool***| ***(Optional)*** |
| `virtualNetworkName` | ***string***| ***(Optional)*** |
| `websocketsEnabled` | ***bool***| ***(Optional)*** |
| `windowsFxVersion` | ***string***| ***(Optional)*** |
## AppServiceSpecSiteConfigCors
##### (Appears on:[AppServiceSpecSiteConfig](#AppServiceSpecSiteConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowedOrigins` | ***[]string***||
| `supportCredentials` | ***bool***| ***(Optional)*** |
## AppServiceSpecSiteConfigIpRestriction
##### (Appears on:[AppServiceSpecSiteConfig](#AppServiceSpecSiteConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ipAddress` | ***string***||
| `subnetMask` | ***string***| ***(Optional)*** |
## AppServiceSpecSiteCredential
##### (Appears on:[AppServiceSpec](#AppServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `username` | ***string***| ***(Optional)*** |
## AppServiceSpecSourceControl
##### (Appears on:[AppServiceSpec](#AppServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `branch` | ***string***| ***(Optional)*** |
| `repoURL` | ***string***| ***(Optional)*** |
## AppServiceSpecStorageAccount
##### (Appears on:[AppServiceSpec](#AppServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `accountName` | ***string***||
| `mountPath` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `shareName` | ***string***||
| `type` | ***string***||
## AppServiceStatus
##### (Appears on:[AppService](#AppService))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppServiceSpec](#AppServiceSpec)***| ***(Optional)*** |
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
| `backup.<index>.storage_account_url` | ***string*** ||
| `connection_string.<index>.value` | ***string*** ||
| `logs.<index>.application_logs.<index>.azure_blob_storage.<index>.sas_url` | ***string*** ||
| `site_credential.<index>.password` | ***string*** ||
| `storage_account.<index>.access_key` | ***string*** ||
