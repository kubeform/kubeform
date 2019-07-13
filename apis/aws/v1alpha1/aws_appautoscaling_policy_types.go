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
	ScalingAdjustment        int    `json:"scaling_adjustment"`
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound"`
}

type AwsAppautoscalingPolicySpecStepScalingPolicyConfiguration struct {
	MetricAggregationType  string                                                      `json:"metric_aggregation_type"`
	MinAdjustmentMagnitude int                                                         `json:"min_adjustment_magnitude"`
	StepAdjustment         []AwsAppautoscalingPolicySpecStepScalingPolicyConfiguration `json:"step_adjustment"`
	AdjustmentType         string                                                      `json:"adjustment_type"`
	Cooldown               int                                                         `json:"cooldown"`
}

type AwsAppautoscalingPolicySpecStepAdjustment struct {
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound"`
	ScalingAdjustment        int    `json:"scaling_adjustment"`
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
	DisableScaleIn                bool                                                                  `json:"disable_scale_in"`
	ScaleInCooldown               int                                                                   `json:"scale_in_cooldown"`
	ScaleOutCooldown              int                                                                   `json:"scale_out_cooldown"`
	TargetValue                   float64                                                               `json:"target_value"`
	CustomizedMetricSpecification []AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration `json:"customized_metric_specification"`
	PredefinedMetricSpecification []AwsAppautoscalingPolicySpecTargetTrackingScalingPolicyConfiguration `json:"predefined_metric_specification"`
}

type AwsAppautoscalingPolicySpec struct {
	StepScalingPolicyConfiguration           []AwsAppautoscalingPolicySpec `json:"step_scaling_policy_configuration"`
	AdjustmentType                           string                        `json:"adjustment_type"`
	MetricAggregationType                    string                        `json:"metric_aggregation_type"`
	MinAdjustmentMagnitude                   int                           `json:"min_adjustment_magnitude"`
	PolicyType                               string                        `json:"policy_type"`
	ScalableDimension                        string                        `json:"scalable_dimension"`
	StepAdjustment                           []AwsAppautoscalingPolicySpec `json:"step_adjustment"`
	Name                                     string                        `json:"name"`
	Alarms                                   []string                      `json:"alarms"`
	Arn                                      string                        `json:"arn"`
	ServiceNamespace                         string                        `json:"service_namespace"`
	Cooldown                                 int                           `json:"cooldown"`
	TargetTrackingScalingPolicyConfiguration []AwsAppautoscalingPolicySpec `json:"target_tracking_scaling_policy_configuration"`
	ResourceId                               string                        `json:"resource_id"`
}



type AwsAppautoscalingPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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