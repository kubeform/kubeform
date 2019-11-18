---
title: AlbTargetGroup
menu:
  docs_{{ .version }}:
    identifier: albtargetgroup-aws.kubeform.com
    name: AlbTargetGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AlbTargetGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AlbTargetGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AlbTargetGroupSpec](#albtargetgroupspec)***||
| `status` | ***[AlbTargetGroupStatus](#albtargetgroupstatus)***||
## AlbTargetGroupSpec

Appears on:[AlbTargetGroup](#albtargetgroup), [AlbTargetGroupStatus](#albtargetgroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `arnSuffix` | ***string***| ***(Optional)*** |
| `deregistrationDelay` | ***int64***| ***(Optional)*** |
| `healthCheck` | ***[[]AlbTargetGroupSpecHealthCheck](#albtargetgroupspechealthcheck)***| ***(Optional)*** |
| `lambdaMultiValueHeadersEnabled` | ***bool***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `port` | ***int64***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** |
| `proxyProtocolV2` | ***bool***| ***(Optional)*** |
| `slowStart` | ***int64***| ***(Optional)*** |
| `stickiness` | ***[[]AlbTargetGroupSpecStickiness](#albtargetgroupspecstickiness)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `targetType` | ***string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
## AlbTargetGroupSpecHealthCheck

Appears on:[AlbTargetGroupSpec](#albtargetgroupspec)

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
## AlbTargetGroupSpecStickiness

Appears on:[AlbTargetGroupSpec](#albtargetgroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cookieDuration` | ***int64***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `type` | ***string***||
## AlbTargetGroupStatus

Appears on:[AlbTargetGroup](#albtargetgroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AlbTargetGroupSpec](#albtargetgroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AlbTargetGroupStatus](#albtargetgroupstatus)

---
