---
title: AppautoscalingPolicy
menu:
  docs_{{ .version }}:
    identifier: appautoscalingpolicy-aws.kubeform.com
    name: AppautoscalingPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AppautoscalingPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AppautoscalingPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppautoscalingPolicySpec](#appautoscalingpolicyspec)***||
| `status` | ***[AppautoscalingPolicyStatus](#appautoscalingpolicystatus)***||
## AppautoscalingPolicySpec

Appears on:[AppautoscalingPolicy](#appautoscalingpolicy), [AppautoscalingPolicyStatus](#appautoscalingpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `alarms` | ***[]string***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `policyType` | ***string***| ***(Optional)*** |
| `resourceID` | ***string***||
| `scalableDimension` | ***string***||
| `serviceNamespace` | ***string***||
| `stepScalingPolicyConfiguration` | ***[[]AppautoscalingPolicySpecStepScalingPolicyConfiguration](#appautoscalingpolicyspecstepscalingpolicyconfiguration)***| ***(Optional)*** |
| `targetTrackingScalingPolicyConfiguration` | ***[[]AppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration](#appautoscalingpolicyspectargettrackingscalingpolicyconfiguration)***| ***(Optional)*** |
## AppautoscalingPolicySpecStepScalingPolicyConfiguration

Appears on:[AppautoscalingPolicySpec](#appautoscalingpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `adjustmentType` | ***string***| ***(Optional)*** |
| `cooldown` | ***int***| ***(Optional)*** |
| `metricAggregationType` | ***string***| ***(Optional)*** |
| `minAdjustmentMagnitude` | ***int***| ***(Optional)*** |
| `stepAdjustment` | ***[[]AppautoscalingPolicySpecStepScalingPolicyConfigurationStepAdjustment](#appautoscalingpolicyspecstepscalingpolicyconfigurationstepadjustment)***| ***(Optional)*** |
## AppautoscalingPolicySpecStepScalingPolicyConfigurationStepAdjustment

Appears on:[AppautoscalingPolicySpecStepScalingPolicyConfiguration](#appautoscalingpolicyspecstepscalingpolicyconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `metricIntervalLowerBound` | ***string***| ***(Optional)*** |
| `metricIntervalUpperBound` | ***string***| ***(Optional)*** |
| `scalingAdjustment` | ***int***||
## AppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration

Appears on:[AppautoscalingPolicySpec](#appautoscalingpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `customizedMetricSpecification` | ***[[]AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification](#appautoscalingpolicyspectargettrackingscalingpolicyconfigurationcustomizedmetricspecification)***| ***(Optional)*** |
| `disableScaleIn` | ***bool***| ***(Optional)*** |
| `predefinedMetricSpecification` | ***[[]AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification](#appautoscalingpolicyspectargettrackingscalingpolicyconfigurationpredefinedmetricspecification)***| ***(Optional)*** |
| `scaleInCooldown` | ***int***| ***(Optional)*** |
| `scaleOutCooldown` | ***int***| ***(Optional)*** |
| `targetValue` | ***encoding/json.Number***||
## AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification

Appears on:[AppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration](#appautoscalingpolicyspectargettrackingscalingpolicyconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dimensions` | ***[[]AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensions](#appautoscalingpolicyspectargettrackingscalingpolicyconfigurationcustomizedmetricspecificationdimensions)***| ***(Optional)*** |
| `metricName` | ***string***||
| `namespace` | ***string***||
| `statistic` | ***string***||
| `unit` | ***string***| ***(Optional)*** |
## AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensions

Appears on:[AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification](#appautoscalingpolicyspectargettrackingscalingpolicyconfigurationcustomizedmetricspecification)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `value` | ***string***||
## AppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification

Appears on:[AppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration](#appautoscalingpolicyspectargettrackingscalingpolicyconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `predefinedMetricType` | ***string***||
| `resourceLabel` | ***string***| ***(Optional)*** |
## AppautoscalingPolicyStatus

Appears on:[AppautoscalingPolicy](#appautoscalingpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppautoscalingPolicySpec](#appautoscalingpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AppautoscalingPolicyStatus](#appautoscalingpolicystatus)

---
