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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DbSnapshotSpec](#dbsnapshotspec)***||
| `status` | ***[DbSnapshotStatus](#dbsnapshotstatus)***||
## DbSnapshotSpec

Appears on:[DbSnapshot](#dbsnapshot), [DbSnapshotStatus](#dbsnapshotstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allocatedStorage` | ***int64***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `dbInstanceIdentifier` | ***string***||
| `dbSnapshotArn` | ***string***| ***(Optional)*** |
| `dbSnapshotIdentifier` | ***string***||
| `encrypted` | ***bool***| ***(Optional)*** |
| `engine` | ***string***| ***(Optional)*** |
| `engineVersion` | ***string***| ***(Optional)*** |
| `iops` | ***int64***| ***(Optional)*** |
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `licenseModel` | ***string***| ***(Optional)*** |
| `optionGroupName` | ***string***| ***(Optional)*** |
| `port` | ***int64***| ***(Optional)*** |
| `snapshotType` | ***string***| ***(Optional)*** |
| `sourceDbSnapshotIdentifier` | ***string***| ***(Optional)*** |
| `sourceRegion` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `storageType` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
## DbSnapshotStatus

Appears on:[DbSnapshot](#dbsnapshot)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DbSnapshotSpec](#dbsnapshotspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DbSnapshotStatus](#dbsnapshotstatus)

---
