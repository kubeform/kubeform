package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AwsAppautoscalingPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppautoscalingPolicySpec   `json:"spec,omitempty"`
	Status            AwsAppautoscalingPolicyStatus `json:"status,omitempty"`
}

type AwsAppautoscalingPolicySpecStepScalingPolicyConfigurationStepAdjustment struct {
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound"`
	ScalingAdjustment        int    `json:"scaling_adjustment"`
}

type AwsAppautoscalingPolicySpecStepScalingPolicyConfiguration struct {
	MinAdjustmentMagnitude int                                                         `json:"min_adjustment_magnitude"`
	StepAdjustment         []AwsAppautoscalingPolicySpecStepScalingPolicyConfiguration `json:"step_adjustment"`
	AdjustmentType         string                                                      `json:"adjustment_type"`
	Cooldown               int                                                         `json:"cooldown"`
	MetricAggregationType  string                                                      `json:"metric_aggregation_type"`
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
	ScaleOutCooldown              int                                                                   `json:"scale_out_cooldown"`
	TargetValue                   float64                                                               `json:"target_value"`
	CustomizedMetricSpecification []AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration `json:"customized_metric_specification"`
	PredefinedMetricSpecification []AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration `json:"predefined_metric_specification"`
	DisableScaleIn                bool                                                                  `json:"disable_scale_in"`
	ScaleInCooldown               int                                                                   `json:"scale_in_cooldown"`
}

type AwsAppautoscalingPolicySpecStepAdjustment struct {
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound"`
	ScalingAdjustment        int    `json:"scaling_adjustment"`
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound"`
}

type AwsAppautoscalingPolicySpec struct {
	Name                                     string                        `json:"name"`
	ResourceId                               string                        `json:"resource_id"`
	ServiceNamespace                         string                        `json:"service_namespace"`
	Arn                                      string                        `json:"arn"`
	StepScalingPolicyConfiguration           []AwsAppautoscalingPolicySpec `json:"step_scaling_policy_configuration"`
	TargetTrackingScalingPolicyConfiguration []AwsAppautoscalingPolicySpec `json:"target_tracking_scaling_policy_configuration"`
	PolicyType                               string                        `json:"policy_type"`
	ScalableDimension                        string                        `json:"scalable_dimension"`
	Alarms                                   []string                      `json:"alarms"`
	MetricAggregationType                    string                        `json:"metric_aggregation_type"`
	StepAdjustment                           []AwsAppautoscalingPolicySpec `json:"step_adjustment"`
	AdjustmentType                           string                        `json:"adjustment_type"`
	Cooldown                                 int                           `json:"cooldown"`
	MinAdjustmentMagnitude                   int                           `json:"min_adjustment_magnitude"`
}

type AwsAppautoscalingPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAppautoscalingPolicyList is a list of AwsAppautoscalingPolicys
type AwsAppautoscalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppautoscalingPolicy CRD objects
	Items []AwsAppautoscalingPolicy `json:"items,omitempty"`
}
