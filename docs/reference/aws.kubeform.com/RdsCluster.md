---
title: RdsCluster
menu:
  docs_{{ .version }}:
    identifier: rdscluster-aws.kubeform.com
    name: RdsCluster
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RdsCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `RdsCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RdsClusterSpec](#rdsclusterspec)***||
| `status` | ***[RdsClusterStatus](#rdsclusterstatus)***||
## RdsClusterSpec

Appears on:[RdsCluster](#rdscluster), [RdsClusterStatus](#rdsclusterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `applyImmediately` | ***bool***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `availabilityZones` | ***[]string***| ***(Optional)*** |
| `backtrackWindow` | ***int***| ***(Optional)*** |
| `backupRetentionPeriod` | ***int***| ***(Optional)*** |
| `clusterIdentifier` | ***string***| ***(Optional)*** |
| `clusterIdentifierPrefix` | ***string***| ***(Optional)*** |
| `clusterMembers` | ***[]string***| ***(Optional)*** |
| `clusterResourceID` | ***string***| ***(Optional)*** |
| `copyTagsToSnapshot` | ***bool***| ***(Optional)*** |
| `databaseName` | ***string***| ***(Optional)*** |
| `dbClusterParameterGroupName` | ***string***| ***(Optional)*** |
| `dbSubnetGroupName` | ***string***| ***(Optional)*** |
| `deletionProtection` | ***bool***| ***(Optional)*** |
| `enabledCloudwatchLogsExports` | ***[]string***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `engine` | ***string***| ***(Optional)*** |
| `engineMode` | ***string***| ***(Optional)*** |
| `engineVersion` | ***string***| ***(Optional)*** |
| `finalSnapshotIdentifier` | ***string***| ***(Optional)*** |
| `globalClusterIdentifier` | ***string***| ***(Optional)*** |
| `hostedZoneID` | ***string***| ***(Optional)*** |
| `iamDatabaseAuthenticationEnabled` | ***bool***| ***(Optional)*** |
| `iamRoles` | ***[]string***| ***(Optional)*** |
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `masterUsername` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `preferredBackupWindow` | ***string***| ***(Optional)*** |
| `preferredMaintenanceWindow` | ***string***| ***(Optional)*** |
| `readerEndpoint` | ***string***| ***(Optional)*** |
| `replicationSourceIdentifier` | ***string***| ***(Optional)*** |
| `s3Import` | ***[[]RdsClusterSpecS3Import](#rdsclusterspecs3import)***| ***(Optional)*** |
| `scalingConfiguration` | ***[[]RdsClusterSpecScalingConfiguration](#rdsclusterspecscalingconfiguration)***| ***(Optional)*** |
| `skipFinalSnapshot` | ***bool***| ***(Optional)*** |
| `snapshotIdentifier` | ***string***| ***(Optional)*** |
| `sourceRegion` | ***string***| ***(Optional)*** |
| `storageEncrypted` | ***bool***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcSecurityGroupIDS` | ***[]string***| ***(Optional)*** |
## RdsClusterSpecS3Import

Appears on:[RdsClusterSpec](#rdsclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketName` | ***string***||
| `bucketPrefix` | ***string***| ***(Optional)*** |
| `ingestionRole` | ***string***||
| `sourceEngine` | ***string***||
| `sourceEngineVersion` | ***string***||
## RdsClusterSpecScalingConfiguration

Appears on:[RdsClusterSpec](#rdsclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoPause` | ***bool***| ***(Optional)*** |
| `maxCapacity` | ***int***| ***(Optional)*** |
| `minCapacity` | ***int***| ***(Optional)*** |
| `secondsUntilAutoPause` | ***int***| ***(Optional)*** |
## RdsClusterStatus

Appears on:[RdsCluster](#rdscluster)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RdsClusterSpec](#rdsclusterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `master_password` | ***string*** ||
