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

type AwsAutoscalingPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAutoscalingPolicySpec   `json:"spec,omitempty"`
	Status            AwsAutoscalingPolicyStatus `json:"status,omitempty"`
}

type AwsAutoscalingPolicySpecStepAdjustment struct {
	ScalingAdjustment        int    `json:"scaling_adjustment"`
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification struct {
	PredefinedMetricType string `json:"predefined_metric_type"`
	ResourceLabel        string `json:"resource_label"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension struct {
	Value string `json:"value"`
	Name  string `json:"name"`
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

type AwsAutoscalingPolicySpec struct {
	AdjustmentType              string                     `json:"adjustment_type"`
	EstimatedInstanceWarmup     int                        `json:"estimated_instance_warmup"`
	MetricAggregationType       string                     `json:"metric_aggregation_type"`
	StepAdjustment              []AwsAutoscalingPolicySpec `json:"step_adjustment"`
	TargetTrackingConfiguration []AwsAutoscalingPolicySpec `json:"target_tracking_configuration"`
	Arn                         string                     `json:"arn"`
	AutoscalingGroupName        string                     `json:"autoscaling_group_name"`
	PolicyType                  string                     `json:"policy_type"`
	Cooldown                    int                        `json:"cooldown"`
	MinAdjustmentMagnitude      int                        `json:"min_adjustment_magnitude"`
	MinAdjustmentStep           int                        `json:"min_adjustment_step"`
	ScalingAdjustment           int                        `json:"scaling_adjustment"`
	Name                        string                     `json:"name"`
}

type AwsAutoscalingPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAutoscalingPolicyList is a list of AwsAutoscalingPolicys
type AwsAutoscalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAutoscalingPolicy CRD objects
	Items []AwsAutoscalingPolicy `json:"items,omitempty"`
}
