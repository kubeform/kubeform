---
title: AutoscalingPolicy
menu:
  docs_v0.0.1:
    identifier: autoscalingpolicy-aws.kubeform.com
    name: AutoscalingPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## AutoscalingPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AutoscalingPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutoscalingPolicySpec](#AutoscalingPolicySpec)***||
| `status` | ***[AutoscalingPolicyStatus](#AutoscalingPolicyStatus)***||
## AutoscalingPolicySpec
##### (Appears on:[AutoscalingPolicy](#AutoscalingPolicy), [AutoscalingPolicyStatus](#AutoscalingPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `adjustmentType` | ***string***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `autoscalingGroupName` | ***string***||
| `cooldown` | ***int***| ***(Optional)*** |
| `estimatedInstanceWarmup` | ***int***| ***(Optional)*** |
| `metricAggregationType` | ***string***| ***(Optional)*** |
| `minAdjustmentMagnitude` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `policyType` | ***string***| ***(Optional)*** |
| `scalingAdjustment` | ***int***| ***(Optional)*** |
| `stepAdjustment` | ***[[]AutoscalingPolicySpecStepAdjustment](#AutoscalingPolicySpecStepAdjustment)***| ***(Optional)*** |
| `targetTrackingConfiguration` | ***[[]AutoscalingPolicySpecTargetTrackingConfiguration](#AutoscalingPolicySpecTargetTrackingConfiguration)***| ***(Optional)*** |
## AutoscalingPolicySpecStepAdjustment
##### (Appears on:[AutoscalingPolicySpec](#AutoscalingPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `metricIntervalLowerBound` | ***string***| ***(Optional)*** |
| `metricIntervalUpperBound` | ***string***| ***(Optional)*** |
| `scalingAdjustment` | ***int***||
## AutoscalingPolicySpecTargetTrackingConfiguration
##### (Appears on:[AutoscalingPolicySpec](#AutoscalingPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `customizedMetricSpecification` | ***[[]AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification](#AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification)***| ***(Optional)*** |
| `disableScaleIn` | ***bool***| ***(Optional)*** |
| `predefinedMetricSpecification` | ***[[]AutoscalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification](#AutoscalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification)***| ***(Optional)*** |
| `targetValue` | ***encoding/json.Number***||
## AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification
##### (Appears on:[AutoscalingPolicySpecTargetTrackingConfiguration](#AutoscalingPolicySpecTargetTrackingConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `metricDimension` | ***[[]AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension](#AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension)***| ***(Optional)*** |
| `metricName` | ***string***||
| `namespace` | ***string***||
| `statistic` | ***string***||
| `unit` | ***string***| ***(Optional)*** |
## AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension
##### (Appears on:[AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification](#AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `value` | ***string***||
## AutoscalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification
##### (Appears on:[AutoscalingPolicySpecTargetTrackingConfiguration](#AutoscalingPolicySpecTargetTrackingConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `predefinedMetricType` | ***string***||
| `resourceLabel` | ***string***| ***(Optional)*** |
## AutoscalingPolicyStatus
##### (Appears on:[AutoscalingPolicy](#AutoscalingPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutoscalingPolicySpec](#AutoscalingPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
