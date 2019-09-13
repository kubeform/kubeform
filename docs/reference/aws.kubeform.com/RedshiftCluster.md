---
title: RedshiftCluster
menu:
  docs_{{ .version }}:
    identifier: redshiftcluster-aws.kubeform.com
    name: RedshiftCluster
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RedshiftCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `RedshiftCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RedshiftClusterSpec](#RedshiftClusterSpec)***||
| `status` | ***[RedshiftClusterStatus](#RedshiftClusterStatus)***||
## RedshiftClusterSpec
##### (Appears on:[RedshiftCluster](#RedshiftCluster), [RedshiftClusterStatus](#RedshiftClusterStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `allowVersionUpgrade` | ***bool***| ***(Optional)*** |
| `automatedSnapshotRetentionPeriod` | ***int***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `clusterIdentifier` | ***string***||
| `clusterParameterGroupName` | ***string***| ***(Optional)*** |
| `clusterPublicKey` | ***string***| ***(Optional)*** |
| `clusterRevisionNumber` | ***string***| ***(Optional)*** |
| `clusterSecurityGroups` | ***[]string***| ***(Optional)*** |
| `clusterSubnetGroupName` | ***string***| ***(Optional)*** |
| `clusterType` | ***string***| ***(Optional)*** |
| `clusterVersion` | ***string***| ***(Optional)*** |
| `databaseName` | ***string***| ***(Optional)*** |
| `dnsName` | ***string***| ***(Optional)*** |
| `elasticIP` | ***string***| ***(Optional)*** |
| `encrypted` | ***bool***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `enhancedVpcRouting` | ***bool***| ***(Optional)*** |
| `finalSnapshotIdentifier` | ***string***| ***(Optional)*** |
| `iamRoles` | ***[]string***| ***(Optional)*** |
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `logging` | ***[[]RedshiftClusterSpecLogging](#RedshiftClusterSpecLogging)***| ***(Optional)*** |
| `masterUsername` | ***string***| ***(Optional)*** |
| `nodeType` | ***string***||
| `numberOfNodes` | ***int***| ***(Optional)*** |
| `ownerAccount` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `preferredMaintenanceWindow` | ***string***| ***(Optional)*** |
| `publiclyAccessible` | ***bool***| ***(Optional)*** |
| `skipFinalSnapshot` | ***bool***| ***(Optional)*** |
| `snapshotClusterIdentifier` | ***string***| ***(Optional)*** |
| `snapshotCopy` | ***[[]RedshiftClusterSpecSnapshotCopy](#RedshiftClusterSpecSnapshotCopy)***| ***(Optional)*** |
| `snapshotIdentifier` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcSecurityGroupIDS` | ***[]string***| ***(Optional)*** |
## RedshiftClusterSpecLogging
##### (Appears on:[RedshiftClusterSpec](#RedshiftClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketName` | ***string***| ***(Optional)*** |
| `enable` | ***bool***||
| `s3KeyPrefix` | ***string***| ***(Optional)*** |
## RedshiftClusterSpecSnapshotCopy
##### (Appears on:[RedshiftClusterSpec](#RedshiftClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `destinationRegion` | ***string***||
| `grantName` | ***string***| ***(Optional)*** |
| `retentionPeriod` | ***int***| ***(Optional)*** |
## RedshiftClusterStatus
##### (Appears on:[RedshiftCluster](#RedshiftCluster))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RedshiftClusterSpec](#RedshiftClusterSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `master_password` | ***string*** ||
