---
title: OpsworksRdsDbInstance
menu:
  docs_v0.0.1:
    identifier: opsworksrdsdbinstance-aws.kubeform.com
    name: OpsworksRdsDbInstance
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## OpsworksRdsDbInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OpsworksRdsDbInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OpsworksRdsDbInstanceSpec](#OpsworksRdsDbInstanceSpec)***||
| `status` | ***[OpsworksRdsDbInstanceStatus](#OpsworksRdsDbInstanceStatus)***||
## OpsworksRdsDbInstanceSpec
##### (Appears on:[OpsworksRdsDbInstance](#OpsworksRdsDbInstance), [OpsworksRdsDbInstanceStatus](#OpsworksRdsDbInstanceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `dbUser` | ***string***||
| `rdsDbInstanceArn` | ***string***||
| `stackID` | ***string***||
## OpsworksRdsDbInstanceStatus
##### (Appears on:[OpsworksRdsDbInstance](#OpsworksRdsDbInstance))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OpsworksRdsDbInstanceSpec](#OpsworksRdsDbInstanceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `db_password` | ***string*** ||
