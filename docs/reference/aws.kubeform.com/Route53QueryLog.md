---
title: Route53QueryLog
menu:
  docs_{{ .version }}:
    identifier: route53querylog-aws.kubeform.com
    name: Route53QueryLog
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Route53QueryLog
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Route53QueryLog` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Route53QueryLogSpec](#route53querylogspec)***||
| `status` | ***[Route53QueryLogStatus](#route53querylogstatus)***||
## Route53QueryLogSpec

Appears on:[Route53QueryLog](#route53querylog), [Route53QueryLogStatus](#route53querylogstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cloudwatchLogGroupArn` | ***string***||
| `zoneID` | ***string***||
## Route53QueryLogStatus

Appears on:[Route53QueryLog](#route53querylog)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Route53QueryLogSpec](#route53querylogspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
