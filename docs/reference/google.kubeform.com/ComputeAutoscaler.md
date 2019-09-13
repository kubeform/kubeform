---
title: ComputeAutoscaler
menu:
  docs_{{ .version }}:
    identifier: computeautoscaler-google.kubeform.com
    name: ComputeAutoscaler
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeAutoscaler
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeAutoscaler` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeAutoscalerSpec](#ComputeAutoscalerSpec)***||
| `status` | ***[ComputeAutoscalerStatus](#ComputeAutoscalerStatus)***||
## ComputeAutoscalerSpec
##### (Appears on:[ComputeAutoscaler](#ComputeAutoscaler), [ComputeAutoscalerStatus](#ComputeAutoscalerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoscalingPolicy` | ***[[]ComputeAutoscalerSpecAutoscalingPolicy](#ComputeAutoscalerSpecAutoscalingPolicy)***||
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `target` | ***string***||
| `zone` | ***string***| ***(Optional)*** |
## ComputeAutoscalerSpecAutoscalingPolicy
##### (Appears on:[ComputeAutoscalerSpec](#ComputeAutoscalerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cooldownPeriod` | ***int***| ***(Optional)*** |
| `cpuUtilization` | ***[[]ComputeAutoscalerSpecAutoscalingPolicyCpuUtilization](#ComputeAutoscalerSpecAutoscalingPolicyCpuUtilization)***| ***(Optional)*** |
| `loadBalancingUtilization` | ***[[]ComputeAutoscalerSpecAutoscalingPolicyLoadBalancingUtilization](#ComputeAutoscalerSpecAutoscalingPolicyLoadBalancingUtilization)***| ***(Optional)*** |
| `maxReplicas` | ***int***||
| `metric` | ***[[]ComputeAutoscalerSpecAutoscalingPolicyMetric](#ComputeAutoscalerSpecAutoscalingPolicyMetric)***| ***(Optional)*** |
| `minReplicas` | ***int***||
## ComputeAutoscalerSpecAutoscalingPolicyCpuUtilization
##### (Appears on:[ComputeAutoscalerSpecAutoscalingPolicy](#ComputeAutoscalerSpecAutoscalingPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `target` | ***encoding/json.Number***||
## ComputeAutoscalerSpecAutoscalingPolicyLoadBalancingUtilization
##### (Appears on:[ComputeAutoscalerSpecAutoscalingPolicy](#ComputeAutoscalerSpecAutoscalingPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `target` | ***encoding/json.Number***||
## ComputeAutoscalerSpecAutoscalingPolicyMetric
##### (Appears on:[ComputeAutoscalerSpecAutoscalingPolicy](#ComputeAutoscalerSpecAutoscalingPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `target` | ***encoding/json.Number***||
| `type` | ***string***||
## ComputeAutoscalerStatus
##### (Appears on:[ComputeAutoscaler](#ComputeAutoscaler))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeAutoscalerSpec](#ComputeAutoscalerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
