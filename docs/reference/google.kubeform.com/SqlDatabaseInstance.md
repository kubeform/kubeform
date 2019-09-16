---
title: SqlDatabaseInstance
menu:
  docs_{{ .version }}:
    identifier: sqldatabaseinstance-google.kubeform.com
    name: SqlDatabaseInstance
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SqlDatabaseInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `SqlDatabaseInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SqlDatabaseInstanceSpec](#SqlDatabaseInstanceSpec)***||
| `status` | ***[SqlDatabaseInstanceStatus](#SqlDatabaseInstanceStatus)***||
## SqlDatabaseInstanceSpec

Appears on:[SqlDatabaseInstance](#SqlDatabaseInstance), [SqlDatabaseInstanceStatus](#SqlDatabaseInstanceStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `connectionName` | ***string***| ***(Optional)*** |
| `databaseVersion` | ***string***| ***(Optional)*** |
| `firstIPAddress` | ***string***| ***(Optional)*** |
| `ipAddress` | ***[[]SqlDatabaseInstanceSpecIpAddress](#SqlDatabaseInstanceSpecIpAddress)***| ***(Optional)*** |
| `masterInstanceName` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `replicaConfiguration` | ***[[]SqlDatabaseInstanceSpecReplicaConfiguration](#SqlDatabaseInstanceSpecReplicaConfiguration)***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `serverCaCert` | ***[[]SqlDatabaseInstanceSpecServerCaCert](#SqlDatabaseInstanceSpecServerCaCert)***| ***(Optional)*** |
| `serviceAccountEmailAddress` | ***string***| ***(Optional)*** |
| `settings` | ***[[]SqlDatabaseInstanceSpecSettings](#SqlDatabaseInstanceSpecSettings)***||
## SqlDatabaseInstanceSpecIpAddress

Appears on:[SqlDatabaseInstanceSpec](#SqlDatabaseInstanceSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ipAddress` | ***string***| ***(Optional)*** |
| `timeToRetire` | ***string***| ***(Optional)*** |
## SqlDatabaseInstanceSpecReplicaConfiguration

Appears on:[SqlDatabaseInstanceSpec](#SqlDatabaseInstanceSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `caCertificate` | ***string***| ***(Optional)*** |
| `clientCertificate` | ***string***| ***(Optional)*** |
| `clientKey` | ***string***| ***(Optional)*** |
| `connectRetryInterval` | ***int***| ***(Optional)*** |
| `dumpFilePath` | ***string***| ***(Optional)*** |
| `failoverTarget` | ***bool***| ***(Optional)*** |
| `masterHeartbeatPeriod` | ***int***| ***(Optional)*** |
| `sslCipher` | ***string***| ***(Optional)*** |
| `username` | ***string***| ***(Optional)*** |
| `verifyServerCertificate` | ***bool***| ***(Optional)*** |
## SqlDatabaseInstanceSpecServerCaCert

Appears on:[SqlDatabaseInstanceSpec](#SqlDatabaseInstanceSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cert` | ***string***| ***(Optional)*** |
| `commonName` | ***string***| ***(Optional)*** |
| `createTime` | ***string***| ***(Optional)*** |
| `expirationTime` | ***string***| ***(Optional)*** |
| `sha1Fingerprint` | ***string***| ***(Optional)*** |
## SqlDatabaseInstanceSpecSettings

Appears on:[SqlDatabaseInstanceSpec](#SqlDatabaseInstanceSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `activationPolicy` | ***string***| ***(Optional)*** |
| `authorizedGaeApplications` | ***[]string***| ***(Optional)*** |
| `availabilityType` | ***string***| ***(Optional)*** |
| `backupConfiguration` | ***[[]SqlDatabaseInstanceSpecSettingsBackupConfiguration](#SqlDatabaseInstanceSpecSettingsBackupConfiguration)***| ***(Optional)*** |
| `crashSafeReplication` | ***bool***| ***(Optional)*** |
| `databaseFlags` | ***[[]SqlDatabaseInstanceSpecSettingsDatabaseFlags](#SqlDatabaseInstanceSpecSettingsDatabaseFlags)***| ***(Optional)*** |
| `diskAutoresize` | ***bool***| ***(Optional)*** |
| `diskSize` | ***int***| ***(Optional)*** |
| `diskType` | ***string***| ***(Optional)*** |
| `ipConfiguration` | ***[[]SqlDatabaseInstanceSpecSettingsIpConfiguration](#SqlDatabaseInstanceSpecSettingsIpConfiguration)***| ***(Optional)*** |
| `locationPreference` | ***[[]SqlDatabaseInstanceSpecSettingsLocationPreference](#SqlDatabaseInstanceSpecSettingsLocationPreference)***| ***(Optional)*** |
| `maintenanceWindow` | ***[[]SqlDatabaseInstanceSpecSettingsMaintenanceWindow](#SqlDatabaseInstanceSpecSettingsMaintenanceWindow)***| ***(Optional)*** |
| `pricingPlan` | ***string***| ***(Optional)*** |
| `replicationType` | ***string***| ***(Optional)*** |
| `tier` | ***string***||
| `userLabels` | ***map[string]string***| ***(Optional)*** |
| `version` | ***int***| ***(Optional)*** |
## SqlDatabaseInstanceSpecSettingsBackupConfiguration

Appears on:[SqlDatabaseInstanceSpecSettings](#SqlDatabaseInstanceSpecSettings)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `binaryLogEnabled` | ***bool***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `startTime` | ***string***| ***(Optional)*** |
## SqlDatabaseInstanceSpecSettingsDatabaseFlags

Appears on:[SqlDatabaseInstanceSpecSettings](#SqlDatabaseInstanceSpecSettings)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
| `value` | ***string***| ***(Optional)*** |
## SqlDatabaseInstanceSpecSettingsIpConfiguration

Appears on:[SqlDatabaseInstanceSpecSettings](#SqlDatabaseInstanceSpecSettings)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authorizedNetworks` | ***[[]SqlDatabaseInstanceSpecSettingsIpConfigurationAuthorizedNetworks](#SqlDatabaseInstanceSpecSettingsIpConfigurationAuthorizedNetworks)***| ***(Optional)*** |
| `ipv4Enabled` | ***bool***| ***(Optional)*** |
| `privateNetwork` | ***string***| ***(Optional)*** |
| `requireSSL` | ***bool***| ***(Optional)*** |
## SqlDatabaseInstanceSpecSettingsIpConfigurationAuthorizedNetworks

Appears on:[SqlDatabaseInstanceSpecSettingsIpConfiguration](#SqlDatabaseInstanceSpecSettingsIpConfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `expirationTime` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `value` | ***string***| ***(Optional)*** |
## SqlDatabaseInstanceSpecSettingsLocationPreference

Appears on:[SqlDatabaseInstanceSpecSettings](#SqlDatabaseInstanceSpecSettings)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `followGaeApplication` | ***string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## SqlDatabaseInstanceSpecSettingsMaintenanceWindow

Appears on:[SqlDatabaseInstanceSpecSettings](#SqlDatabaseInstanceSpecSettings)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `day` | ***int***| ***(Optional)*** |
| `hour` | ***int***| ***(Optional)*** |
| `updateTrack` | ***string***| ***(Optional)*** |
## SqlDatabaseInstanceStatus

Appears on:[SqlDatabaseInstance](#SqlDatabaseInstance)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SqlDatabaseInstanceSpec](#SqlDatabaseInstanceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `replica_configuration.<index>.password` | ***string*** ||
