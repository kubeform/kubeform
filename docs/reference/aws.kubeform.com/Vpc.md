---
title: Vpc
menu:
  docs_v0.0.1:
    identifier: vpc-aws.kubeform.com
    name: Vpc
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## Vpc
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Vpc` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpcSpec](#VpcSpec)***||
| `status` | ***[VpcStatus](#VpcStatus)***||
## VpcSpec
##### (Appears on:[Vpc](#Vpc), [VpcStatus](#VpcStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `assignGeneratedIpv6CIDRBlock` | ***bool***| ***(Optional)*** |
| `cidrBlock` | ***string***||
| `defaultNetworkACLID` | ***string***| ***(Optional)*** |
| `defaultRouteTableID` | ***string***| ***(Optional)*** |
| `defaultSecurityGroupID` | ***string***| ***(Optional)*** |
| `dhcpOptionsID` | ***string***| ***(Optional)*** |
| `enableClassiclink` | ***bool***| ***(Optional)*** |
| `enableClassiclinkDNSSupport` | ***bool***| ***(Optional)*** |
| `enableDNSHostnames` | ***bool***| ***(Optional)*** |
| `enableDNSSupport` | ***bool***| ***(Optional)*** |
| `instanceTenancy` | ***string***| ***(Optional)*** |
| `ipv6AssociationID` | ***string***| ***(Optional)*** |
| `ipv6CIDRBlock` | ***string***| ***(Optional)*** |
| `mainRouteTableID` | ***string***| ***(Optional)*** |
| `ownerID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## VpcStatus
##### (Appears on:[Vpc](#Vpc))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpcSpec](#VpcSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
