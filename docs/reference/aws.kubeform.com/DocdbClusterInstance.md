---
title: DocdbClusterInstance
menu:
  docs_{{ .version }}:
    identifier: docdbclusterinstance-aws.kubeform.com
    name: DocdbClusterInstance
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DocdbClusterInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DocdbClusterInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DocdbClusterInstanceSpec](#docdbclusterinstancespec)***||
| `status` | ***[DocdbClusterInstanceStatus](#docdbclusterinstancestatus)***||
## DocdbClusterInstanceSpec

Appears on:[DocdbClusterInstance](#docdbclusterinstance), [DocdbClusterInstanceStatus](#docdbclusterinstancestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `applyImmediately` | ***bool***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `autoMinorVersionUpgrade` | ***bool***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `clusterIdentifier` | ***string***||
| `dbSubnetGroupName` | ***string***| ***(Optional)*** |
| `dbiResourceID` | ***string***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `engine` | ***string***| ***(Optional)*** |
| `engineVersion` | ***string***| ***(Optional)*** |
| `identifier` | ***string***| ***(Optional)*** |
| `identifierPrefix` | ***string***| ***(Optional)*** |
| `instanceClass` | ***string***||
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `preferredBackupWindow` | ***string***| ***(Optional)*** |
| `preferredMaintenanceWindow` | ***string***| ***(Optional)*** |
| `promotionTier` | ***int***| ***(Optional)*** |
| `publiclyAccessible` | ***bool***| ***(Optional)*** |
| `storageEncrypted` | ***bool***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `writer` | ***bool***| ***(Optional)*** |
## DocdbClusterInstanceStatus

Appears on:[DocdbClusterInstance](#docdbclusterinstance)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DocdbClusterInstanceSpec](#docdbclusterinstancespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DocdbClusterInstanceStatus](#docdbclusterinstancestatus)

---
