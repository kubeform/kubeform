---
title: VpcEndpointService
menu:
  docs_{{ .version }}:
    identifier: vpcendpointservice-aws.kubeform.com
    name: VpcEndpointService
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VpcEndpointService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `VpcEndpointService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpcEndpointServiceSpec](#vpcendpointservicespec)***||
| `status` | ***[VpcEndpointServiceStatus](#vpcendpointservicestatus)***||
## Phase(`string` alias)

Appears on:[VpcEndpointServiceStatus](#vpcendpointservicestatus)

## VpcEndpointServiceSpec

Appears on:[VpcEndpointService](#vpcendpointservice), [VpcEndpointServiceStatus](#vpcendpointservicestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `acceptanceRequired` | ***bool***||
| `allowedPrincipals` | ***[]string***| ***(Optional)*** |
| `availabilityZones` | ***[]string***| ***(Optional)*** |
| `baseEndpointDNSNames` | ***[]string***| ***(Optional)*** |
| `networkLoadBalancerArns` | ***[]string***||
| `privateDNSName` | ***string***| ***(Optional)*** |
| `serviceName` | ***string***| ***(Optional)*** |
| `serviceType` | ***string***| ***(Optional)*** |
| `state` | ***string***| ***(Optional)*** |
## VpcEndpointServiceStatus

Appears on:[VpcEndpointService](#vpcendpointservice)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpcEndpointServiceSpec](#vpcendpointservicespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
