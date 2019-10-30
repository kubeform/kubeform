---
title: DxConnection
menu:
  docs_{{ .version }}:
    identifier: dxconnection-aws.kubeform.com
    name: DxConnection
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DxConnection
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DxConnection` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DxConnectionSpec](#dxconnectionspec)***||
| `status` | ***[DxConnectionStatus](#dxconnectionstatus)***||
## DxConnectionSpec

Appears on:[DxConnection](#dxconnection), [DxConnectionStatus](#dxconnectionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `awsDevice` | ***string***| ***(Optional)*** |
| `bandwidth` | ***string***||
| `hasLogicalRedundancy` | ***string***| ***(Optional)*** |
| `jumboFrameCapable` | ***bool***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## DxConnectionStatus

Appears on:[DxConnection](#dxconnection)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DxConnectionSpec](#dxconnectionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DxConnectionStatus](#dxconnectionstatus)

---
