---
title: LbTargetGroup
menu:
  docs_{{ .version }}:
    identifier: lbtargetgroup-aws.kubeform.com
    name: LbTargetGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LbTargetGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LbTargetGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbTargetGroupSpec](#lbtargetgroupspec)***||
| `status` | ***[LbTargetGroupStatus](#lbtargetgroupstatus)***||
## LbTargetGroupSpec

Appears on:[LbTargetGroup](#lbtargetgroup), [LbTargetGroupStatus](#lbtargetgroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `arnSuffix` | ***string***| ***(Optional)*** |
| `deregistrationDelay` | ***int64***| ***(Optional)*** |
| `healthCheck` | ***[[]LbTargetGroupSpecHealthCheck](#lbtargetgroupspechealthcheck)***| ***(Optional)*** |
| `lambdaMultiValueHeadersEnabled` | ***bool***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `port` | ***int64***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** |
| `proxyProtocolV2` | ***bool***| ***(Optional)*** |
| `slowStart` | ***int64***| ***(Optional)*** |
| `stickiness` | ***[[]LbTargetGroupSpecStickiness](#lbtargetgroupspecstickiness)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `targetType` | ***string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
## LbTargetGroupSpecHealthCheck

Appears on:[LbTargetGroupSpec](#lbtargetgroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `healthyThreshold` | ***int64***| ***(Optional)*** |
| `interval` | ***int64***| ***(Optional)*** |
| `matcher` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `port` | ***string***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** |
| `timeout` | ***int64***| ***(Optional)*** |
| `unhealthyThreshold` | ***int64***| ***(Optional)*** |
## LbTargetGroupSpecStickiness

Appears on:[LbTargetGroupSpec](#lbtargetgroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cookieDuration` | ***int64***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `type` | ***string***||
## LbTargetGroupStatus

Appears on:[LbTargetGroup](#lbtargetgroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbTargetGroupSpec](#lbtargetgroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LbTargetGroupStatus](#lbtargetgroupstatus)

---
