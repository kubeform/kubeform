---
title: RdsClusterInstance
menu:
  docs_{{ .version }}:
    identifier: rdsclusterinstance-aws.kubeform.com
    name: RdsClusterInstance
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RdsClusterInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `RdsClusterInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RdsClusterInstanceSpec](#rdsclusterinstancespec)***||
| `status` | ***[RdsClusterInstanceStatus](#rdsclusterinstancestatus)***||
## Phase(`string` alias)

Appears on:[RdsClusterInstanceStatus](#rdsclusterinstancestatus)

## RdsClusterInstanceSpec

Appears on:[RdsClusterInstance](#rdsclusterinstance), [RdsClusterInstanceStatus](#rdsclusterinstancestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `applyImmediately` | ***bool***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `autoMinorVersionUpgrade` | ***bool***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `clusterIdentifier` | ***string***||
| `copyTagsToSnapshot` | ***bool***| ***(Optional)*** |
| `dbParameterGroupName` | ***string***| ***(Optional)*** |
| `dbSubnetGroupName` | ***string***| ***(Optional)*** |
| `dbiResourceID` | ***string***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `engine` | ***string***| ***(Optional)*** |
| `engineVersion` | ***string***| ***(Optional)*** |
| `identifier` | ***string***| ***(Optional)*** |
| `identifierPrefix` | ***string***| ***(Optional)*** |
| `instanceClass` | ***string***||
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `monitoringInterval` | ***int***| ***(Optional)*** |
| `monitoringRoleArn` | ***string***| ***(Optional)*** |
| `performanceInsightsEnabled` | ***bool***| ***(Optional)*** |
| `performanceInsightsKmsKeyID` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `preferredBackupWindow` | ***string***| ***(Optional)*** |
| `preferredMaintenanceWindow` | ***string***| ***(Optional)*** |
| `promotionTier` | ***int***| ***(Optional)*** |
| `publiclyAccessible` | ***bool***| ***(Optional)*** |
| `storageEncrypted` | ***bool***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `writer` | ***bool***| ***(Optional)*** |
## RdsClusterInstanceStatus

Appears on:[RdsClusterInstance](#rdsclusterinstance)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RdsClusterInstanceSpec](#rdsclusterinstancespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
