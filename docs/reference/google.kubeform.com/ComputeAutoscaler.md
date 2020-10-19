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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeAutoscalerSpec](#computeautoscalerspec)***||
| `status` | ***[ComputeAutoscalerStatus](#computeautoscalerstatus)***||
## ComputeAutoscalerSpec

Appears on:[ComputeAutoscaler](#computeautoscaler), [ComputeAutoscalerStatus](#computeautoscalerstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoscalingPolicy` | ***[[]ComputeAutoscalerSpecAutoscalingPolicy](#computeautoscalerspecautoscalingpolicy)***||
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `target` | ***string***||
| `zone` | ***string***| ***(Optional)*** |
## ComputeAutoscalerSpecAutoscalingPolicy

Appears on:[ComputeAutoscalerSpec](#computeautoscalerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cooldownPeriod` | ***int64***| ***(Optional)*** |
| `cpuUtilization` | ***[[]ComputeAutoscalerSpecAutoscalingPolicyCpuUtilization](#computeautoscalerspecautoscalingpolicycpuutilization)***| ***(Optional)*** |
| `loadBalancingUtilization` | ***[[]ComputeAutoscalerSpecAutoscalingPolicyLoadBalancingUtilization](#computeautoscalerspecautoscalingpolicyloadbalancingutilization)***| ***(Optional)*** |
| `maxReplicas` | ***int64***||
| `metric` | ***[[]ComputeAutoscalerSpecAutoscalingPolicyMetric](#computeautoscalerspecautoscalingpolicymetric)***| ***(Optional)*** |
| `minReplicas` | ***int64***||
## ComputeAutoscalerSpecAutoscalingPolicyCpuUtilization

Appears on:[ComputeAutoscalerSpecAutoscalingPolicy](#computeautoscalerspecautoscalingpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `target` | ***float64***||
## ComputeAutoscalerSpecAutoscalingPolicyLoadBalancingUtilization

Appears on:[ComputeAutoscalerSpecAutoscalingPolicy](#computeautoscalerspecautoscalingpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `target` | ***float64***||
## ComputeAutoscalerSpecAutoscalingPolicyMetric

Appears on:[ComputeAutoscalerSpecAutoscalingPolicy](#computeautoscalerspecautoscalingpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `target` | ***float64***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## ComputeAutoscalerStatus

Appears on:[ComputeAutoscaler](#computeautoscaler)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeAutoscalerSpec](#computeautoscalerspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeAutoscalerStatus](#computeautoscalerstatus)

---
