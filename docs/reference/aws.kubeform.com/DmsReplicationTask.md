---
title: DmsReplicationTask
menu:
  docs_{{ .version }}:
    identifier: dmsreplicationtask-aws.kubeform.com
    name: DmsReplicationTask
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DmsReplicationTask
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DmsReplicationTask` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DmsReplicationTaskSpec](#DmsReplicationTaskSpec)***||
| `status` | ***[DmsReplicationTaskStatus](#DmsReplicationTaskStatus)***||
## DmsReplicationTaskSpec

Appears on:[DmsReplicationTask](#DmsReplicationTask), [DmsReplicationTaskStatus](#DmsReplicationTaskStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cdcStartTime` | ***string***| ***(Optional)*** |
| `migrationType` | ***string***||
| `replicationInstanceArn` | ***string***||
| `replicationTaskArn` | ***string***| ***(Optional)*** |
| `replicationTaskID` | ***string***||
| `replicationTaskSettings` | ***string***| ***(Optional)*** |
| `sourceEndpointArn` | ***string***||
| `tableMappings` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `targetEndpointArn` | ***string***||
## DmsReplicationTaskStatus

Appears on:[DmsReplicationTask](#DmsReplicationTask)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DmsReplicationTaskSpec](#DmsReplicationTaskSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
