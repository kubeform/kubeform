---
title: AutoscalingPolicy
menu:
  docs_{{ .version }}:
    identifier: autoscalingpolicy-aws.kubeform.com
    name: AutoscalingPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AutoscalingPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AutoscalingPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutoscalingPolicySpec](#autoscalingpolicyspec)***||
| `status` | ***[AutoscalingPolicyStatus](#autoscalingpolicystatus)***||
## AutoscalingPolicySpec

Appears on:[AutoscalingPolicy](#autoscalingpolicy), [AutoscalingPolicyStatus](#autoscalingpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `adjustmentType` | ***string***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `autoscalingGroupName` | ***string***||
| `cooldown` | ***int64***| ***(Optional)*** |
| `estimatedInstanceWarmup` | ***int64***| ***(Optional)*** |
| `metricAggregationType` | ***string***| ***(Optional)*** |
| `minAdjustmentMagnitude` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
| `policyType` | ***string***| ***(Optional)*** |
| `scalingAdjustment` | ***int64***| ***(Optional)*** |
| `stepAdjustment` | ***[[]AutoscalingPolicySpecStepAdjustment](#autoscalingpolicyspecstepadjustment)***| ***(Optional)*** |
| `targetTrackingConfiguration` | ***[[]AutoscalingPolicySpecTargetTrackingConfiguration](#autoscalingpolicyspectargettrackingconfiguration)***| ***(Optional)*** |
## AutoscalingPolicySpecStepAdjustment

Appears on:[AutoscalingPolicySpec](#autoscalingpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `metricIntervalLowerBound` | ***string***| ***(Optional)*** |
| `metricIntervalUpperBound` | ***string***| ***(Optional)*** |
| `scalingAdjustment` | ***int64***||
## AutoscalingPolicySpecTargetTrackingConfiguration

Appears on:[AutoscalingPolicySpec](#autoscalingpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `customizedMetricSpecification` | ***[[]AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification](#autoscalingpolicyspectargettrackingconfigurationcustomizedmetricspecification)***| ***(Optional)*** |
| `disableScaleIn` | ***bool***| ***(Optional)*** |
| `predefinedMetricSpecification` | ***[[]AutoscalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification](#autoscalingpolicyspectargettrackingconfigurationpredefinedmetricspecification)***| ***(Optional)*** |
| `targetValue` | ***float64***||
## AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification

Appears on:[AutoscalingPolicySpecTargetTrackingConfiguration](#autoscalingpolicyspectargettrackingconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `metricDimension` | ***[[]AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension](#autoscalingpolicyspectargettrackingconfigurationcustomizedmetricspecificationmetricdimension)***| ***(Optional)*** |
| `metricName` | ***string***||
| `namespace` | ***string***||
| `statistic` | ***string***||
| `unit` | ***string***| ***(Optional)*** |
## AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension

Appears on:[AutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification](#autoscalingpolicyspectargettrackingconfigurationcustomizedmetricspecification)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `value` | ***string***||
## AutoscalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification

Appears on:[AutoscalingPolicySpecTargetTrackingConfiguration](#autoscalingpolicyspectargettrackingconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `predefinedMetricType` | ***string***||
| `resourceLabel` | ***string***| ***(Optional)*** |
## AutoscalingPolicyStatus

Appears on:[AutoscalingPolicy](#autoscalingpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutoscalingPolicySpec](#autoscalingpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AutoscalingPolicyStatus](#autoscalingpolicystatus)

---
