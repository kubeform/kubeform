---
title: FlowLog
menu:
  docs_{{ .version }}:
    identifier: flowlog-aws.kubeform.com
    name: FlowLog
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## FlowLog
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `FlowLog` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FlowLogSpec](#flowlogspec)***||
| `status` | ***[FlowLogStatus](#flowlogstatus)***||
## FlowLogSpec

Appears on:[FlowLog](#flowlog), [FlowLogStatus](#flowlogstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `eniID` | ***string***| ***(Optional)*** |
| `iamRoleArn` | ***string***| ***(Optional)*** |
| `logDestination` | ***string***| ***(Optional)*** |
| `logDestinationType` | ***string***| ***(Optional)*** |
| `logGroupName` | ***string***| ***(Optional)*** Deprecated|
| `subnetID` | ***string***| ***(Optional)*** |
| `trafficType` | ***string***||
| `vpcID` | ***string***| ***(Optional)*** |
## FlowLogStatus

Appears on:[FlowLog](#flowlog)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FlowLogSpec](#flowlogspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
