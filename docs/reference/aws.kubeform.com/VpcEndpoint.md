---
title: VpcEndpoint
menu:
  docs_v0.0.1:
    identifier: vpcendpoint-aws.kubeform.com
    name: VpcEndpoint
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## VpcEndpoint
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `VpcEndpoint` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpcEndpointSpec](#VpcEndpointSpec)***||
| `status` | ***[VpcEndpointStatus](#VpcEndpointStatus)***||
## VpcEndpointSpec
##### (Appears on:[VpcEndpoint](#VpcEndpoint), [VpcEndpointStatus](#VpcEndpointStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoAccept` | ***bool***| ***(Optional)*** |
| `cidrBlocks` | ***[]string***| ***(Optional)*** |
| `dnsEntry` | ***[[]VpcEndpointSpecDnsEntry](#VpcEndpointSpecDnsEntry)***| ***(Optional)*** |
| `networkInterfaceIDS` | ***[]string***| ***(Optional)*** |
| `policy` | ***string***| ***(Optional)*** |
| `prefixListID` | ***string***| ***(Optional)*** |
| `privateDNSEnabled` | ***bool***| ***(Optional)*** |
| `routeTableIDS` | ***[]string***| ***(Optional)*** |
| `securityGroupIDS` | ***[]string***| ***(Optional)*** |
| `serviceName` | ***string***||
| `state` | ***string***| ***(Optional)*** |
| `subnetIDS` | ***[]string***| ***(Optional)*** |
| `vpcEndpointType` | ***string***| ***(Optional)*** |
| `vpcID` | ***string***||
## VpcEndpointSpecDnsEntry
##### (Appears on:[VpcEndpointSpec](#VpcEndpointSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `dnsName` | ***string***| ***(Optional)*** |
| `hostedZoneID` | ***string***| ***(Optional)*** |
## VpcEndpointStatus
##### (Appears on:[VpcEndpoint](#VpcEndpoint))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpcEndpointSpec](#VpcEndpointSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
