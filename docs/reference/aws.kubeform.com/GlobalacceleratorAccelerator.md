---
title: GlobalacceleratorAccelerator
menu:
  docs_v0.0.1:
    identifier: globalacceleratoraccelerator-aws.kubeform.com
    name: GlobalacceleratorAccelerator
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## GlobalacceleratorAccelerator
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GlobalacceleratorAccelerator` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GlobalacceleratorAcceleratorSpec](#GlobalacceleratorAcceleratorSpec)***||
| `status` | ***[GlobalacceleratorAcceleratorStatus](#GlobalacceleratorAcceleratorStatus)***||
## GlobalacceleratorAcceleratorSpec
##### (Appears on:[GlobalacceleratorAccelerator](#GlobalacceleratorAccelerator), [GlobalacceleratorAcceleratorStatus](#GlobalacceleratorAcceleratorStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `attributes` | ***[[]GlobalacceleratorAcceleratorSpecAttributes](#GlobalacceleratorAcceleratorSpecAttributes)***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `ipAddressType` | ***string***| ***(Optional)*** |
| `ipSets` | ***[[]GlobalacceleratorAcceleratorSpecIpSets](#GlobalacceleratorAcceleratorSpecIpSets)***| ***(Optional)*** |
| `name` | ***string***||
## GlobalacceleratorAcceleratorSpecAttributes
##### (Appears on:[GlobalacceleratorAcceleratorSpec](#GlobalacceleratorAcceleratorSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `flowLogsEnabled` | ***bool***| ***(Optional)*** |
| `flowLogsS3Bucket` | ***string***| ***(Optional)*** |
| `flowLogsS3Prefix` | ***string***| ***(Optional)*** |
## GlobalacceleratorAcceleratorSpecIpSets
##### (Appears on:[GlobalacceleratorAcceleratorSpec](#GlobalacceleratorAcceleratorSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ipAddresses` | ***[]string***| ***(Optional)*** |
| `ipFamily` | ***string***| ***(Optional)*** |
## GlobalacceleratorAcceleratorStatus
##### (Appears on:[GlobalacceleratorAccelerator](#GlobalacceleratorAccelerator))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlobalacceleratorAcceleratorSpec](#GlobalacceleratorAcceleratorSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
