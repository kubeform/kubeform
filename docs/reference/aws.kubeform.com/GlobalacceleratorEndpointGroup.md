---
title: GlobalacceleratorEndpointGroup
menu:
  docs_{{ .version }}:
    identifier: globalacceleratorendpointgroup-aws.kubeform.com
    name: GlobalacceleratorEndpointGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GlobalacceleratorEndpointGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GlobalacceleratorEndpointGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GlobalacceleratorEndpointGroupSpec](#globalacceleratorendpointgroupspec)***||
| `status` | ***[GlobalacceleratorEndpointGroupStatus](#globalacceleratorendpointgroupstatus)***||
## GlobalacceleratorEndpointGroupSpec


Appears on:[GlobalacceleratorEndpointGroup](#globalacceleratorendpointgroup), [GlobalacceleratorEndpointGroupStatus](#globalacceleratorendpointgroupstatus)


| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `endpointConfiguration` | ***[[]GlobalacceleratorEndpointGroupSpecEndpointConfiguration](#globalacceleratorendpointgroupspecendpointconfiguration)***| ***(Optional)*** |
| `endpointGroupRegion` | ***string***| ***(Optional)*** |
| `healthCheckIntervalSeconds` | ***int***| ***(Optional)*** |
| `healthCheckPath` | ***string***| ***(Optional)*** |
| `healthCheckPort` | ***int***| ***(Optional)*** |
| `healthCheckProtocol` | ***string***| ***(Optional)*** |
| `listenerArn` | ***string***||
| `thresholdCount` | ***int***| ***(Optional)*** |
| `trafficDialPercentage` | ***encoding/json.Number***| ***(Optional)*** |
## GlobalacceleratorEndpointGroupSpecEndpointConfiguration


Appears on:[GlobalacceleratorEndpointGroupSpec](#globalacceleratorendpointgroupspec)


| Field | Type | Description |
| ------ | ----- | ----------- |
| `endpointID` | ***string***| ***(Optional)*** |
| `weight` | ***int***| ***(Optional)*** |
## GlobalacceleratorEndpointGroupStatus


Appears on:[GlobalacceleratorEndpointGroup](#globalacceleratorendpointgroup)


| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlobalacceleratorEndpointGroupSpec](#globalacceleratorendpointgroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
