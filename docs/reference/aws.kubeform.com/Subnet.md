---
title: Subnet
menu:
  docs_v0.0.1:
    identifier: subnet-aws.kubeform.com
    name: Subnet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## Subnet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Subnet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SubnetSpec](#SubnetSpec)***||
| `status` | ***[SubnetStatus](#SubnetStatus)***||
## SubnetSpec
##### (Appears on:[Subnet](#Subnet), [SubnetStatus](#SubnetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `assignIpv6AddressOnCreation` | ***bool***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `availabilityZoneID` | ***string***| ***(Optional)*** |
| `cidrBlock` | ***string***||
| `ipv6CIDRBlock` | ***string***| ***(Optional)*** |
| `ipv6CIDRBlockAssociationID` | ***string***| ***(Optional)*** |
| `mapPublicIPOnLaunch` | ***bool***| ***(Optional)*** |
| `ownerID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***||
## SubnetStatus
##### (Appears on:[Subnet](#Subnet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SubnetSpec](#SubnetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
