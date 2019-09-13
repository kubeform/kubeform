---
title: NeptuneClusterInstance
menu:
  docs_{{ .version }}:
    identifier: neptuneclusterinstance-aws.kubeform.com
    name: NeptuneClusterInstance
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NeptuneClusterInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NeptuneClusterInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NeptuneClusterInstanceSpec](#NeptuneClusterInstanceSpec)***||
| `status` | ***[NeptuneClusterInstanceStatus](#NeptuneClusterInstanceStatus)***||
## NeptuneClusterInstanceSpec
##### (Appears on:[NeptuneClusterInstance](#NeptuneClusterInstance), [NeptuneClusterInstanceStatus](#NeptuneClusterInstanceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `address` | ***string***| ***(Optional)*** |
| `applyImmediately` | ***bool***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `autoMinorVersionUpgrade` | ***bool***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `clusterIdentifier` | ***string***||
| `dbiResourceID` | ***string***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `engine` | ***string***| ***(Optional)*** |
| `engineVersion` | ***string***| ***(Optional)*** |
| `identifier` | ***string***| ***(Optional)*** |
| `identifierPrefix` | ***string***| ***(Optional)*** |
| `instanceClass` | ***string***||
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `neptuneParameterGroupName` | ***string***| ***(Optional)*** |
| `neptuneSubnetGroupName` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `preferredBackupWindow` | ***string***| ***(Optional)*** |
| `preferredMaintenanceWindow` | ***string***| ***(Optional)*** |
| `promotionTier` | ***int***| ***(Optional)*** |
| `publiclyAccessible` | ***bool***| ***(Optional)*** |
| `storageEncrypted` | ***bool***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `writer` | ***bool***| ***(Optional)*** |
## NeptuneClusterInstanceStatus
##### (Appears on:[NeptuneClusterInstance](#NeptuneClusterInstance))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NeptuneClusterInstanceSpec](#NeptuneClusterInstanceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
