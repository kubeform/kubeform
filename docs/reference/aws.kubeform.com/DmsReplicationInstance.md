---
title: DmsReplicationInstance
menu:
  docs_{{ .version }}:
    identifier: dmsreplicationinstance-aws.kubeform.com
    name: DmsReplicationInstance
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DmsReplicationInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DmsReplicationInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DmsReplicationInstanceSpec](#dmsreplicationinstancespec)***||
| `status` | ***[DmsReplicationInstanceStatus](#dmsreplicationinstancestatus)***||
## DmsReplicationInstanceSpec

Appears on:[DmsReplicationInstance](#dmsreplicationinstance), [DmsReplicationInstanceStatus](#dmsreplicationinstancestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allocatedStorage` | ***int***| ***(Optional)*** |
| `applyImmediately` | ***bool***| ***(Optional)*** |
| `autoMinorVersionUpgrade` | ***bool***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `engineVersion` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `multiAz` | ***bool***| ***(Optional)*** |
| `preferredMaintenanceWindow` | ***string***| ***(Optional)*** |
| `publiclyAccessible` | ***bool***| ***(Optional)*** |
| `replicationInstanceArn` | ***string***| ***(Optional)*** |
| `replicationInstanceClass` | ***string***||
| `replicationInstanceID` | ***string***||
| `replicationInstancePrivateIPS` | ***[]string***| ***(Optional)*** |
| `replicationInstancePublicIPS` | ***[]string***| ***(Optional)*** |
| `replicationSubnetGroupID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcSecurityGroupIDS` | ***[]string***| ***(Optional)*** |
## DmsReplicationInstanceStatus

Appears on:[DmsReplicationInstance](#dmsreplicationinstance)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DmsReplicationInstanceSpec](#dmsreplicationinstancespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
