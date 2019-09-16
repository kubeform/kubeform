---
title: SqlDatabase
menu:
  docs_{{ .version }}:
    identifier: sqldatabase-azurerm.kubeform.com
    name: SqlDatabase
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SqlDatabase
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SqlDatabase` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SqlDatabaseSpec](#sqldatabasespec)***||
| `status` | ***[SqlDatabaseStatus](#sqldatabasestatus)***||
## SqlDatabaseSpec

Appears on:[SqlDatabase](#sqldatabase), [SqlDatabaseStatus](#sqldatabasestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `collation` | ***string***| ***(Optional)*** |
| `createMode` | ***string***| ***(Optional)*** |
| `creationDate` | ***string***| ***(Optional)*** |
| `defaultSecondaryLocation` | ***string***| ***(Optional)*** |
| `edition` | ***string***| ***(Optional)*** |
| `elasticPoolName` | ***string***| ***(Optional)*** |
| `encryption` | ***string***| ***(Optional)*** |
| `import` | ***[[]SqlDatabaseSpecImport](#sqldatabasespecimport)***| ***(Optional)*** |
| `location` | ***string***||
| `maxSizeBytes` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `readScale` | ***bool***| ***(Optional)*** |
| `requestedServiceObjectiveID` | ***string***| ***(Optional)*** |
| `requestedServiceObjectiveName` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `restorePointInTime` | ***string***| ***(Optional)*** |
| `serverName` | ***string***||
| `sourceDatabaseDeletionDate` | ***string***| ***(Optional)*** |
| `sourceDatabaseID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `threatDetectionPolicy` | ***[[]SqlDatabaseSpecThreatDetectionPolicy](#sqldatabasespecthreatdetectionpolicy)***| ***(Optional)*** |
## SqlDatabaseSpecImport

Appears on:[SqlDatabaseSpec](#sqldatabasespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `administratorLogin` | ***string***||
| `authenticationType` | ***string***||
| `operationMode` | ***string***| ***(Optional)*** |
| `storageKeyType` | ***string***||
| `storageURI` | ***string***||
## SqlDatabaseSpecThreatDetectionPolicy

Appears on:[SqlDatabaseSpec](#sqldatabasespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `disabledAlerts` | ***[]string***| ***(Optional)*** |
| `emailAccountAdmins` | ***string***| ***(Optional)*** |
| `emailAddresses` | ***[]string***| ***(Optional)*** |
| `retentionDays` | ***int***| ***(Optional)*** |
| `state` | ***string***| ***(Optional)*** |
| `storageEndpoint` | ***string***| ***(Optional)*** |
| `useServerDefault` | ***string***| ***(Optional)*** |
## SqlDatabaseStatus

Appears on:[SqlDatabase](#sqldatabase)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SqlDatabaseSpec](#sqldatabasespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `import.<index>.administrator_login_password` | ***string*** ||
| `import.<index>.storage_key` | ***string*** ||
| `threat_detection_policy.<index>.storage_account_access_key` | ***string*** ||
