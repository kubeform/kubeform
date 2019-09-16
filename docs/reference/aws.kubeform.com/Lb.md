---
title: Lb
menu:
  docs_{{ .version }}:
    identifier: lb-aws.kubeform.com
    name: Lb
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Lb
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Lb` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbSpec](#LbSpec)***||
| `status` | ***[LbStatus](#LbStatus)***||
## LbSpec

Appears on:[Lb](#Lb), [LbStatus](#LbStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accessLogs` | ***[[]LbSpecAccessLogs](#LbSpecAccessLogs)***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `arnSuffix` | ***string***| ***(Optional)*** |
| `dnsName` | ***string***| ***(Optional)*** |
| `enableCrossZoneLoadBalancing` | ***bool***| ***(Optional)*** |
| `enableDeletionProtection` | ***bool***| ***(Optional)*** |
| `enableHttp2` | ***bool***| ***(Optional)*** |
| `idleTimeout` | ***int***| ***(Optional)*** |
| `internal` | ***bool***| ***(Optional)*** |
| `ipAddressType` | ***string***| ***(Optional)*** |
| `loadBalancerType` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `securityGroups` | ***[]string***| ***(Optional)*** |
| `subnetMapping` | ***[[]LbSpecSubnetMapping](#LbSpecSubnetMapping)***| ***(Optional)*** |
| `subnets` | ***[]string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
| `zoneID` | ***string***| ***(Optional)*** |
## LbSpecAccessLogs

Appears on:[LbSpec](#LbSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucket` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
## LbSpecSubnetMapping

Appears on:[LbSpec](#LbSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allocationID` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***||
## LbStatus

Appears on:[Lb](#Lb)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbSpec](#LbSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
