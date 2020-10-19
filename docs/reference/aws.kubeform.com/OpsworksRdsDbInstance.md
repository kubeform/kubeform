---
title: OpsworksRdsDbInstance
menu:
  docs_{{ .version }}:
    identifier: opsworksrdsdbinstance-aws.kubeform.com
    name: OpsworksRdsDbInstance
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## OpsworksRdsDbInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OpsworksRdsDbInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OpsworksRdsDbInstanceSpec](#opsworksrdsdbinstancespec)***||
| `status` | ***[OpsworksRdsDbInstanceStatus](#opsworksrdsdbinstancestatus)***||
## OpsworksRdsDbInstanceSpec

Appears on:[OpsworksRdsDbInstance](#opsworksrdsdbinstance), [OpsworksRdsDbInstanceStatus](#opsworksrdsdbinstancestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `dbUser` | ***string***||
| `rdsDbInstanceArn` | ***string***||
| `stackID` | ***string***||
## OpsworksRdsDbInstanceStatus

Appears on:[OpsworksRdsDbInstance](#opsworksrdsdbinstance)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OpsworksRdsDbInstanceSpec](#opsworksrdsdbinstancespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[OpsworksRdsDbInstanceStatus](#opsworksrdsdbinstancestatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `db_password` | ***string*** ||
