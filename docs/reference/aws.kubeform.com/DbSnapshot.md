---
title: DbSnapshot
menu:
  docs_{{ .version }}:
    identifier: dbsnapshot-aws.kubeform.com
    name: DbSnapshot
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DbSnapshot
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DbSnapshot` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DbSnapshotSpec](#DbSnapshotSpec)***||
| `status` | ***[DbSnapshotStatus](#DbSnapshotStatus)***||
## DbSnapshotSpec
##### (Appears on:[DbSnapshot](#DbSnapshot), [DbSnapshotStatus](#DbSnapshotStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allocatedStorage` | ***int***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `dbInstanceIdentifier` | ***string***||
| `dbSnapshotArn` | ***string***| ***(Optional)*** |
| `dbSnapshotIdentifier` | ***string***||
| `encrypted` | ***bool***| ***(Optional)*** |
| `engine` | ***string***| ***(Optional)*** |
| `engineVersion` | ***string***| ***(Optional)*** |
| `iops` | ***int***| ***(Optional)*** |
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `licenseModel` | ***string***| ***(Optional)*** |
| `optionGroupName` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `snapshotType` | ***string***| ***(Optional)*** |
| `sourceDbSnapshotIdentifier` | ***string***| ***(Optional)*** |
| `sourceRegion` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `storageType` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
## DbSnapshotStatus
##### (Appears on:[DbSnapshot](#DbSnapshot))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DbSnapshotSpec](#DbSnapshotSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
