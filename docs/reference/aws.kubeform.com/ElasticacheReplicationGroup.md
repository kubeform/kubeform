---
title: ElasticacheReplicationGroup
menu:
  docs_v0.0.1:
    identifier: elasticachereplicationgroup-aws.kubeform.com
    name: ElasticacheReplicationGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ElasticacheReplicationGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ElasticacheReplicationGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElasticacheReplicationGroupSpec](#ElasticacheReplicationGroupSpec)***||
| `status` | ***[ElasticacheReplicationGroupStatus](#ElasticacheReplicationGroupStatus)***||
## ElasticacheReplicationGroupSpec
##### (Appears on:[ElasticacheReplicationGroup](#ElasticacheReplicationGroup), [ElasticacheReplicationGroupStatus](#ElasticacheReplicationGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `applyImmediately` | ***bool***| ***(Optional)*** |
| `atRestEncryptionEnabled` | ***bool***| ***(Optional)*** |
| `autoMinorVersionUpgrade` | ***bool***| ***(Optional)*** |
| `automaticFailoverEnabled` | ***bool***| ***(Optional)*** |
| `availabilityZones` | ***[]string***| ***(Optional)*** |
| `clusterMode` | ***[[]ElasticacheReplicationGroupSpecClusterMode](#ElasticacheReplicationGroupSpecClusterMode)***| ***(Optional)*** |
| `configurationEndpointAddress` | ***string***| ***(Optional)*** |
| `engine` | ***string***| ***(Optional)*** |
| `engineVersion` | ***string***| ***(Optional)*** |
| `maintenanceWindow` | ***string***| ***(Optional)*** |
| `memberClusters` | ***[]string***| ***(Optional)*** |
| `nodeType` | ***string***| ***(Optional)*** |
| `notificationTopicArn` | ***string***| ***(Optional)*** |
| `numberCacheClusters` | ***int***| ***(Optional)*** |
| `parameterGroupName` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `primaryEndpointAddress` | ***string***| ***(Optional)*** |
| `replicationGroupDescription` | ***string***||
| `replicationGroupID` | ***string***||
| `securityGroupIDS` | ***[]string***| ***(Optional)*** |
| `securityGroupNames` | ***[]string***| ***(Optional)*** |
| `snapshotArns` | ***[]string***| ***(Optional)*** |
| `snapshotName` | ***string***| ***(Optional)*** |
| `snapshotRetentionLimit` | ***int***| ***(Optional)*** |
| `snapshotWindow` | ***string***| ***(Optional)*** |
| `subnetGroupName` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `transitEncryptionEnabled` | ***bool***| ***(Optional)*** |
## ElasticacheReplicationGroupSpecClusterMode
##### (Appears on:[ElasticacheReplicationGroupSpec](#ElasticacheReplicationGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `numNodeGroups` | ***int***||
| `replicasPerNodeGroup` | ***int***||
## ElasticacheReplicationGroupStatus
##### (Appears on:[ElasticacheReplicationGroup](#ElasticacheReplicationGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElasticacheReplicationGroupSpec](#ElasticacheReplicationGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `auth_token` | ***string*** ||
