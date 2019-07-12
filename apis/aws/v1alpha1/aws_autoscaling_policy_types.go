package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAutoscalingPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAutoscalingPolicySpec   `json:"spec,omitempty"`
	Status            AwsAutoscalingPolicyStatus `json:"status,omitempty"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification struct {
	PredefinedMetricType string `json:"predefined_metric_type"`
	ResourceLabel        string `json:"resource_label"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification struct {
	MetricDimension []AwsAutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification `json:"metric_dimension"`
	MetricName      string                                                                             `json:"metric_name"`
	Namespace       string                                                                             `json:"namespace"`
	Statistic       string                                                                             `json:"statistic"`
	Unit            string                                                                             `json:"unit"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfiguration struct {
	PredefinedMetricSpecification []AwsAutoscalingPolicySpecTargetTrackingConfiguration `json:"predefined_metric_specification"`
	CustomizedMetricSpecification []AwsAutoscalingPolicySpecTargetTrackingConfiguration `json:"customized_metric_specification"`
	TargetValue                   float64                                               `json:"target_value"`
	DisableScaleIn                bool                                                  `json:"disable_scale_in"`
}

type AwsAutoscalingPolicySpecStepAdjustment struct {
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound"`
	ScalingAdjustment        int    `json:"scaling_adjustment"`
}

type AwsAutoscalingPolicySpec struct {
	Arn                         string                     `json:"arn"`
	EstimatedInstanceWarmup     int                        `json:"estimated_instance_warmup"`
	TargetTrackingConfiguration []AwsAutoscalingPolicySpec `json:"target_tracking_configuration"`
	MinAdjustmentMagnitude      int                        `json:"min_adjustment_magnitude"`
	MinAdjustmentStep           int                        `json:"min_adjustment_step"`
	Name                        string                     `json:"name"`
	AdjustmentType              string                     `json:"adjustment_type"`
	AutoscalingGroupName        string                     `json:"autoscaling_group_name"`
	PolicyType                  string                     `json:"policy_type"`
	Cooldown                    int                        `json:"cooldown"`
	MetricAggregationType       string                     `json:"metric_aggregation_type"`
	ScalingAdjustment           int                        `json:"scaling_adjustment"`
	StepAdjustment              []AwsAutoscalingPolicySpec `json:"step_adjustment"`
}

type AwsAutoscalingPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAutoscalingPolicyList is a list of AwsAutoscalingPolicys
type AwsAutoscalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAutoscalingPolicy CRD objects
	Items []AwsAutoscalingPolicy `json:"items,omitempty"`
}
