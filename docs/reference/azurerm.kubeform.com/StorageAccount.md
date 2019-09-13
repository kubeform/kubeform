---
title: StorageAccount
menu:
  docs_{{ .version }}:
    identifier: storageaccount-azurerm.kubeform.com
    name: StorageAccount
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageAccount
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageAccount` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageAccountSpec](#StorageAccountSpec)***||
| `status` | ***[StorageAccountStatus](#StorageAccountStatus)***||
## StorageAccountSpec
##### (Appears on:[StorageAccount](#StorageAccount), [StorageAccountStatus](#StorageAccountStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `accessTier` | ***string***| ***(Optional)*** |
| `accountEncryptionSource` | ***string***| ***(Optional)*** |
| `accountKind` | ***string***| ***(Optional)*** |
| `accountReplicationType` | ***string***||
| `accountTier` | ***string***||
| `accountType` | ***string***| ***(Optional)*** Deprecated|
| `customDomain` | ***[[]StorageAccountSpecCustomDomain](#StorageAccountSpecCustomDomain)***| ***(Optional)*** |
| `enableAdvancedThreatProtection` | ***bool***| ***(Optional)*** |
| `enableBlobEncryption` | ***bool***| ***(Optional)*** |
| `enableFileEncryption` | ***bool***| ***(Optional)*** |
| `enableHTTPSTrafficOnly` | ***bool***| ***(Optional)*** |
| `identity` | ***[[]StorageAccountSpecIdentity](#StorageAccountSpecIdentity)***| ***(Optional)*** |
| `isHnsEnabled` | ***bool***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `networkRules` | ***[[]StorageAccountSpecNetworkRules](#StorageAccountSpecNetworkRules)***| ***(Optional)*** |
| `primaryBlobEndpoint` | ***string***| ***(Optional)*** |
| `primaryBlobHost` | ***string***| ***(Optional)*** |
| `primaryDfsEndpoint` | ***string***| ***(Optional)*** |
| `primaryDfsHost` | ***string***| ***(Optional)*** |
| `primaryFileEndpoint` | ***string***| ***(Optional)*** |
| `primaryFileHost` | ***string***| ***(Optional)*** |
| `primaryLocation` | ***string***| ***(Optional)*** |
| `primaryQueueEndpoint` | ***string***| ***(Optional)*** |
| `primaryQueueHost` | ***string***| ***(Optional)*** |
| `primaryTableEndpoint` | ***string***| ***(Optional)*** |
| `primaryTableHost` | ***string***| ***(Optional)*** |
| `primaryWebEndpoint` | ***string***| ***(Optional)*** |
| `primaryWebHost` | ***string***| ***(Optional)*** |
| `queueProperties` | ***[[]StorageAccountSpecQueueProperties](#StorageAccountSpecQueueProperties)***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `secondaryBlobEndpoint` | ***string***| ***(Optional)*** |
| `secondaryBlobHost` | ***string***| ***(Optional)*** |
| `secondaryDfsEndpoint` | ***string***| ***(Optional)*** |
| `secondaryDfsHost` | ***string***| ***(Optional)*** |
| `secondaryFileEndpoint` | ***string***| ***(Optional)*** |
| `secondaryFileHost` | ***string***| ***(Optional)*** |
| `secondaryLocation` | ***string***| ***(Optional)*** |
| `secondaryQueueEndpoint` | ***string***| ***(Optional)*** |
| `secondaryQueueHost` | ***string***| ***(Optional)*** |
| `secondaryTableEndpoint` | ***string***| ***(Optional)*** |
| `secondaryTableHost` | ***string***| ***(Optional)*** |
| `secondaryWebEndpoint` | ***string***| ***(Optional)*** |
| `secondaryWebHost` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## StorageAccountSpecCustomDomain
##### (Appears on:[StorageAccountSpec](#StorageAccountSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `useSubdomain` | ***bool***| ***(Optional)*** |
## StorageAccountSpecIdentity
##### (Appears on:[StorageAccountSpec](#StorageAccountSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `principalID` | ***string***| ***(Optional)*** |
| `tenantID` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## StorageAccountSpecNetworkRules
##### (Appears on:[StorageAccountSpec](#StorageAccountSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bypass` | ***[]string***| ***(Optional)*** |
| `defaultAction` | ***string***| ***(Optional)*** |
| `ipRules` | ***[]string***| ***(Optional)*** |
| `virtualNetworkSubnetIDS` | ***[]string***| ***(Optional)*** |
## StorageAccountSpecQueueProperties
##### (Appears on:[StorageAccountSpec](#StorageAccountSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `corsRule` | ***[[]StorageAccountSpecQueuePropertiesCorsRule](#StorageAccountSpecQueuePropertiesCorsRule)***| ***(Optional)*** |
| `hourMetrics` | ***[[]StorageAccountSpecQueuePropertiesHourMetrics](#StorageAccountSpecQueuePropertiesHourMetrics)***| ***(Optional)*** |
| `logging` | ***[[]StorageAccountSpecQueuePropertiesLogging](#StorageAccountSpecQueuePropertiesLogging)***| ***(Optional)*** |
| `minuteMetrics` | ***[[]StorageAccountSpecQueuePropertiesMinuteMetrics](#StorageAccountSpecQueuePropertiesMinuteMetrics)***| ***(Optional)*** |
## StorageAccountSpecQueuePropertiesCorsRule
##### (Appears on:[StorageAccountSpecQueueProperties](#StorageAccountSpecQueueProperties))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowedHeaders` | ***[]string***||
| `allowedMethods` | ***[]string***||
| `allowedOrigins` | ***[]string***||
| `exposedHeaders` | ***[]string***||
| `maxAgeInSeconds` | ***int***||
## StorageAccountSpecQueuePropertiesHourMetrics
##### (Appears on:[StorageAccountSpecQueueProperties](#StorageAccountSpecQueueProperties))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `includeApis` | ***bool***| ***(Optional)*** |
| `retentionPolicyDays` | ***int***| ***(Optional)*** |
| `version` | ***string***||
## StorageAccountSpecQueuePropertiesLogging
##### (Appears on:[StorageAccountSpecQueueProperties](#StorageAccountSpecQueueProperties))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `delete` | ***bool***||
| `read` | ***bool***||
| `retentionPolicyDays` | ***int***| ***(Optional)*** |
| `version` | ***string***||
| `write` | ***bool***||
## StorageAccountSpecQueuePropertiesMinuteMetrics
##### (Appears on:[StorageAccountSpecQueueProperties](#StorageAccountSpecQueueProperties))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `includeApis` | ***bool***| ***(Optional)*** |
| `retentionPolicyDays` | ***int***| ***(Optional)*** |
| `version` | ***string***||
## StorageAccountStatus
##### (Appears on:[StorageAccount](#StorageAccount))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageAccountSpec](#StorageAccountSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_access_key` | ***string*** ||
| `primary_blob_connection_string` | ***string*** ||
| `primary_connection_string` | ***string*** ||
| `secondary_access_key` | ***string*** ||
| `secondary_blob_connection_string` | ***string*** ||
| `secondary_connection_string` | ***string*** ||
