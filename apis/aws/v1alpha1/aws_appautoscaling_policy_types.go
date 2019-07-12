package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppautoscalingPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppautoscalingPolicySpec   `json:"spec,omitempty"`
	Status            AwsAppautoscalingPolicyStatus `json:"status,omitempty"`
}

type AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensions struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification struct {
	Dimensions []AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification `json:"dimensions"`
	MetricName string                                                                                             `json:"metric_name"`
	Namespace  string                                                                                             `json:"namespace"`
	Statistic  string                                                                                             `json:"statistic"`
	Unit       string                                                                                             `json:"unit"`
}

type AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification struct {
	PredefinedMetricType string `json:"predefined_metric_type"`
	ResourceLabel        string `json:"resource_label"`
}

type AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration struct {
	CustomizedMetricSpecification []AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration `json:"customized_metric_specification"`
	PredefinedMetricSpecification []AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration `json:"predefined_metric_specification"`
	DisableScaleIn                bool                                                                  `json:"disable_scale_in"`
	ScaleInCooldown               int                                                                   `json:"scale_in_cooldown"`
	ScaleOutCooldown              int                                                                   `json:"scale_out_cooldown"`
	TargetValue                   float64                                                               `json:"target_value"`
}

type AwsAppautoscalingPolicySpecStepScalingPolicyConfigurationStepAdjustment struct {
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound"`
	ScalingAdjustment        int    `json:"scaling_adjustment"`
}

type AwsAppautoscalingPolicySpecStepScalingPolicyConfiguration struct {
	AdjustmentType         string                                                      `json:"adjustment_type"`
	Cooldown               int                                                         `json:"cooldown"`
	MetricAggregationType  string                                                      `json:"metric_aggregation_type"`
	MinAdjustmentMagnitude int                                                         `json:"min_adjustment_magnitude"`
	StepAdjustment         []AwsAppautoscalingPolicySpecStepScalingPolicyConfiguration `json:"step_adjustment"`
}

type AwsAppautoscalingPolicySpecStepAdjustment struct {
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound"`
	ScalingAdjustment        int    `json:"scaling_adjustment"`
}

type AwsAppautoscalingPolicySpec struct {
	AdjustmentType                           string                        `json:"adjustment_type"`
	Cooldown                                 int                           `json:"cooldown"`
	Arn                                      string                        `json:"arn"`
	TargetTrackingScalingPolicyConfiguration []AwsAppautoscalingPolicySpec `json:"target_tracking_scaling_policy_configuration"`
	ScalableDimension                        string                        `json:"scalable_dimension"`
	ServiceNamespace                         string                        `json:"service_namespace"`
	StepScalingPolicyConfiguration           []AwsAppautoscalingPolicySpec `json:"step_scaling_policy_configuration"`
	Alarms                                   []string                      `json:"alarms"`
	MinAdjustmentMagnitude                   int                           `json:"min_adjustment_magnitude"`
	Name                                     string                        `json:"name"`
	PolicyType                               string                        `json:"policy_type"`
	ResourceId                               string                        `json:"resource_id"`
	MetricAggregationType                    string                        `json:"metric_aggregation_type"`
	StepAdjustment                           []AwsAppautoscalingPolicySpec `json:"step_adjustment"`
}

type AwsAppautoscalingPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppautoscalingPolicyList is a list of AwsAppautoscalingPolicys
type AwsAppautoscalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppautoscalingPolicy CRD objects
	Items []AwsAppautoscalingPolicy `json:"items,omitempty"`
}
