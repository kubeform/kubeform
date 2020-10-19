---
title: ElasticacheReplicationGroup
menu:
  docs_{{ .version }}:
    identifier: elasticachereplicationgroup-aws.kubeform.com
    name: ElasticacheReplicationGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ElasticacheReplicationGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ElasticacheReplicationGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElasticacheReplicationGroupSpec](#elasticachereplicationgroupspec)***||
| `status` | ***[ElasticacheReplicationGroupStatus](#elasticachereplicationgroupstatus)***||
## ElasticacheReplicationGroupSpec

Appears on:[ElasticacheReplicationGroup](#elasticachereplicationgroup), [ElasticacheReplicationGroupStatus](#elasticachereplicationgroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `applyImmediately` | ***bool***| ***(Optional)*** |
| `atRestEncryptionEnabled` | ***bool***| ***(Optional)*** |
| `autoMinorVersionUpgrade` | ***bool***| ***(Optional)*** |
| `automaticFailoverEnabled` | ***bool***| ***(Optional)*** |
| `availabilityZones` | ***[]string***| ***(Optional)*** |
| `clusterMode` | ***[[]ElasticacheReplicationGroupSpecClusterMode](#elasticachereplicationgroupspecclustermode)***| ***(Optional)*** |
| `configurationEndpointAddress` | ***string***| ***(Optional)*** |
| `engine` | ***string***| ***(Optional)*** |
| `engineVersion` | ***string***| ***(Optional)*** |
| `maintenanceWindow` | ***string***| ***(Optional)*** |
| `memberClusters` | ***[]string***| ***(Optional)*** |
| `nodeType` | ***string***| ***(Optional)*** |
| `notificationTopicArn` | ***string***| ***(Optional)*** |
| `numberCacheClusters` | ***int64***| ***(Optional)*** |
| `parameterGroupName` | ***string***| ***(Optional)*** |
| `port` | ***int64***| ***(Optional)*** |
| `primaryEndpointAddress` | ***string***| ***(Optional)*** |
| `replicationGroupDescription` | ***string***||
| `replicationGroupID` | ***string***||
| `securityGroupIDS` | ***[]string***| ***(Optional)*** |
| `securityGroupNames` | ***[]string***| ***(Optional)*** |
| `snapshotArns` | ***[]string***| ***(Optional)*** |
| `snapshotName` | ***string***| ***(Optional)*** |
| `snapshotRetentionLimit` | ***int64***| ***(Optional)*** |
| `snapshotWindow` | ***string***| ***(Optional)*** |
| `subnetGroupName` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `transitEncryptionEnabled` | ***bool***| ***(Optional)*** |
## ElasticacheReplicationGroupSpecClusterMode

Appears on:[ElasticacheReplicationGroupSpec](#elasticachereplicationgroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `numNodeGroups` | ***int64***||
| `replicasPerNodeGroup` | ***int64***||
## ElasticacheReplicationGroupStatus

Appears on:[ElasticacheReplicationGroup](#elasticachereplicationgroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElasticacheReplicationGroupSpec](#elasticachereplicationgroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ElasticacheReplicationGroupStatus](#elasticachereplicationgroupstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `auth_token` | ***string*** ||
