---
title: Elb
menu:
  docs_{{ .version }}:
    identifier: elb-aws.kubeform.com
    name: Elb
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Elb
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Elb` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElbSpec](#ElbSpec)***||
| `status` | ***[ElbStatus](#ElbStatus)***||
## ElbSpec
##### (Appears on:[Elb](#Elb), [ElbStatus](#ElbStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accessLogs` | ***[[]ElbSpecAccessLogs](#ElbSpecAccessLogs)***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `availabilityZones` | ***[]string***| ***(Optional)*** |
| `connectionDraining` | ***bool***| ***(Optional)*** |
| `connectionDrainingTimeout` | ***int***| ***(Optional)*** |
| `crossZoneLoadBalancing` | ***bool***| ***(Optional)*** |
| `dnsName` | ***string***| ***(Optional)*** |
| `healthCheck` | ***[[]ElbSpecHealthCheck](#ElbSpecHealthCheck)***| ***(Optional)*** |
| `idleTimeout` | ***int***| ***(Optional)*** |
| `instances` | ***[]string***| ***(Optional)*** |
| `internal` | ***bool***| ***(Optional)*** |
| `listener` | ***[[]ElbSpecListener](#ElbSpecListener)***||
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `securityGroups` | ***[]string***| ***(Optional)*** |
| `sourceSecurityGroup` | ***string***| ***(Optional)*** |
| `sourceSecurityGroupID` | ***string***| ***(Optional)*** |
| `subnets` | ***[]string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `zoneID` | ***string***| ***(Optional)*** |
## ElbSpecAccessLogs
##### (Appears on:[ElbSpec](#ElbSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucket` | ***string***||
| `bucketPrefix` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `interval` | ***int***| ***(Optional)*** |
## ElbSpecHealthCheck
##### (Appears on:[ElbSpec](#ElbSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `healthyThreshold` | ***int***||
| `interval` | ***int***||
| `target` | ***string***||
| `timeout` | ***int***||
| `unhealthyThreshold` | ***int***||
## ElbSpecListener
##### (Appears on:[ElbSpec](#ElbSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `instancePort` | ***int***||
| `instanceProtocol` | ***string***||
| `lbPort` | ***int***||
| `lbProtocol` | ***string***||
| `sslCertificateID` | ***string***| ***(Optional)*** |
## ElbStatus
##### (Appears on:[Elb](#Elb))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElbSpec](#ElbSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
