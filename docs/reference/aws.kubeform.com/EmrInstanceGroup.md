---
title: EmrInstanceGroup
menu:
  docs_{{ .version }}:
    identifier: emrinstancegroup-aws.kubeform.com
    name: EmrInstanceGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## EmrInstanceGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `EmrInstanceGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EmrInstanceGroupSpec](#EmrInstanceGroupSpec)***||
| `status` | ***[EmrInstanceGroupStatus](#EmrInstanceGroupStatus)***||
## EmrInstanceGroupSpec
##### (Appears on:[EmrInstanceGroup](#EmrInstanceGroup), [EmrInstanceGroupStatus](#EmrInstanceGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `clusterID` | ***string***||
| `ebsConfig` | ***[[]EmrInstanceGroupSpecEbsConfig](#EmrInstanceGroupSpecEbsConfig)***| ***(Optional)*** |
| `ebsOptimized` | ***bool***| ***(Optional)*** |
| `instanceCount` | ***int***| ***(Optional)*** |
| `instanceType` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `runningInstanceCount` | ***int***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
## EmrInstanceGroupSpecEbsConfig
##### (Appears on:[EmrInstanceGroupSpec](#EmrInstanceGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `iops` | ***int***| ***(Optional)*** |
| `size` | ***int***||
| `type` | ***string***||
| `volumesPerInstance` | ***int***| ***(Optional)*** |
## EmrInstanceGroupStatus
##### (Appears on:[EmrInstanceGroup](#EmrInstanceGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EmrInstanceGroupSpec](#EmrInstanceGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
