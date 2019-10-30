---
title: RdsClusterEndpoint
menu:
  docs_{{ .version }}:
    identifier: rdsclusterendpoint-aws.kubeform.com
    name: RdsClusterEndpoint
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RdsClusterEndpoint
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `RdsClusterEndpoint` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RdsClusterEndpointSpec](#rdsclusterendpointspec)***||
| `status` | ***[RdsClusterEndpointStatus](#rdsclusterendpointstatus)***||
## Phase(`string` alias)

Appears on:[RdsClusterEndpointStatus](#rdsclusterendpointstatus)

## RdsClusterEndpointSpec

Appears on:[RdsClusterEndpoint](#rdsclusterendpoint), [RdsClusterEndpointStatus](#rdsclusterendpointstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `clusterEndpointIdentifier` | ***string***||
| `clusterIdentifier` | ***string***||
| `customEndpointType` | ***string***||
| `endpoint` | ***string***| ***(Optional)*** |
| `excludedMembers` | ***[]string***| ***(Optional)*** |
| `staticMembers` | ***[]string***| ***(Optional)*** |
## RdsClusterEndpointStatus

Appears on:[RdsClusterEndpoint](#rdsclusterendpoint)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RdsClusterEndpointSpec](#rdsclusterendpointspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
