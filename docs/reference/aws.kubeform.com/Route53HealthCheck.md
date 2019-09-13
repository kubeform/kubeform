---
title: Route53HealthCheck
menu:
  docs_{{ .version }}:
    identifier: route53healthcheck-aws.kubeform.com
    name: Route53HealthCheck
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Route53HealthCheck
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Route53HealthCheck` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Route53HealthCheckSpec](#Route53HealthCheckSpec)***||
| `status` | ***[Route53HealthCheckStatus](#Route53HealthCheckStatus)***||
## Route53HealthCheckSpec
##### (Appears on:[Route53HealthCheck](#Route53HealthCheck), [Route53HealthCheckStatus](#Route53HealthCheckStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `childHealthThreshold` | ***int***| ***(Optional)*** |
| `childHealthchecks` | ***[]string***| ***(Optional)*** |
| `cloudwatchAlarmName` | ***string***| ***(Optional)*** |
| `cloudwatchAlarmRegion` | ***string***| ***(Optional)*** |
| `enableSni` | ***bool***| ***(Optional)*** |
| `failureThreshold` | ***int***| ***(Optional)*** |
| `fqdn` | ***string***| ***(Optional)*** |
| `insufficientDataHealthStatus` | ***string***| ***(Optional)*** |
| `invertHealthcheck` | ***bool***| ***(Optional)*** |
| `ipAddress` | ***string***| ***(Optional)*** |
| `measureLatency` | ***bool***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `referenceName` | ***string***| ***(Optional)*** |
| `regions` | ***[]string***| ***(Optional)*** |
| `requestInterval` | ***int***| ***(Optional)*** |
| `resourcePath` | ***string***| ***(Optional)*** |
| `searchString` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `type` | ***string***||
## Route53HealthCheckStatus
##### (Appears on:[Route53HealthCheck](#Route53HealthCheck))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Route53HealthCheckSpec](#Route53HealthCheckSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
