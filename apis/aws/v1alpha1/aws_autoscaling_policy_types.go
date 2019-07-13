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

type AwsAutoscalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification struct {
	PredefinedMetricType string `json:"predefined_metric_type"`
	ResourceLabel        string `json:"resource_label"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification struct {
	Statistic       string                                                                             `json:"statistic"`
	Unit            string                                                                             `json:"unit"`
	MetricDimension []AwsAutoscalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification `json:"metric_dimension"`
	MetricName      string                                                                             `json:"metric_name"`
	Namespace       string                                                                             `json:"namespace"`
}

type AwsAutoscalingPolicySpecTargetTrackingConfiguration struct {
	DisableScaleIn                bool                                                  `json:"disable_scale_in"`
	PredefinedMetricSpecification []AwsAutoscalingPolicySpecTargetTrackingConfiguration `json:"predefined_metric_specification"`
	CustomizedMetricSpecification []AwsAutoscalingPolicySpecTargetTrackingConfiguration `json:"customized_metric_specification"`
	TargetValue                   float64                                               `json:"target_value"`
}

type AwsAutoscalingPolicySpecStepAdjustment struct {
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound"`
	ScalingAdjustment        int    `json:"scaling_adjustment"`
}

type AwsAutoscalingPolicySpec struct {
	TargetTrackingConfiguration []AwsAutoscalingPolicySpec `json:"target_tracking_configuration"`
	Arn                         string                     `json:"arn"`
	AdjustmentType              string                     `json:"adjustment_type"`
	PolicyType                  string                     `json:"policy_type"`
	MinAdjustmentMagnitude      int                        `json:"min_adjustment_magnitude"`
	MinAdjustmentStep           int                        `json:"min_adjustment_step"`
	ScalingAdjustment           int                        `json:"scaling_adjustment"`
	Name                        string                     `json:"name"`
	AutoscalingGroupName        string                     `json:"autoscaling_group_name"`
	Cooldown                    int                        `json:"cooldown"`
	EstimatedInstanceWarmup     int                        `json:"estimated_instance_warmup"`
	MetricAggregationType       string                     `json:"metric_aggregation_type"`
	StepAdjustment              []AwsAutoscalingPolicySpec `json:"step_adjustment"`
}



type AwsAutoscalingPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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