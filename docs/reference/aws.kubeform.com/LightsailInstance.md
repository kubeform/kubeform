---
title: LightsailInstance
menu:
  docs_{{ .version }}:
    identifier: lightsailinstance-aws.kubeform.com
    name: LightsailInstance
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LightsailInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LightsailInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LightsailInstanceSpec](#LightsailInstanceSpec)***||
| `status` | ***[LightsailInstanceStatus](#LightsailInstanceStatus)***||
## LightsailInstanceSpec

Appears on:[LightsailInstance](#LightsailInstance), [LightsailInstanceStatus](#LightsailInstanceStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `availabilityZone` | ***string***||
| `blueprintID` | ***string***||
| `bundleID` | ***string***||
| `cpuCount` | ***int***| ***(Optional)*** |
| `createdAt` | ***string***| ***(Optional)*** |
| `ipv6Address` | ***string***| ***(Optional)*** |
| `isStaticIP` | ***bool***| ***(Optional)*** |
| `keyPairName` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `privateIPAddress` | ***string***| ***(Optional)*** |
| `publicIPAddress` | ***string***| ***(Optional)*** |
| `ramSize` | ***int***| ***(Optional)*** |
| `userData` | ***string***| ***(Optional)*** |
| `username` | ***string***| ***(Optional)*** |
## LightsailInstanceStatus

Appears on:[LightsailInstance](#LightsailInstance)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LightsailInstanceSpec](#LightsailInstanceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
